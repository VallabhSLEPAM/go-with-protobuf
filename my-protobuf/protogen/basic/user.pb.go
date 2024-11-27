// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: my-protobuf/proto/basic/user.proto

package basic

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Gender int32

const (
	Gender_GENDER_MALE   Gender = 0
	Gender_GENDER_FEMALE Gender = 1
	Gender_GENDER_OTHERS Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "GENDER_MALE",
		1: "GENDER_FEMALE",
		2: "GENDER_OTHERS",
	}
	Gender_value = map[string]int32{
		"GENDER_MALE":   0,
		"GENDER_FEMALE": 1,
		"GENDER_OTHERS": 2,
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
	return file_my_protobuf_proto_basic_user_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_my_protobuf_proto_basic_user_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_my_protobuf_proto_basic_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string     `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	IsActive             bool       `protobuf:"varint,3,opt,name=is_active,proto3" json:"is_active,omitempty"`
	Password             []byte     `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Gender               Gender     `protobuf:"varint,16,opt,name=gender,proto3,enum=Gender" json:"gender,omitempty"`
	Emails               []string   `protobuf:"bytes,17,rep,name=emails,proto3" json:"emails,omitempty"`
	Address              *Address   `protobuf:"bytes,18,opt,name=address,proto3" json:"address,omitempty"`
	CommunicationChannel *anypb.Any `protobuf:"bytes,19,opt,name=communication_channel,proto3" json:"communication_channel,omitempty"`
	// Types that are assignable to ElectronicCommChannel:
	//
	//	*User_SocialMedia
	//	*User_InstantMessaging
	ElectronicCommChannel isUser_ElectronicCommChannel `protobuf_oneof:"electronic_comm_channel"`
	SkillRating           map[string]uint32            `protobuf:"bytes,22,rep,name=skill_rating,proto3" json:"skill_rating,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_my_protobuf_proto_basic_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_my_protobuf_proto_basic_user_proto_msgTypes[0]
	if x != nil {
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
	return file_my_protobuf_proto_basic_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *User) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *User) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_GENDER_MALE
}

func (x *User) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *User) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *User) GetCommunicationChannel() *anypb.Any {
	if x != nil {
		return x.CommunicationChannel
	}
	return nil
}

func (m *User) GetElectronicCommChannel() isUser_ElectronicCommChannel {
	if m != nil {
		return m.ElectronicCommChannel
	}
	return nil
}

func (x *User) GetSocialMedia() *SocialMediaPlatform {
	if x, ok := x.GetElectronicCommChannel().(*User_SocialMedia); ok {
		return x.SocialMedia
	}
	return nil
}

func (x *User) GetInstantMessaging() *InstantMessagingPlatform {
	if x, ok := x.GetElectronicCommChannel().(*User_InstantMessaging); ok {
		return x.InstantMessaging
	}
	return nil
}

func (x *User) GetSkillRating() map[string]uint32 {
	if x != nil {
		return x.SkillRating
	}
	return nil
}

type isUser_ElectronicCommChannel interface {
	isUser_ElectronicCommChannel()
}

type User_SocialMedia struct {
	SocialMedia *SocialMediaPlatform `protobuf:"bytes,20,opt,name=socialMedia,json=social_media,proto3,oneof"`
}

type User_InstantMessaging struct {
	InstantMessaging *InstantMessagingPlatform `protobuf:"bytes,21,opt,name=InstantMessaging,json=instant_messaging,proto3,oneof"`
}

func (*User_SocialMedia) isUser_ElectronicCommChannel() {}

func (*User_InstantMessaging) isUser_ElectronicCommChannel() {}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street     string              `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City       string              `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country    string              `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Coordinate *Address_Coordinate `protobuf:"bytes,16,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_my_protobuf_proto_basic_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_my_protobuf_proto_basic_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_my_protobuf_proto_basic_user_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetCoordinate() *Address_Coordinate {
	if x != nil {
		return x.Coordinate
	}
	return nil
}

type PaperMail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaperMailAddress string `protobuf:"bytes,1,opt,name=paper_mail_address,proto3" json:"paper_mail_address,omitempty"`
}

func (x *PaperMail) Reset() {
	*x = PaperMail{}
	mi := &file_my_protobuf_proto_basic_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaperMail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaperMail) ProtoMessage() {}

func (x *PaperMail) ProtoReflect() protoreflect.Message {
	mi := &file_my_protobuf_proto_basic_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaperMail.ProtoReflect.Descriptor instead.
func (*PaperMail) Descriptor() ([]byte, []int) {
	return file_my_protobuf_proto_basic_user_proto_rawDescGZIP(), []int{2}
}

func (x *PaperMail) GetPaperMailAddress() string {
	if x != nil {
		return x.PaperMailAddress
	}
	return ""
}

type SocialMediaPlatform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SocialMediaPlatform string `protobuf:"bytes,1,opt,name=social_media_platform,proto3" json:"social_media_platform,omitempty"`
	SocialMediaUsername string `protobuf:"bytes,2,opt,name=social_media_username,proto3" json:"social_media_username,omitempty"`
}

func (x *SocialMediaPlatform) Reset() {
	*x = SocialMediaPlatform{}
	mi := &file_my_protobuf_proto_basic_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SocialMediaPlatform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialMediaPlatform) ProtoMessage() {}

func (x *SocialMediaPlatform) ProtoReflect() protoreflect.Message {
	mi := &file_my_protobuf_proto_basic_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialMediaPlatform.ProtoReflect.Descriptor instead.
func (*SocialMediaPlatform) Descriptor() ([]byte, []int) {
	return file_my_protobuf_proto_basic_user_proto_rawDescGZIP(), []int{3}
}

func (x *SocialMediaPlatform) GetSocialMediaPlatform() string {
	if x != nil {
		return x.SocialMediaPlatform
	}
	return ""
}

func (x *SocialMediaPlatform) GetSocialMediaUsername() string {
	if x != nil {
		return x.SocialMediaUsername
	}
	return ""
}

type InstantMessagingPlatform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstantMessagingPlatform string `protobuf:"bytes,1,opt,name=instant_messaging_platform,proto3" json:"instant_messaging_platform,omitempty"`
	InstantMessagingUsername string `protobuf:"bytes,2,opt,name=instant_messaging_username,proto3" json:"instant_messaging_username,omitempty"`
}

func (x *InstantMessagingPlatform) Reset() {
	*x = InstantMessagingPlatform{}
	mi := &file_my_protobuf_proto_basic_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstantMessagingPlatform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantMessagingPlatform) ProtoMessage() {}

func (x *InstantMessagingPlatform) ProtoReflect() protoreflect.Message {
	mi := &file_my_protobuf_proto_basic_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantMessagingPlatform.ProtoReflect.Descriptor instead.
func (*InstantMessagingPlatform) Descriptor() ([]byte, []int) {
	return file_my_protobuf_proto_basic_user_proto_rawDescGZIP(), []int{4}
}

func (x *InstantMessagingPlatform) GetInstantMessagingPlatform() string {
	if x != nil {
		return x.InstantMessagingPlatform
	}
	return ""
}

func (x *InstantMessagingPlatform) GetInstantMessagingUsername() string {
	if x != nil {
		return x.InstantMessagingUsername
	}
	return ""
}

type Address_Coordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Address_Coordinate) Reset() {
	*x = Address_Coordinate{}
	mi := &file_my_protobuf_proto_basic_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address_Coordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address_Coordinate) ProtoMessage() {}

func (x *Address_Coordinate) ProtoReflect() protoreflect.Message {
	mi := &file_my_protobuf_proto_basic_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address_Coordinate.ProtoReflect.Descriptor instead.
func (*Address_Coordinate) Descriptor() ([]byte, []int) {
	return file_my_protobuf_proto_basic_user_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Address_Coordinate) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Address_Coordinate) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_my_protobuf_proto_basic_user_proto protoreflect.FileDescriptor

var file_my_protobuf_proto_basic_user_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb1, 0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07,
	0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4a, 0x0a, 0x15, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x39, 0x0a, 0x0b, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x12, 0x48, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x48, 0x00, 0x52, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x0a, 0x0c,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x16, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x3e, 0x0a, 0x10, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x19, 0x0a, 0x17, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x22, 0xcc, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x0a,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x1a, 0x46, 0x0a, 0x0a, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x22, 0x3b, 0x0a, 0x09, 0x70, 0x61, 0x70, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6c, 0x12,
	0x2e, 0x0a, 0x12, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x81, 0x01, 0x0a, 0x13, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x34, 0x0a,
	0x15, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x18, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x3e, 0x0a, 0x1a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x3e, 0x0a, 0x1a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x2a, 0x3f, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x45,
	0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x47,
	0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x53, 0x10,
	0x02, 0x42, 0x1c, 0x5a, 0x1a, 0x6d, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_my_protobuf_proto_basic_user_proto_rawDescOnce sync.Once
	file_my_protobuf_proto_basic_user_proto_rawDescData = file_my_protobuf_proto_basic_user_proto_rawDesc
)

func file_my_protobuf_proto_basic_user_proto_rawDescGZIP() []byte {
	file_my_protobuf_proto_basic_user_proto_rawDescOnce.Do(func() {
		file_my_protobuf_proto_basic_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_my_protobuf_proto_basic_user_proto_rawDescData)
	})
	return file_my_protobuf_proto_basic_user_proto_rawDescData
}

var file_my_protobuf_proto_basic_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_my_protobuf_proto_basic_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_my_protobuf_proto_basic_user_proto_goTypes = []any{
	(Gender)(0),                      // 0: Gender
	(*User)(nil),                     // 1: User
	(*Address)(nil),                  // 2: Address
	(*PaperMail)(nil),                // 3: paperMail
	(*SocialMediaPlatform)(nil),      // 4: socialMediaPlatform
	(*InstantMessagingPlatform)(nil), // 5: instantMessagingPlatform
	nil,                              // 6: User.SkillRatingEntry
	(*Address_Coordinate)(nil),       // 7: Address.Coordinate
	(*anypb.Any)(nil),                // 8: google.protobuf.Any
}
var file_my_protobuf_proto_basic_user_proto_depIdxs = []int32{
	0, // 0: User.gender:type_name -> Gender
	2, // 1: User.address:type_name -> Address
	8, // 2: User.communication_channel:type_name -> google.protobuf.Any
	4, // 3: User.socialMedia:type_name -> socialMediaPlatform
	5, // 4: User.InstantMessaging:type_name -> instantMessagingPlatform
	6, // 5: User.skill_rating:type_name -> User.SkillRatingEntry
	7, // 6: Address.coordinate:type_name -> Address.Coordinate
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_my_protobuf_proto_basic_user_proto_init() }
func file_my_protobuf_proto_basic_user_proto_init() {
	if File_my_protobuf_proto_basic_user_proto != nil {
		return
	}
	file_my_protobuf_proto_basic_user_proto_msgTypes[0].OneofWrappers = []any{
		(*User_SocialMedia)(nil),
		(*User_InstantMessaging)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_my_protobuf_proto_basic_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_my_protobuf_proto_basic_user_proto_goTypes,
		DependencyIndexes: file_my_protobuf_proto_basic_user_proto_depIdxs,
		EnumInfos:         file_my_protobuf_proto_basic_user_proto_enumTypes,
		MessageInfos:      file_my_protobuf_proto_basic_user_proto_msgTypes,
	}.Build()
	File_my_protobuf_proto_basic_user_proto = out.File
	file_my_protobuf_proto_basic_user_proto_rawDesc = nil
	file_my_protobuf_proto_basic_user_proto_goTypes = nil
	file_my_protobuf_proto_basic_user_proto_depIdxs = nil
}
