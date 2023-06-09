// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: staff.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Staff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TeamName     string                 `protobuf:"bytes,3,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	Organization string                 `protobuf:"bytes,4,opt,name=organization,proto3" json:"organization,omitempty"`
	Title        string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	OnboardDate  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=onboard_date,json=onboardDate,proto3" json:"onboard_date,omitempty"`
	Active       bool                   `protobuf:"varint,7,opt,name=active,proto3" json:"active,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ModifiedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
	DeletedAt    *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Staff) Reset() {
	*x = Staff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Staff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Staff) ProtoMessage() {}

func (x *Staff) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Staff.ProtoReflect.Descriptor instead.
func (*Staff) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{0}
}

func (x *Staff) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Staff) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Staff) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *Staff) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *Staff) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Staff) GetOnboardDate() *timestamppb.Timestamp {
	if x != nil {
		return x.OnboardDate
	}
	return nil
}

func (x *Staff) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Staff) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Staff) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

func (x *Staff) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type CreateStaffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TeamName     string                 `protobuf:"bytes,3,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	Organization string                 `protobuf:"bytes,4,opt,name=organization,proto3" json:"organization,omitempty"`
	Title        string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	OnboardDate  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=onboard_date,json=onboardDate,proto3" json:"onboard_date,omitempty"`
}

func (x *CreateStaffRequest) Reset() {
	*x = CreateStaffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStaffRequest) ProtoMessage() {}

func (x *CreateStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStaffRequest.ProtoReflect.Descriptor instead.
func (*CreateStaffRequest) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStaffRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateStaffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStaffRequest) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *CreateStaffRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *CreateStaffRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateStaffRequest) GetOnboardDate() *timestamppb.Timestamp {
	if x != nil {
		return x.OnboardDate
	}
	return nil
}

type CreateStaffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TeamName     string                 `protobuf:"bytes,3,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	Organization string                 `protobuf:"bytes,4,opt,name=organization,proto3" json:"organization,omitempty"`
	Title        string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	OnboardDate  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=onboard_date,json=onboardDate,proto3" json:"onboard_date,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CreateStaffResponse) Reset() {
	*x = CreateStaffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStaffResponse) ProtoMessage() {}

func (x *CreateStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStaffResponse.ProtoReflect.Descriptor instead.
func (*CreateStaffResponse) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStaffResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateStaffResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStaffResponse) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *CreateStaffResponse) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *CreateStaffResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateStaffResponse) GetOnboardDate() *timestamppb.Timestamp {
	if x != nil {
		return x.OnboardDate
	}
	return nil
}

func (x *CreateStaffResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type UpdateStaffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TeamName     string                 `protobuf:"bytes,3,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	Organization string                 `protobuf:"bytes,4,opt,name=organization,proto3" json:"organization,omitempty"`
	Title        string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	OnboardDate  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=onboard_date,json=onboardDate,proto3" json:"onboard_date,omitempty"`
}

func (x *UpdateStaffRequest) Reset() {
	*x = UpdateStaffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStaffRequest) ProtoMessage() {}

func (x *UpdateStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStaffRequest.ProtoReflect.Descriptor instead.
func (*UpdateStaffRequest) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateStaffRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateStaffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateStaffRequest) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *UpdateStaffRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *UpdateStaffRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateStaffRequest) GetOnboardDate() *timestamppb.Timestamp {
	if x != nil {
		return x.OnboardDate
	}
	return nil
}

type UpdateStaffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TeamName     string                 `protobuf:"bytes,3,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	Organization string                 `protobuf:"bytes,4,opt,name=organization,proto3" json:"organization,omitempty"`
	Title        string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	OnboardDate  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=onboard_date,json=onboardDate,proto3" json:"onboard_date,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ModifiedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
}

func (x *UpdateStaffResponse) Reset() {
	*x = UpdateStaffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStaffResponse) ProtoMessage() {}

func (x *UpdateStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStaffResponse.ProtoReflect.Descriptor instead.
func (*UpdateStaffResponse) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateStaffResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateStaffResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateStaffResponse) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *UpdateStaffResponse) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *UpdateStaffResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateStaffResponse) GetOnboardDate() *timestamppb.Timestamp {
	if x != nil {
		return x.OnboardDate
	}
	return nil
}

func (x *UpdateStaffResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UpdateStaffResponse) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

type ListStaffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Staffs []*Staff `protobuf:"bytes,1,rep,name=staffs,proto3" json:"staffs,omitempty"`
}

func (x *ListStaffResponse) Reset() {
	*x = ListStaffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaffResponse) ProtoMessage() {}

func (x *ListStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaffResponse.ProtoReflect.Descriptor instead.
func (*ListStaffResponse) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{5}
}

func (x *ListStaffResponse) GetStaffs() []*Staff {
	if x != nil {
		return x.Staffs
	}
	return nil
}

type StaffId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StaffId) Reset() {
	*x = StaffId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffId) ProtoMessage() {}

func (x *StaffId) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffId.ProtoReflect.Descriptor instead.
func (*StaffId) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{6}
}

func (x *StaffId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_staff_proto protoreflect.FileDescriptor

var file_staff_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x3d, 0x0a, 0x0c, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xce, 0x01, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0c,
	0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b,
	0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x8a, 0x02, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3d,
	0x0a, 0x0c, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xce, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xc7, 0x02, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0c,
	0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b,
	0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x33, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x66, 0x66, 0x73, 0x22, 0x19, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x32, 0x8f, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x13, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x13, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x08, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x1a, 0x06, 0x2e, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x12, 0x2f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x12, 0x08, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x68, 0x61, 0x75, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2d, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staff_proto_rawDescOnce sync.Once
	file_staff_proto_rawDescData = file_staff_proto_rawDesc
)

func file_staff_proto_rawDescGZIP() []byte {
	file_staff_proto_rawDescOnce.Do(func() {
		file_staff_proto_rawDescData = protoimpl.X.CompressGZIP(file_staff_proto_rawDescData)
	})
	return file_staff_proto_rawDescData
}

var file_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_staff_proto_goTypes = []interface{}{
	(*Staff)(nil),                 // 0: Staff
	(*CreateStaffRequest)(nil),    // 1: CreateStaffRequest
	(*CreateStaffResponse)(nil),   // 2: CreateStaffResponse
	(*UpdateStaffRequest)(nil),    // 3: UpdateStaffRequest
	(*UpdateStaffResponse)(nil),   // 4: UpdateStaffResponse
	(*ListStaffResponse)(nil),     // 5: ListStaffResponse
	(*StaffId)(nil),               // 6: StaffId
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_staff_proto_depIdxs = []int32{
	7,  // 0: Staff.onboard_date:type_name -> google.protobuf.Timestamp
	7,  // 1: Staff.created_at:type_name -> google.protobuf.Timestamp
	7,  // 2: Staff.modified_at:type_name -> google.protobuf.Timestamp
	7,  // 3: Staff.deleted_at:type_name -> google.protobuf.Timestamp
	7,  // 4: CreateStaffRequest.onboard_date:type_name -> google.protobuf.Timestamp
	7,  // 5: CreateStaffResponse.onboard_date:type_name -> google.protobuf.Timestamp
	7,  // 6: CreateStaffResponse.created_at:type_name -> google.protobuf.Timestamp
	7,  // 7: UpdateStaffRequest.onboard_date:type_name -> google.protobuf.Timestamp
	7,  // 8: UpdateStaffResponse.onboard_date:type_name -> google.protobuf.Timestamp
	7,  // 9: UpdateStaffResponse.created_at:type_name -> google.protobuf.Timestamp
	7,  // 10: UpdateStaffResponse.modified_at:type_name -> google.protobuf.Timestamp
	0,  // 11: ListStaffResponse.staffs:type_name -> Staff
	8,  // 12: StaffService.listStaff:input_type -> google.protobuf.Empty
	1,  // 13: StaffService.createStaff:input_type -> CreateStaffRequest
	3,  // 14: StaffService.updateStaff:input_type -> UpdateStaffRequest
	6,  // 15: StaffService.findStaffById:input_type -> StaffId
	6,  // 16: StaffService.deleteStaff:input_type -> StaffId
	5,  // 17: StaffService.listStaff:output_type -> ListStaffResponse
	2,  // 18: StaffService.createStaff:output_type -> CreateStaffResponse
	4,  // 19: StaffService.updateStaff:output_type -> UpdateStaffResponse
	0,  // 20: StaffService.findStaffById:output_type -> Staff
	8,  // 21: StaffService.deleteStaff:output_type -> google.protobuf.Empty
	17, // [17:22] is the sub-list for method output_type
	12, // [12:17] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_staff_proto_init() }
func file_staff_proto_init() {
	if File_staff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_staff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Staff); i {
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
		file_staff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStaffRequest); i {
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
		file_staff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStaffResponse); i {
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
		file_staff_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStaffRequest); i {
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
		file_staff_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStaffResponse); i {
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
		file_staff_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStaffResponse); i {
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
		file_staff_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffId); i {
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
			RawDescriptor: file_staff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staff_proto_goTypes,
		DependencyIndexes: file_staff_proto_depIdxs,
		MessageInfos:      file_staff_proto_msgTypes,
	}.Build()
	File_staff_proto = out.File
	file_staff_proto_rawDesc = nil
	file_staff_proto_goTypes = nil
	file_staff_proto_depIdxs = nil
}
