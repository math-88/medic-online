// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: medical.proto

package v1

import (
	_ "github.com/ysugimoto/grpc-graphql-gateway/graphql"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Role int32

const (
	Role_user  Role = 0
	Role_admin Role = 1
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "user",
		1: "admin",
	}
	Role_value = map[string]int32{
		"user":  0,
		"admin": 1,
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
	return file_medical_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_medical_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_medical_proto_rawDescGZIP(), []int{0}
}

// The request message containing the login and password
type UserGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserGetRequest) Reset() {
	*x = UserGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetRequest) ProtoMessage() {}

func (x *UserGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_medical_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetRequest.ProtoReflect.Descriptor instead.
func (*UserGetRequest) Descriptor() ([]byte, []int) {
	return file_medical_proto_rawDescGZIP(), []int{0}
}

func (x *UserGetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// The response message containing the answer
type UserGetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserGetReply) Reset() {
	*x = UserGetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetReply) ProtoMessage() {}

func (x *UserGetReply) ProtoReflect() protoreflect.Message {
	mi := &file_medical_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetReply.ProtoReflect.Descriptor instead.
func (*UserGetReply) Descriptor() ([]byte, []int) {
	return file_medical_proto_rawDescGZIP(), []int{1}
}

func (x *UserGetReply) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname   string                 `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Email     string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Role      Role                   `protobuf:"varint,7,opt,name=role,proto3,enum=v1.Role" json:"role,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_medical_proto_msgTypes[2]
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
	return file_medical_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_user
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// The request message containing the login and password
type ProtocolGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string     `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Anamnesis *Anamnesis `protobuf:"bytes,2,opt,name=anamnesis,proto3" json:"anamnesis,omitempty"`
}

func (x *ProtocolGetRequest) Reset() {
	*x = ProtocolGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolGetRequest) ProtoMessage() {}

func (x *ProtocolGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_medical_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolGetRequest.ProtoReflect.Descriptor instead.
func (*ProtocolGetRequest) Descriptor() ([]byte, []int) {
	return file_medical_proto_rawDescGZIP(), []int{3}
}

func (x *ProtocolGetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ProtocolGetRequest) GetAnamnesis() *Anamnesis {
	if x != nil {
		return x.Anamnesis
	}
	return nil
}

// The response message containing the protocol
type ProtocolGetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol *Protocol `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *ProtocolGetReply) Reset() {
	*x = ProtocolGetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolGetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolGetReply) ProtoMessage() {}

func (x *ProtocolGetReply) ProtoReflect() protoreflect.Message {
	mi := &file_medical_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolGetReply.ProtoReflect.Descriptor instead.
func (*ProtocolGetReply) Descriptor() ([]byte, []int) {
	return file_medical_proto_rawDescGZIP(), []int{4}
}

func (x *ProtocolGetReply) GetProtocol() *Protocol {
	if x != nil {
		return x.Protocol
	}
	return nil
}

type Protocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Actions     []*Actions `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *Protocol) Reset() {
	*x = Protocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Protocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Protocol) ProtoMessage() {}

func (x *Protocol) ProtoReflect() protoreflect.Message {
	mi := &file_medical_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Protocol.ProtoReflect.Descriptor instead.
func (*Protocol) Descriptor() ([]byte, []int) {
	return file_medical_proto_rawDescGZIP(), []int{5}
}

func (x *Protocol) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Protocol) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Protocol) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Protocol) GetActions() []*Actions {
	if x != nil {
		return x.Actions
	}
	return nil
}

type Actions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// string name = 2;
	Drug        string `protobuf:"bytes,3,opt,name=drug,proto3" json:"drug,omitempty"`
	Dosage      string `protobuf:"bytes,4,opt,name=dosage,proto3" json:"dosage,omitempty"`
	Time        string `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	Duration    string `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Actions) Reset() {
	*x = Actions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actions) ProtoMessage() {}

func (x *Actions) ProtoReflect() protoreflect.Message {
	mi := &file_medical_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actions.ProtoReflect.Descriptor instead.
func (*Actions) Descriptor() ([]byte, []int) {
	return file_medical_proto_rawDescGZIP(), []int{6}
}

func (x *Actions) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Actions) GetDrug() string {
	if x != nil {
		return x.Drug
	}
	return ""
}

func (x *Actions) GetDosage() string {
	if x != nil {
		return x.Dosage
	}
	return ""
}

func (x *Actions) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Actions) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Actions) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Anamnesis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Anamnesis) Reset() {
	*x = Anamnesis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Anamnesis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Anamnesis) ProtoMessage() {}

func (x *Anamnesis) ProtoReflect() protoreflect.Message {
	mi := &file_medical_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Anamnesis.ProtoReflect.Descriptor instead.
func (*Anamnesis) Descriptor() ([]byte, []int) {
	return file_medical_proto_rawDescGZIP(), []int{7}
}

func (x *Anamnesis) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_medical_proto protoreflect.FileDescriptor

var file_medical_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0e, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0xed, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1c, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e,
	0x76, 0x31, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x5a, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x09, 0x61, 0x6e, 0x61, 0x6d, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6d, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x52, 0x09, 0x61, 0x6e, 0x61, 0x6d, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x22, 0x3c,
	0x0a, 0x10, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x28, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x77, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x72, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x72, 0x75, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x2d, 0x0a, 0x09, 0x41, 0x6e, 0x61, 0x6d, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x1b,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x01, 0x32, 0xd8, 0x01, 0x0a, 0x0e,
	0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58,
	0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x27, 0xba, 0x43, 0x09, 0x12, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x6c, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2f, 0xba, 0x43, 0x0d, 0x12, 0x0b, 0x67, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a,
	0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x14, 0x5a, 0x12, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_medical_proto_rawDescOnce sync.Once
	file_medical_proto_rawDescData = file_medical_proto_rawDesc
)

func file_medical_proto_rawDescGZIP() []byte {
	file_medical_proto_rawDescOnce.Do(func() {
		file_medical_proto_rawDescData = protoimpl.X.CompressGZIP(file_medical_proto_rawDescData)
	})
	return file_medical_proto_rawDescData
}

var file_medical_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_medical_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_medical_proto_goTypes = []interface{}{
	(Role)(0),                     // 0: v1.role
	(*UserGetRequest)(nil),        // 1: v1.UserGetRequest
	(*UserGetReply)(nil),          // 2: v1.UserGetReply
	(*User)(nil),                  // 3: v1.User
	(*ProtocolGetRequest)(nil),    // 4: v1.ProtocolGetRequest
	(*ProtocolGetReply)(nil),      // 5: v1.ProtocolGetReply
	(*Protocol)(nil),              // 6: v1.Protocol
	(*Actions)(nil),               // 7: v1.Actions
	(*Anamnesis)(nil),             // 8: v1.Anamnesis
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_medical_proto_depIdxs = []int32{
	3, // 0: v1.UserGetReply.user:type_name -> v1.User
	0, // 1: v1.User.role:type_name -> v1.role
	9, // 2: v1.User.createdAt:type_name -> google.protobuf.Timestamp
	9, // 3: v1.User.updated_at:type_name -> google.protobuf.Timestamp
	8, // 4: v1.ProtocolGetRequest.anamnesis:type_name -> v1.Anamnesis
	6, // 5: v1.ProtocolGetReply.protocol:type_name -> v1.Protocol
	7, // 6: v1.Protocol.actions:type_name -> v1.Actions
	1, // 7: v1.MedicalService.UserGet:input_type -> v1.UserGetRequest
	4, // 8: v1.MedicalService.ProtocolGet:input_type -> v1.ProtocolGetRequest
	2, // 9: v1.MedicalService.UserGet:output_type -> v1.UserGetReply
	5, // 10: v1.MedicalService.ProtocolGet:output_type -> v1.ProtocolGetReply
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_medical_proto_init() }
func file_medical_proto_init() {
	if File_medical_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_medical_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetRequest); i {
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
		file_medical_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetReply); i {
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
		file_medical_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_medical_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolGetRequest); i {
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
		file_medical_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolGetReply); i {
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
		file_medical_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Protocol); i {
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
		file_medical_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actions); i {
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
		file_medical_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Anamnesis); i {
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
			RawDescriptor: file_medical_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_medical_proto_goTypes,
		DependencyIndexes: file_medical_proto_depIdxs,
		EnumInfos:         file_medical_proto_enumTypes,
		MessageInfos:      file_medical_proto_msgTypes,
	}.Build()
	File_medical_proto = out.File
	file_medical_proto_rawDesc = nil
	file_medical_proto_goTypes = nil
	file_medical_proto_depIdxs = nil
}
