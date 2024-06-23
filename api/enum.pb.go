// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: enum.proto

package api

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

// 数据状态
type Status int32

const (
	// 全部
	Status_STATUS_ALL Status = 0
	// 启用
	Status_STATUS_ENABLE Status = 1
	// 禁用
	Status_STATUS_DISABLE Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_ALL",
		1: "STATUS_ENABLE",
		2: "STATUS_DISABLE",
	}
	Status_value = map[string]int32{
		"STATUS_ALL":     0,
		"STATUS_ENABLE":  1,
		"STATUS_DISABLE": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{0}
}

// 性别
type Gender int32

const (
	// 全部
	Gender_GENDER_ALL Gender = 0
	// 男
	Gender_GENDER_MALE Gender = 1
	// 女
	Gender_GENDER_FEMALE Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "GENDER_ALL",
		1: "GENDER_MALE",
		2: "GENDER_FEMALE",
	}
	Gender_value = map[string]int32{
		"GENDER_ALL":    0,
		"GENDER_MALE":   1,
		"GENDER_FEMALE": 2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[1].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[1]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{1}
}

// 系统全局默认角色， 区别于空间下自定义角色类型
type Role int32

const (
	// 全部 / 未知
	Role_ROLE_ALL Role = 0
	// 管理员
	Role_ROLE_SUPPER_ADMIN Role = 1
	// 普通管理员
	Role_ROLE_ADMIN Role = 2
	// 普通用户
	Role_ROLE_USER Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_ALL",
		1: "ROLE_SUPPER_ADMIN",
		2: "ROLE_ADMIN",
		3: "ROLE_USER",
	}
	Role_value = map[string]int32{
		"ROLE_ALL":          0,
		"ROLE_SUPPER_ADMIN": 1,
		"ROLE_ADMIN":        2,
		"ROLE_USER":         3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[2].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[2]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{2}
}

// 数据来源
type DataSource int32

const (
	// 全部
	DataSource_DATA_SOURCE_ALL DataSource = 0
	// 本地
	DataSource_DATA_SOURCE_LOCAL DataSource = 1
	// 远程
	DataSource_DATA_SOURCE_REMOTE DataSource = 2
)

// Enum value maps for DataSource.
var (
	DataSource_name = map[int32]string{
		0: "DATA_SOURCE_ALL",
		1: "DATA_SOURCE_LOCAL",
		2: "DATA_SOURCE_REMOTE",
	}
	DataSource_value = map[string]int32{
		"DATA_SOURCE_ALL":    0,
		"DATA_SOURCE_LOCAL":  1,
		"DATA_SOURCE_REMOTE": 2,
	}
)

func (x DataSource) Enum() *DataSource {
	p := new(DataSource)
	*p = x
	return p
}

func (x DataSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSource) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[3].Descriptor()
}

func (DataSource) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[3]
}

func (x DataSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSource.Descriptor instead.
func (DataSource) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{3}
}

// 验证码类型
type CaptchaType int32

const (
	// 图片验证码
	CaptchaType_CAPTCHA_TYPE_IMAGE CaptchaType = 0
	// 音频验证码
	CaptchaType_CAPTCHA_TYPE_AUDIO CaptchaType = 1
)

// Enum value maps for CaptchaType.
var (
	CaptchaType_name = map[int32]string{
		0: "CAPTCHA_TYPE_IMAGE",
		1: "CAPTCHA_TYPE_AUDIO",
	}
	CaptchaType_value = map[string]int32{
		"CAPTCHA_TYPE_IMAGE": 0,
		"CAPTCHA_TYPE_AUDIO": 1,
	}
)

func (x CaptchaType) Enum() *CaptchaType {
	p := new(CaptchaType)
	*p = x
	return p
}

func (x CaptchaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CaptchaType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[4].Descriptor()
}

func (CaptchaType) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[4]
}

func (x CaptchaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CaptchaType.Descriptor instead.
func (CaptchaType) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{4}
}

// 领域类型枚举
type DomainType int32

const (
	// 未知领域类型
	DomainType_DomainTypeUnknown DomainType = 0
	// 系统领域
	DomainType_DomainTypeSystem DomainType = 1
	// 监控领域
	DomainType_DomainTypeMonitor DomainType = 2
)

