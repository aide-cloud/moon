// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: admin/model.proto

package admin

import (
	api "github.com/aide-cloud/moon/api"
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

type SelectExtend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 图标
	Icon string `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`
	// 颜色
	Color string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	// 描述
	Remark string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	// 图片URL
	Image string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *SelectExtend) Reset() {
	*x = SelectExtend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectExtend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectExtend) ProtoMessage() {}

func (x *SelectExtend) ProtoReflect() protoreflect.Message {
	mi := &file_admin_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectExtend.ProtoReflect.Descriptor instead.
func (*SelectExtend) Descriptor() ([]byte, []int) {
	return file_admin_model_proto_rawDescGZIP(), []int{0}
}

func (x *SelectExtend) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *SelectExtend) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *SelectExtend) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *SelectExtend) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

// 下拉选择基础数据
type Select struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数据值
	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	// 数据label
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// 子级数据, 针对级联选择
	Children []*Select `protobuf:"bytes,3,rep,name=children,proto3" json:"children,omitempty"`
	// 是否禁用
	Disabled bool `protobuf:"varint,4,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// 针对有图标类型的配置项(可选，默认为null)
	Extend *SelectExtend `protobuf:"bytes,5,opt,name=extend,proto3,oneof" json:"extend,omitempty"`
}

func (x *Select) Reset() {
	*x = Select{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Select) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Select) ProtoMessage() {}

func (x *Select) ProtoReflect() protoreflect.Message {
	mi := &file_admin_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Select.ProtoReflect.Descriptor instead.
func (*Select) Descriptor() ([]byte, []int) {
	return file_admin_model_proto_rawDescGZIP(), []int{1}
}

func (x *Select) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Select) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Select) GetChildren() []*Select {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Select) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Select) GetExtend() *SelectExtend {
	if x != nil {
		return x.Extend
	}
	return nil
}

// 用户模块
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 用户名
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 昵称
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// 邮箱
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// 手机
	Phone string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	// 状态
	Status api.Status `protobuf:"varint,6,opt,name=status,proto3,enum=api.Status" json:"status,omitempty"`
	// 性别
	Gender api.Gender `protobuf:"varint,7,opt,name=gender,proto3,enum=api.Gender" json:"gender,omitempty"`
	// 角色
	Role api.Role `protobuf:"varint,8,opt,name=role,proto3,enum=api.Role" json:"role,omitempty"`
	// 头像
	Avatar string `protobuf:"bytes,9,opt,name=avatar,proto3" json:"avatar,omitempty"`
	// 个人说明
	Remark string `protobuf:"bytes,10,opt,name=remark,proto3" json:"remark,omitempty"`
	// 创建时间
	CreatedAt string `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt string `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_admin_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_admin_model_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetStatus() api.Status {
	if x != nil {
		return x.Status
	}
	return api.Status(0)
}

func (x *User) GetGender() api.Gender {
	if x != nil {
		return x.Gender
	}
	return api.Gender(0)
}

func (x *User) GetRole() api.Role {
	if x != nil {
		return x.Role
	}
	return api.Role(0)
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *User) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// 系统API资源模块
type ResourceItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 接口ID
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 接口名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 接口路径
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// 接口状态
	Status api.Status `protobuf:"varint,4,opt,name=status,proto3,enum=api.Status" json:"status,omitempty"`
	// 备注
	Remark string `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	// 创建时间
	CreatedAt string `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// 更新时间
	UpdatedAt string `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// 删除时间
	DeletedAt string `protobuf:"bytes,8,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	// 所属模块
	Module api.ModuleType `protobuf:"varint,9,opt,name=module,proto3,enum=api.ModuleType" json:"module,omitempty"`
	// 所属领域
	Domain api.DomainType `protobuf:"varint,10,opt,name=domain,proto3,enum=api.DomainType" json:"domain,omitempty"`
}

func (x *ResourceItem) Reset() {
	*x = ResourceItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceItem) ProtoMessage() {}

func (x *ResourceItem) ProtoReflect() protoreflect.Message {
	mi := &file_admin_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceItem.ProtoReflect.Descriptor instead.
func (*ResourceItem) Descriptor() ([]byte, []int) {
	return file_admin_model_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceItem) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResourceItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceItem) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ResourceItem) GetStatus() api.Status {
	if x != nil {
		return x.Status
	}
	return api.Status(0)
}

func (x *ResourceItem) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ResourceItem) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ResourceItem) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ResourceItem) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *ResourceItem) GetModule() api.ModuleType {
	if x != nil {
		return x.Module
	}
	return api.ModuleType(0)
}

func (x *ResourceItem) GetDomain() api.DomainType {
	if x != nil {
		return x.Domain
	}
	return api.DomainType(0)
}

var File_admin_model_proto protoreflect.FileDescriptor

var file_admin_model_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x0a,
	0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x0c, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x08,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x48, 0x00, 0x52,
	0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x22, 0xc9, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x23, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xaf, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x27, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x42, 0x39, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x69, 0x64, 0x65, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x6f, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_model_proto_rawDescOnce sync.Once
	file_admin_model_proto_rawDescData = file_admin_model_proto_rawDesc
)

func file_admin_model_proto_rawDescGZIP() []byte {
	file_admin_model_proto_rawDescOnce.Do(func() {
		file_admin_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_model_proto_rawDescData)
	})
	return file_admin_model_proto_rawDescData
}

var file_admin_model_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_admin_model_proto_goTypes = []interface{}{
	(*SelectExtend)(nil), // 0: api.admin.SelectExtend
	(*Select)(nil),       // 1: api.admin.Select
	(*User)(nil),         // 2: api.admin.User
	(*ResourceItem)(nil), // 3: api.admin.ResourceItem
	(api.Status)(0),      // 4: api.Status
	(api.Gender)(0),      // 5: api.Gender
	(api.Role)(0),        // 6: api.Role
	(api.ModuleType)(0),  // 7: api.ModuleType
	(api.DomainType)(0),  // 8: api.DomainType
}
var file_admin_model_proto_depIdxs = []int32{
	1, // 0: api.admin.Select.children:type_name -> api.admin.Select
	0, // 1: api.admin.Select.extend:type_name -> api.admin.SelectExtend
	4, // 2: api.admin.User.status:type_name -> api.Status
	5, // 3: api.admin.User.gender:type_name -> api.Gender
	6, // 4: api.admin.User.role:type_name -> api.Role
	4, // 5: api.admin.ResourceItem.status:type_name -> api.Status
	7, // 6: api.admin.ResourceItem.module:type_name -> api.ModuleType
	8, // 7: api.admin.ResourceItem.domain:type_name -> api.DomainType
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_admin_model_proto_init() }
func file_admin_model_proto_init() {
	if File_admin_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectExtend); i {
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
		file_admin_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Select); i {
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
		file_admin_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_admin_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceItem); i {
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
	file_admin_model_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_model_proto_goTypes,
		DependencyIndexes: file_admin_model_proto_depIdxs,
		MessageInfos:      file_admin_model_proto_msgTypes,
	}.Build()
	File_admin_model_proto = out.File
	file_admin_model_proto_rawDesc = nil
	file_admin_model_proto_goTypes = nil
	file_admin_model_proto_depIdxs = nil
}
