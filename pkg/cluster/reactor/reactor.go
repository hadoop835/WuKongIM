package reactor

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/WuKongIM/WuKongIM/pkg/cluster/replica"
	"github.com/WuKongIM/WuKongIM/pkg/wklog"
	"github.com/lni/goutils/syncutil"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

type Reactor struct {
	subReactors []*ReactorSub
	opts        *Options
	mu          sync.RWMutex
	taskPool    *ants.Pool
	wklog.Log

	processInitC              chan *initReq              // 处理频道初始化
	processConflictCheckC     chan *conflictCheckReq     // 冲突检查请求
	processGetLogC            chan *getLogReq            // 获取日志请求
	processStoreAppendC       chan AppendLogReq          // 存储追加日志请求
	processApplyLogC          chan *applyLogReq          // 应用日志请求
	processLearnerToFollowerC chan *learnerToFollowerReq // 从learner转为follower
	processLearnerToLeaderC   chan *learnerToLeaderReq   // 从learner转为leader
	processFollowerToLeaderC  chan *followerToLeaderReq  // 从follower转为leader

	stopper *syncutil.Stopper

	request IRequest
}

func New(opts *Options) *Reactor {
	r := &Reactor{
		opts:    opts,
		Log:     wklog.NewWKLog(fmt.Sprintf("Reactor[%d][%s]", opts.NodeId, opts.ReactorType.String())),
		stopper: syncutil.NewStopper(),

		processInitC:              make(chan *initReq, 1024),
		processConflictCheckC:     make(chan *conflictCheckReq, 1024),
		processGetLogC:            make(chan *getLogReq, 1024),
		processStoreAppendC:       make(chan AppendLogReq, 1024),
		processApplyLogC:          make(chan *applyLogReq, 1024),
		processLearnerToFollowerC: make(chan *learnerToFollowerReq, 1024),
		processLearnerToLeaderC:   make(chan *learnerToLeaderReq, 1024),
		processFollowerToLeaderC:  make(chan *followerToLeaderReq, 1024),
		request:                   opts.Request,
	}
	taskPool, err := ants.NewPool(opts.TaskPoolSize, ants.WithPanicHandler(func(err interface{}) {
		stack := debug.Stack()
		r.Panic("执行任务失败", zap.Any("error", err), zap.String("stack", string(stack)))

	}))
	if err != nil {
		r.Panic("create task pool error", zap.Error(err))
	}
	r.taskPool = taskPool

	for i := 0; i < int(r.opts.SubReactorNum); i++ {
		sub := r.newReactorSub(i)
		r.subReactors = append(r.subReactors, sub)
	}
	return r
}

func (r *Reactor) Start() error {

	// 高并发处理，适用于分散的耗时任务
	for i := 0; i < 100; i++ {

	}

	// 中并发处理，适合于分散但是不是很耗时的任务
	for i := 0; i < 10; i++ {

		r.stopper.RunWorker(r.processGetLogLoop)

		r.stopper.RunWorker(r.processLearnerToFollowerLoop)
		r.stopper.RunWorker(r.processLearnerToLeaderLoop)
		r.stopper.RunWorker(r.processFollowerToLeaderLoop)
	}

	// 低并发处理，适合于集中的耗时任务，这样可以合并请求批量处理
	for i := 0; i < 1; i++ {
		r.stopper.RunWorker(r.processInitLoop)
		r.stopper.RunWorker(r.processApplyLogLoop)
		r.stopper.RunWorker(r.processStoreAppendLoop) // 追加日志的协程不需要太多，因为追加日志会进行日志合并，如果协程太多反而频繁操作db导致性能下降
		r.stopper.RunWorker(r.processConflictCheckLoop)
	}

	for _, sub := range r.subReactors {
		err := sub.Start()
		if err != nil {
			return err
		}

	}
	return nil
}

func (r *Reactor) Stop() {

	for _, sub := range r.subReactors {
		sub.Stop()
	}
	r.stopper.Stop()

}
func (r *Reactor) ProposeAndWait(ctx context.Context, handleKey string, logs []replica.Log) ([]ProposeResult, error) {
	sub := r.reactorSub(handleKey)
	return sub.proposeAndWait(ctx, handleKey, logs)
}

func (r *Reactor) AddHandler(key string, handler IHandler) {
	h := getHandlerFromPool()
	h.init(key, handler, r)
	sub := r.reactorSub(key)
	sub.addHandler(h)
}

func (r *Reactor) RemoveHandler(key string) {

	sub := r.reactorSub(key)
	h := sub.removeHandler(key)
	if h != nil {
		putHandlerToPool(h)
	}
}

func (r *Reactor) Handler(key string) IHandler {
	h := r.handler(key)
	if h == nil {
		return nil
	}
	return h.handler
}

func (r *Reactor) handler(key string) *handler {
	sub := r.reactorSub(key)
	h := sub.handler(key)
	if h == nil {
		return nil
	}
	return h
}

func (r *Reactor) Step(key string, msg replica.Message) {
	sub := r.reactorSub(key)
	sub.step(key, msg)
}

func (r *Reactor) StepWait(key string, msg replica.Message) error {
	sub := r.reactorSub(key)
	return sub.stepWait(key, msg)
}

func (r *Reactor) ExistHandler(key string) bool {
	sub := r.reactorSub(key)
	return sub.existHandler(key)
}

func (r *Reactor) AddMessage(m Message) {
	sub := r.reactorSub(m.HandlerKey)
	sub.addMessage(m)
}

func (r *Reactor) HandlerLen() int {
	len := 0
	for _, sub := range r.subReactors {
		len += sub.handlerLen()
	}
	return len
}

func (r *Reactor) IteratorHandler(f func(h IHandler) bool) {
	for _, sub := range r.subReactors {
		sub.iterator(func(h *handler) bool {
			return f(h.handler)
		})
	}
}

func (r *Reactor) newReactorSub(i int) *ReactorSub {
	return NewReactorSub(i, r)
}

func (r *Reactor) reactorSub(key string) *ReactorSub {
	r.mu.RLock()
	defer r.mu.RUnlock()
	idx := hashWthString(key)
	return r.subReactors[idx%uint32(r.opts.SubReactorNum)]
}

func (r *Reactor) newMessageQueue() *MessageQueue {
	return NewMessageQueue(r.opts.ReceiveQueueLength, r.opts.MaxReceiveQueueSize)
}

func (r *Reactor) WithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*10)
}