// Enum value maps for DomainType.
var (
	DomainType_name = map[int32]string{
		0: "DomainTypeUnknown",
		1: "DomainTypeSystem",
		2: "DomainTypeMonitor",
	}
	DomainType_value = map[string]int32{
		"DomainTypeUnknown": 0,
		"DomainTypeSystem":  1,
		"DomainTypeMonitor": 2,
	}
)

func (x DomainType) Enum() *DomainType {
	p := new(DomainType)
	*p = x
	return p
}

func (x DomainType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DomainType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[5].Descriptor()
}

func (DomainType) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[5]
}

func (x DomainType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DomainType.Descriptor instead.
func (DomainType) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{5}
}

// 模块类型枚举
type ModuleType int32

const (
	// 未知模块类型
	ModuleType_ModuleTypeUnknown ModuleType = 0
	// 接口模块
	ModuleType_ModelTypeApi ModuleType = 1
	// 菜单模块
	ModuleType_ModelTypeMenu ModuleType = 2
	// 角色模块
	ModuleType_ModelTypeRole ModuleType = 3
	// 用户模块
	ModuleType_ModelTypeUser ModuleType = 4
	// 字典模块
	ModuleType_ModelTypeDict ModuleType = 5
	// 配置模块
	ModuleType_ModelTypeConfig ModuleType = 6
	// 日志模块
	ModuleType_ModelTypeLog ModuleType = 7
	// 任务模块
	ModuleType_ModelTypeJob ModuleType = 8
	// 通知模块
	ModuleType_ModelTypeNotify ModuleType = 9
	// 系统模块
	ModuleType_ModelTypeSystem ModuleType = 10
	// 监控模块
	ModuleType_ModelTypeMonitor ModuleType = 11
)

// Enum value maps for ModuleType.
var (
	ModuleType_name = map[int32]string{
		0:  "ModuleTypeUnknown",
		1:  "ModelTypeApi",
		2:  "ModelTypeMenu",
		3:  "ModelTypeRole",
		4:  "ModelTypeUser",
		5:  "ModelTypeDict",
		6:  "ModelTypeConfig",
		7:  "ModelTypeLog",
		8:  "ModelTypeJob",
		9:  "ModelTypeNotify",
		10: "ModelTypeSystem",
		11: "ModelTypeMonitor",
	}
	ModuleType_value = map[string]int32{
		"ModuleTypeUnknown": 0,
		"ModelTypeApi":      1,
		"ModelTypeMenu":     2,
		"ModelTypeRole":     3,
		"ModelTypeUser":     4,
		"ModelTypeDict":     5,
		"ModelTypeConfig":   6,
		"ModelTypeLog":      7,
		"ModelTypeJob":      8,
		"ModelTypeNotify":   9,
		"ModelTypeSystem":   10,
		"ModelTypeMonitor":  11,
	}
)

func (x ModuleType) Enum() *ModuleType {
	p := new(ModuleType)
	*p = x
	return p
}

func (x ModuleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModuleType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[6].Descriptor()
}

func (ModuleType) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[6]
}

func (x ModuleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ModuleType.Descriptor instead.
func (ModuleType) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{6}
}

// 数据源类型
type DatasourceType int32

const (
	// 未知数据源类型
	DatasourceType_DatasourceTypeUnknown DatasourceType = 0
	// Metric
	DatasourceType_DatasourceTypeMetric DatasourceType = 1
	// Trace
	DatasourceType_DatasourceTypeTrace DatasourceType = 2
	// Log
	DatasourceType_DatasourceTypeLog DatasourceType = 3
)

// Enum value maps for DatasourceType.
var (
	DatasourceType_name = map[int32]string{
		0: "DatasourceTypeUnknown",
		1: "DatasourceTypeMetric",
		2: "DatasourceTypeTrace",
		3: "DatasourceTypeLog",
	}
	DatasourceType_value = map[string]int32{
		"DatasourceTypeUnknown": 0,
		"DatasourceTypeMetric":  1,
		"DatasourceTypeTrace":   2,
		"DatasourceTypeLog":     3,
	}
)

func (x DatasourceType) Enum() *DatasourceType {
	p := new(DatasourceType)
	*p = x
	return p
}

func (x DatasourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatasourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[7].Descriptor()
}

func (DatasourceType) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[7]
}

func (x DatasourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DatasourceType.Descriptor instead.
func (DatasourceType) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{7}
}

// 存储器类型 prometheus、victoriametrics等
type StorageType int32

const (
	// 未知存储器类型
	StorageType_StorageTypeUnknown StorageType = 0
	// Prometheus
	StorageType_StorageTypePrometheus StorageType = 1
	// VictoriaMetrics
	StorageType_StorageTypeVictoriaMetrics StorageType = 2
)

