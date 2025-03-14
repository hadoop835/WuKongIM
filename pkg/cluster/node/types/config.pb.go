// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: pkg/cluster2/node/types/config.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodeRole int32

const (
	NodeRole_NodeRoleReplica NodeRole = 0 // 副本服务，参与投票，可以成为领导
	NodeRole_NodeRoleProxy   NodeRole = 1 // 代理服务，仅仅做转发加速效果，类似cdn （没有投票权）
)

// Enum value maps for NodeRole.
var (
	NodeRole_name = map[int32]string{
		0: "NodeRoleReplica",
		1: "NodeRoleProxy",
	}
	NodeRole_value = map[string]int32{
		"NodeRoleReplica": 0,
		"NodeRoleProxy":   1,
	}
)

func (x NodeRole) Enum() *NodeRole {
	p := new(NodeRole)
	*p = x
	return p
}

func (x NodeRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeRole) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_cluster2_node_types_config_proto_enumTypes[0].Descriptor()
}

func (NodeRole) Type() protoreflect.EnumType {
	return &file_pkg_cluster2_node_types_config_proto_enumTypes[0]
}

func (x NodeRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeRole.Descriptor instead.
func (NodeRole) EnumDescriptor() ([]byte, []int) {
	return file_pkg_cluster2_node_types_config_proto_rawDescGZIP(), []int{0}
}

type NodeStatus int32

const (
	NodeStatus_NodeStatusUnkown   NodeStatus = 0 // 未知
	NodeStatus_NodeStatusWillJoin NodeStatus = 1 // 将要加入
	NodeStatus_NodeStatusJoining  NodeStatus = 2 // 加入中
	NodeStatus_NodeStatusJoined   NodeStatus = 3 // 加入完成
)

// Enum value maps for NodeStatus.
var (
	NodeStatus_name = map[int32]string{
		0: "NodeStatusUnkown",
		1: "NodeStatusWillJoin",
		2: "NodeStatusJoining",
		3: "NodeStatusJoined",
	}
	NodeStatus_value = map[string]int32{
		"NodeStatusUnkown":   0,
		"NodeStatusWillJoin": 1,
		"NodeStatusJoining":  2,
		"NodeStatusJoined":   3,
	}
)

func (x NodeStatus) Enum() *NodeStatus {
	p := new(NodeStatus)
	*p = x
	return p
}

func (x NodeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_cluster2_node_types_config_proto_enumTypes[1].Descriptor()
}

func (NodeStatus) Type() protoreflect.EnumType {
	return &file_pkg_cluster2_node_types_config_proto_enumTypes[1]
}

func (x NodeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeStatus.Descriptor instead.
func (NodeStatus) EnumDescriptor() ([]byte, []int) {
	return file_pkg_cluster2_node_types_config_proto_rawDescGZIP(), []int{1}
}

type MigrateStatus int32

const (
	MigrateStatus_MigrateStatusUnkown MigrateStatus = 0 // 未知
	MigrateStatus_MigrateStatusWill   MigrateStatus = 1 // 准备迁移
	MigrateStatus_MigrateStatusDoing  MigrateStatus = 2 // 迁移中
	MigrateStatus_MigrateStatusDone   MigrateStatus = 3 // 迁移完成
)

// Enum value maps for MigrateStatus.
var (
	MigrateStatus_name = map[int32]string{
		0: "MigrateStatusUnkown",
		1: "MigrateStatusWill",
		2: "MigrateStatusDoing",
		3: "MigrateStatusDone",
	}
	MigrateStatus_value = map[string]int32{
		"MigrateStatusUnkown": 0,
		"MigrateStatusWill":   1,
		"MigrateStatusDoing":  2,
		"MigrateStatusDone":   3,
	}
)

func (x MigrateStatus) Enum() *MigrateStatus {
	p := new(MigrateStatus)
	*p = x
	return p
}

func (x MigrateStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MigrateStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_cluster2_node_types_config_proto_enumTypes[2].Descriptor()
}

func (MigrateStatus) Type() protoreflect.EnumType {
	return &file_pkg_cluster2_node_types_config_proto_enumTypes[2]
}

func (x MigrateStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MigrateStatus.Descriptor instead.
func (MigrateStatus) EnumDescriptor() ([]byte, []int) {
	return file_pkg_cluster2_node_types_config_proto_rawDescGZIP(), []int{2}
}

type SlotStatus int32

const (
	SlotStatus_SlotStatusNormal         SlotStatus = 0 // 未知
	SlotStatus_SlotStatusCandidate      SlotStatus = 1 // 进入领导候选状态
	SlotStatus_SlotStatusLeaderTransfer SlotStatus = 2 // 领导转移
)

// Enum value maps for SlotStatus.
var (
	SlotStatus_name = map[int32]string{
		0: "SlotStatusNormal",
		1: "SlotStatusCandidate",
		2: "SlotStatusLeaderTransfer",
	}
	SlotStatus_value = map[string]int32{
		"SlotStatusNormal":         0,
		"SlotStatusCandidate":      1,
		"SlotStatusLeaderTransfer": 2,
	}
)

func (x SlotStatus) Enum() *SlotStatus {
	p := new(SlotStatus)
	*p = x
	return p
}

func (x SlotStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SlotStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_cluster2_node_types_config_proto_enumTypes[3].Descriptor()
}

func (SlotStatus) Type() protoreflect.EnumType {
	return &file_pkg_cluster2_node_types_config_proto_enumTypes[3]
}

func (x SlotStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SlotStatus.Descriptor instead.
func (SlotStatus) EnumDescriptor() ([]byte, []int) {
	return file_pkg_cluster2_node_types_config_proto_rawDescGZIP(), []int{3}
}

type LearnerStatus int32

const (
	LearnerStatus_LearnerStatusLearning LearnerStatus = 0 // 学习中
	LearnerStatus_LearnerStatusGraduate LearnerStatus = 1 // 已毕业
)

// Enum value maps for LearnerStatus.
var (
	LearnerStatus_name = map[int32]string{
		0: "LearnerStatusLearning",
		1: "LearnerStatusGraduate",
	}
	LearnerStatus_value = map[string]int32{
		"LearnerStatusLearning": 0,
		"LearnerStatusGraduate": 1,
	}
)

func (x LearnerStatus) Enum() *LearnerStatus {
	p := new(LearnerStatus)
	*p = x
	return p
}

func (x LearnerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LearnerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_cluster2_node_types_config_proto_enumTypes[4].Descriptor()
}

func (LearnerStatus) Type() protoreflect.EnumType {
	return &file_pkg_cluster2_node_types_config_proto_enumTypes[4]
}

func (x LearnerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LearnerStatus.Descriptor instead.
func (LearnerStatus) EnumDescriptor() ([]byte, []int) {
	return file_pkg_cluster2_node_types_config_proto_rawDescGZIP(), []int{4}
}

type Config struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Version             uint64                 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`                         // 配置版本
	SlotCount           uint32                 `protobuf:"varint,2,opt,name=slotCount,proto3" json:"slotCount,omitempty"`                     // 槽位数量
	SlotReplicaCount    uint32                 `protobuf:"varint,3,opt,name=slotReplicaCount,proto3" json:"slotReplicaCount,omitempty"`       // 槽最大副本数量  (这个数量包含领导，比如副本为3，则是1个领导2个副本节点)
	ChannelReplicaCount uint32                 `protobuf:"varint,4,opt,name=channelReplicaCount,proto3" json:"channelReplicaCount,omitempty"` // 频道最大副本数量
	Term                uint32                 `protobuf:"varint,5,opt,name=term,proto3" json:"term,omitempty"`                               // 领导任期
	MigrateFrom         uint64                 `protobuf:"varint,6,opt,name=migrateFrom,proto3" json:"migrateFrom,omitempty"`                 // 迁移的源节点
	MigrateTo           uint64                 `protobuf:"varint,7,opt,name=migrateTo,proto3" json:"migrateTo,omitempty"`                     // 迁移的目标节点
	Learners            []uint64               `protobuf:"varint,8,rep,packed,name=learners,proto3" json:"learners,omitempty"`                // 学习者列表
	Nodes               []*Node                `protobuf:"bytes,9,rep,name=nodes,proto3" json:"nodes,omitempty"`                              // 分布式中的节点
	Slots               []*Slot                `protobuf:"bytes,10,rep,name=slots,proto3" json:"slots,omitempty"`                             // 分布式中的槽位
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_pkg_cluster2_node_types_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cluster2_node_types_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_pkg_cluster2_node_types_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Config) GetSlotCount() uint32 {
	if x != nil {
		return x.SlotCount
	}
	return 0
}

func (x *Config) GetSlotReplicaCount() uint32 {
	if x != nil {
		return x.SlotReplicaCount
	}
	return 0
}

func (x *Config) GetChannelReplicaCount() uint32 {
	if x != nil {
		return x.ChannelReplicaCount
	}
	return 0
}

func (x *Config) GetTerm() uint32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *Config) GetMigrateFrom() uint64 {
	if x != nil {
		return x.MigrateFrom
	}
	return 0
}

func (x *Config) GetMigrateTo() uint64 {
	if x != nil {
		return x.MigrateTo
	}
	return 0
}

func (x *Config) GetLearners() []uint64 {
	if x != nil {
		return x.Learners
	}
	return nil
}

func (x *Config) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Config) GetSlots() []*Slot {
	if x != nil {
		return x.Slots
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                      // 节点id
	ClusterAddr   string                 `protobuf:"bytes,2,opt,name=clusterAddr,proto3" json:"clusterAddr,omitempty"`     // 节点分布式通讯地址
	ApiServerAddr string                 `protobuf:"bytes,3,opt,name=apiServerAddr,proto3" json:"apiServerAddr,omitempty"` // 节点api服务地址
	Join          bool                   `protobuf:"varint,4,opt,name=join,proto3" json:"join,omitempty"`                  // 是否是加入集群的节点，false表示初始节点 true表示后面新加入的节点
	// repeated SlotMigrate exports = 5; // 正在迁出的槽位
	// repeated SlotMigrate imports = 6; // 正在迁入的槽位
	Online        bool       `protobuf:"varint,5,opt,name=online,proto3" json:"online,omitempty"`                        // 是否在线
	OfflineCount  uint32     `protobuf:"varint,6,opt,name=offlineCount,proto3" json:"offlineCount,omitempty"`            // 离线次数
	LastOffline   int64      `protobuf:"varint,7,opt,name=lastOffline,proto3" json:"lastOffline,omitempty"`              // 最后一次离线时间
	AllowVote     bool       `protobuf:"varint,8,opt,name=allowVote,proto3" json:"allowVote,omitempty"`                  // 节点是否允许投票
	Role          NodeRole   `protobuf:"varint,9,opt,name=role,proto3,enum=types.NodeRole" json:"role,omitempty"`        // 节点角色
	Status        NodeStatus `protobuf:"varint,10,opt,name=status,proto3,enum=types.NodeStatus" json:"status,omitempty"` // 节点状态
	CreatedAt     int64      `protobuf:"varint,11,opt,name=createdAt,proto3" json:"createdAt,omitempty"`                 // 创建时间
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Node) Reset() {
	*x = Node{}
	mi := &file_pkg_cluster2_node_types_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cluster2_node_types_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_pkg_cluster2_node_types_config_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetClusterAddr() string {
	if x != nil {
		return x.ClusterAddr
	}
	return ""
}

func (x *Node) GetApiServerAddr() string {
	if x != nil {
		return x.ApiServerAddr
	}
	return ""
}

func (x *Node) GetJoin() bool {
	if x != nil {
		return x.Join
	}
	return false
}

func (x *Node) GetOnline() bool {
	if x != nil {
		return x.Online
	}
	return false
}

func (x *Node) GetOfflineCount() uint32 {
	if x != nil {
		return x.OfflineCount
	}
	return 0
}

func (x *Node) GetLastOffline() int64 {
	if x != nil {
		return x.LastOffline
	}
	return 0
}

func (x *Node) GetAllowVote() bool {
	if x != nil {
		return x.AllowVote
	}
	return false
}

func (x *Node) GetRole() NodeRole {
	if x != nil {
		return x.Role
	}
	return NodeRole_NodeRoleReplica
}

func (x *Node) GetStatus() NodeStatus {
	if x != nil {
		return x.Status
	}
	return NodeStatus_NodeStatusUnkown
}

func (x *Node) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type Slot struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                // 槽位id
	Leader        uint64                 `protobuf:"varint,2,opt,name=leader,proto3" json:"leader,omitempty"`                        // 槽位的领导节点
	Term          uint32                 `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`                            // 槽位的领导任期
	Replicas      []uint64               `protobuf:"varint,4,rep,packed,name=replicas,proto3" json:"replicas,omitempty"`             // 副本节点(包含领导节点)
	Learners      []uint64               `protobuf:"varint,6,rep,packed,name=learners,proto3" json:"learners,omitempty"`             // 学习者列表
	MigrateFrom   uint64                 `protobuf:"varint,7,opt,name=migrateFrom,proto3" json:"migrateFrom,omitempty"`              // 迁移的源节点
	MigrateTo     uint64                 `protobuf:"varint,8,opt,name=migrateTo,proto3" json:"migrateTo,omitempty"`                  // 迁移的目标节点
	ExpectLeader  uint64                 `protobuf:"varint,9,opt,name=expectLeader,proto3" json:"expectLeader,omitempty"`            // 期望选举到的领导节点
	Status        SlotStatus             `protobuf:"varint,10,opt,name=status,proto3,enum=types.SlotStatus" json:"status,omitempty"` // 槽位状态
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Slot) Reset() {
	*x = Slot{}
	mi := &file_pkg_cluster2_node_types_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Slot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slot) ProtoMessage() {}

func (x *Slot) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cluster2_node_types_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slot.ProtoReflect.Descriptor instead.
func (*Slot) Descriptor() ([]byte, []int) {
	return file_pkg_cluster2_node_types_config_proto_rawDescGZIP(), []int{2}
}

func (x *Slot) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Slot) GetLeader() uint64 {
	if x != nil {
		return x.Leader
	}
	return 0
}

func (x *Slot) GetTerm() uint32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *Slot) GetReplicas() []uint64 {
	if x != nil {
		return x.Replicas
	}
	return nil
}

func (x *Slot) GetLearners() []uint64 {
	if x != nil {
		return x.Learners
	}
	return nil
}

func (x *Slot) GetMigrateFrom() uint64 {
	if x != nil {
		return x.MigrateFrom
	}
	return 0
}

func (x *Slot) GetMigrateTo() uint64 {
	if x != nil {
		return x.MigrateTo
	}
	return 0
}

func (x *Slot) GetExpectLeader() uint64 {
	if x != nil {
		return x.ExpectLeader
	}
	return 0
}

func (x *Slot) GetStatus() SlotStatus {
	if x != nil {
		return x.Status
	}
	return SlotStatus_SlotStatusNormal
}

type SlotMigrate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          uint64                 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"` // 迁移的源节点
	To            uint64                 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`     // 迁移的目标节点
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SlotMigrate) Reset() {
	*x = SlotMigrate{}
	mi := &file_pkg_cluster2_node_types_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SlotMigrate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlotMigrate) ProtoMessage() {}

func (x *SlotMigrate) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cluster2_node_types_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlotMigrate.ProtoReflect.Descriptor instead.
func (*SlotMigrate) Descriptor() ([]byte, []int) {
	return file_pkg_cluster2_node_types_config_proto_rawDescGZIP(), []int{3}
}

func (x *SlotMigrate) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *SlotMigrate) GetTo() uint64 {
	if x != nil {
		return x.To
	}
	return 0
}

type Learner struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LearnerId     uint64                 `protobuf:"varint,1,opt,name=learnerId,proto3" json:"learnerId,omitempty"`                    // 学习者节点id
	Status        LearnerStatus          `protobuf:"varint,2,opt,name=status,proto3,enum=types.LearnerStatus" json:"status,omitempty"` // 学习状态
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Learner) Reset() {
	*x = Learner{}
	mi := &file_pkg_cluster2_node_types_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Learner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Learner) ProtoMessage() {}

func (x *Learner) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cluster2_node_types_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Learner.ProtoReflect.Descriptor instead.
func (*Learner) Descriptor() ([]byte, []int) {
	return file_pkg_cluster2_node_types_config_proto_rawDescGZIP(), []int{4}
}

func (x *Learner) GetLearnerId() uint64 {
	if x != nil {
		return x.LearnerId
	}
	return 0
}

func (x *Learner) GetStatus() LearnerStatus {
	if x != nil {
		return x.Status
	}
	return LearnerStatus_LearnerStatusLearning
}

var File_pkg_cluster2_node_types_config_proto protoreflect.FileDescriptor

var file_pkg_cluster2_node_types_config_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x32, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xd4, 0x02,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6c, 0x6f, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x6c, 0x6f, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x73, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x73, 0x6c, 0x6f, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x13,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x21,
	0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x05, 0x73,
	0x6c, 0x6f, 0x74, 0x73, 0x22, 0xdc, 0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x24, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x56, 0x6f, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x89, 0x02, 0x0a, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x31, 0x0a, 0x0b, 0x53, 0x6c, 0x6f, 0x74, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x74, 0x6f, 0x22, 0x55, 0x0a, 0x07, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x32, 0x0a, 0x08, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x10, 0x01, 0x2a, 0x67, 0x0a,
	0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x6e, 0x6b, 0x6f, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x57,
	0x69, 0x6c, 0x6c, 0x4a, 0x6f, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x02,
	0x12, 0x14, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a, 0x6f,
	0x69, 0x6e, 0x65, 0x64, 0x10, 0x03, 0x2a, 0x6e, 0x0a, 0x0d, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x6e, 0x6b, 0x6f, 0x77, 0x6e, 0x10, 0x00,
	0x12, 0x15, 0x0a, 0x11, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x57, 0x69, 0x6c, 0x6c, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x6f, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12,
	0x15, 0x0a, 0x11, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x44, 0x6f, 0x6e, 0x65, 0x10, 0x03, 0x2a, 0x59, 0x0a, 0x0a, 0x53, 0x6c, 0x6f, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x6c, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x6c,
	0x6f, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x6c, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x10,
	0x02, 0x2a, 0x45, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x19, 0x0a,
	0x15, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x47, 0x72,
	0x61, 0x64, 0x75, 0x61, 0x74, 0x65, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_cluster2_node_types_config_proto_rawDescOnce sync.Once
	file_pkg_cluster2_node_types_config_proto_rawDescData = file_pkg_cluster2_node_types_config_proto_rawDesc
)

func file_pkg_cluster2_node_types_config_proto_rawDescGZIP() []byte {
	file_pkg_cluster2_node_types_config_proto_rawDescOnce.Do(func() {
		file_pkg_cluster2_node_types_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_cluster2_node_types_config_proto_rawDescData)
	})
	return file_pkg_cluster2_node_types_config_proto_rawDescData
}

var file_pkg_cluster2_node_types_config_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_pkg_cluster2_node_types_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_cluster2_node_types_config_proto_goTypes = []any{
	(NodeRole)(0),       // 0: types.NodeRole
	(NodeStatus)(0),     // 1: types.NodeStatus
	(MigrateStatus)(0),  // 2: types.MigrateStatus
	(SlotStatus)(0),     // 3: types.SlotStatus
	(LearnerStatus)(0),  // 4: types.LearnerStatus
	(*Config)(nil),      // 5: types.Config
	(*Node)(nil),        // 6: types.Node
	(*Slot)(nil),        // 7: types.Slot
	(*SlotMigrate)(nil), // 8: types.SlotMigrate
	(*Learner)(nil),     // 9: types.Learner
}
var file_pkg_cluster2_node_types_config_proto_depIdxs = []int32{
	6, // 0: types.Config.nodes:type_name -> types.Node
	7, // 1: types.Config.slots:type_name -> types.Slot
	0, // 2: types.Node.role:type_name -> types.NodeRole
	1, // 3: types.Node.status:type_name -> types.NodeStatus
	3, // 4: types.Slot.status:type_name -> types.SlotStatus
	4, // 5: types.Learner.status:type_name -> types.LearnerStatus
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_cluster2_node_types_config_proto_init() }
func file_pkg_cluster2_node_types_config_proto_init() {
	if File_pkg_cluster2_node_types_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_cluster2_node_types_config_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_cluster2_node_types_config_proto_goTypes,
		DependencyIndexes: file_pkg_cluster2_node_types_config_proto_depIdxs,
		EnumInfos:         file_pkg_cluster2_node_types_config_proto_enumTypes,
		MessageInfos:      file_pkg_cluster2_node_types_config_proto_msgTypes,
	}.Build()
	File_pkg_cluster2_node_types_config_proto = out.File
	file_pkg_cluster2_node_types_config_proto_rawDesc = nil
	file_pkg_cluster2_node_types_config_proto_goTypes = nil
	file_pkg_cluster2_node_types_config_proto_depIdxs = nil
}