// Enum value maps for StorageType.
var (
	StorageType_name = map[int32]string{
		0: "StorageTypeUnknown",
		1: "StorageTypePrometheus",
		2: "StorageTypeVictoriaMetrics",
	}
	StorageType_value = map[string]int32{
		"StorageTypeUnknown":         0,
		"StorageTypePrometheus":      1,
		"StorageTypeVictoriaMetrics": 2,
	}
)

func (x StorageType) Enum() *StorageType {
	p := new(StorageType)
	*p = x
	return p
}

func (x StorageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[8].Descriptor()
}

func (StorageType) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[8]
}

func (x StorageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageType.Descriptor instead.
func (StorageType) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{8}
}

// MetricType 指标类型
type MetricType int32

const (
	// 未知指标类型
	MetricType_MetricTypeUnknown MetricType = 0
	// Counter
	MetricType_MetricTypeCounter MetricType = 1
	// Gauge
	MetricType_MetricTypeGauge MetricType = 2
	// Histogram
	MetricType_MetricTypeHistogram MetricType = 3
	// Summary
	MetricType_MetricTypeSummary MetricType = 4
)

// Enum value maps for MetricType.
var (
	MetricType_name = map[int32]string{
		0: "MetricTypeUnknown",
		1: "MetricTypeCounter",
		2: "MetricTypeGauge",
		3: "MetricTypeHistogram",
		4: "MetricTypeSummary",
	}
	MetricType_value = map[string]int32{
		"MetricTypeUnknown":   0,
		"MetricTypeCounter":   1,
		"MetricTypeGauge":     2,
		"MetricTypeHistogram": 3,
		"MetricTypeSummary":   4,
	}
)

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[9].Descriptor()
}

func (MetricType) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[9]
}

func (x MetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricType.Descriptor instead.
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{9}
}

// 分类, 区分字典中的各个模块数据
type DictType int32

const (
	// UNKNOWN 未知, 用于默认值
	DictType_CATEGORY_UNKNOWN DictType = 0
	// PromLabel 标签
	DictType_CATEGORY_PROM_LABEL DictType = 1
	// PromAnnotation 注解
	DictType_CATEGORY_PROM_ANNOTATION DictType = 2
	// PromStrategy 策略
	DictType_CATEGORY_PROM_STRATEGY DictType = 3
	// PromStrategyGroup 策略组
	DictType_CATEGORY_PROM_STRATEGY_GROUP DictType = 4
	// AlarmLevel 告警级别
	DictType_CATEGORY_ALARM_LEVEL DictType = 5
	// AlarmStatus 告警状态
	DictType_CATEGORY_ALARM_STATUS DictType = 6
	// NotifyType 通知类型
	DictType_CATEGORY_NOTIFY_TYPE DictType = 7
)

// Enum value maps for DictType.
var (
	DictType_name = map[int32]string{
		0: "CATEGORY_UNKNOWN",
		1: "CATEGORY_PROM_LABEL",
		2: "CATEGORY_PROM_ANNOTATION",
		3: "CATEGORY_PROM_STRATEGY",
		4: "CATEGORY_PROM_STRATEGY_GROUP",
		5: "CATEGORY_ALARM_LEVEL",
		6: "CATEGORY_ALARM_STATUS",
		7: "CATEGORY_NOTIFY_TYPE",
	}
	DictType_value = map[string]int32{
		"CATEGORY_UNKNOWN":             0,
		"CATEGORY_PROM_LABEL":          1,
		"CATEGORY_PROM_ANNOTATION":     2,
		"CATEGORY_PROM_STRATEGY":       3,
		"CATEGORY_PROM_STRATEGY_GROUP": 4,
		"CATEGORY_ALARM_LEVEL":         5,
		"CATEGORY_ALARM_STATUS":        6,
		"CATEGORY_NOTIFY_TYPE":         7,
	}
)

func (x DictType) Enum() *DictType {
	p := new(DictType)
	*p = x
	return p
}

func (x DictType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DictType) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[10].Descriptor()
}

func (DictType) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[10]
}

func (x DictType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DictType.Descriptor instead.
func (DictType) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{10}
}

// 枚举类型
type EnumItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 枚举值
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	// 枚举描述
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *EnumItem) Reset() {
	*x = EnumItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumItem) ProtoMessage() {}

func (x *EnumItem) ProtoReflect() protoreflect.Message {
	mi := &file_enum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumItem.ProtoReflect.Descriptor instead.
func (*EnumItem) Descriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{0}
}

func (x *EnumItem) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *EnumItem) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_enum_proto protoreflect.FileDescriptor

var file_enum_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x22, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x75, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2a, 0x3f, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x4c,
	0x4c, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x4e,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x2a, 0x3c, 0x0a, 0x06, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x41,
	0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4d,
	0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f,
	0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x2a, 0x4a, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x15,
	0x0a, 0x11, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x45, 0x52, 0x5f, 0x41, 0x44,
	0x4d, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x44,
	0x4d, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x10, 0x03, 0x2a, 0x50, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x16,
	0x0a, 0x12, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x52, 0x45,
	0x4d, 0x4f, 0x54, 0x45, 0x10, 0x02, 0x2a, 0x3d, 0x0a, 0x0b, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x41, 0x50, 0x54, 0x43, 0x48, 0x41,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x43, 0x41, 0x50, 0x54, 0x43, 0x48, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55,
	0x44, 0x49, 0x4f, 0x10, 0x01, 0x2a, 0x50, 0x0a, 0x0a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x10, 0x02, 0x2a, 0xfa, 0x01, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x41, 0x70, 0x69, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x65, 0x6e, 0x75,
	0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x55, 0x73, 0x65, 0x72, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x44, 0x69, 0x63, 0x74, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x10, 0x06,
	0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x6f, 0x67,
	0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x4a,
	0x6f, 0x62, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x10, 0x0a, 0x12, 0x14,
	0x0a, 0x10, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x10, 0x0b, 0x2a, 0x75, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x6f, 0x67, 0x10, 0x03, 0x2a, 0x60, 0x0a, 0x0b, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x10, 0x01, 0x12, 0x1e, 0x0a,
	0x1a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x56, 0x69, 0x63, 0x74,
	0x6f, 0x72, 0x69, 0x61, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x10, 0x02, 0x2a, 0x7f, 0x0a,
	0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x47, 0x61, 0x75, 0x67, 0x65, 0x10, 0x02, 0x12, 0x17,
	0x0a, 0x13, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x10, 0x04, 0x2a, 0xe4,
	0x01, 0x0a, 0x08, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x52,
	0x4f, 0x4d, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x5f, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45,
	0x47, 0x59, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x41, 0x4c, 0x41, 0x52, 0x4d, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x05,
	0x12, 0x19, 0x0a, 0x15, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x41, 0x4c, 0x41,
	0x52, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x10, 0x07, 0x42, 0x2c, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x50, 0x01, 0x5a, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x64, 0x65, 0x2d,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2f, 0x6d, 0x6f, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x3b,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enum_proto_rawDescOnce sync.Once
	file_enum_proto_rawDescData = file_enum_proto_rawDesc
)

func file_enum_proto_rawDescGZIP() []byte {
	file_enum_proto_rawDescOnce.Do(func() {
		file_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_enum_proto_rawDescData)
	})
	return file_enum_proto_rawDescData
}

var file_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 11)
var file_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enum_proto_goTypes = []any{
	(Status)(0),         // 0: api.Status
	(Gender)(0),         // 1: api.Gender
	(Role)(0),           // 2: api.Role
	(DataSource)(0),     // 3: api.DataSource
	(CaptchaType)(0),    // 4: api.CaptchaType
	(DomainType)(0),     // 5: api.DomainType
	(ModuleType)(0),     // 6: api.ModuleType
	(DatasourceType)(0), // 7: api.DatasourceType
	(StorageType)(0),    // 8: api.StorageType
	(MetricType)(0),     // 9: api.MetricType
	(DictType)(0),       // 10: api.DictType
	(*EnumItem)(nil),    // 11: api.EnumItem
}
var file_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enum_proto_init() }
func file_enum_proto_init() {
	if File_enum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enum_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EnumItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enum_proto_rawDesc,
			NumEnums:      11,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enum_proto_goTypes,
		DependencyIndexes: file_enum_proto_depIdxs,
		EnumInfos:         file_enum_proto_enumTypes,
		MessageInfos:      file_enum_proto_msgTypes,
	}.Build()
	File_enum_proto = out.File
	file_enum_proto_rawDesc = nil
	file_enum_proto_goTypes = nil
	file_enum_proto_depIdxs = nil
}
