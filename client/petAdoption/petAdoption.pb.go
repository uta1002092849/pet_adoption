// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: petAdoption.proto

package petAdoption

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

// The pet object definition
type Pet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`           // ID of the pet
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`       // Name of the pet
	Gender  string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`   // Gender of the pet
	Age     int32  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`        // Age of the pet
	Breed   string `protobuf:"bytes,5,opt,name=breed,proto3" json:"breed,omitempty"`     // Breed of the pet
	Picture []byte `protobuf:"bytes,6,opt,name=picture,proto3" json:"picture,omitempty"` // Picture of the pet
}

func (x *Pet) Reset() {
	*x = Pet{}
	mi := &file_petAdoption_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_petAdoption_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_petAdoption_proto_rawDescGZIP(), []int{0}
}

func (x *Pet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Pet) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Pet) GetBreed() string {
	if x != nil {
		return x.Breed
	}
	return ""
}

func (x *Pet) GetPicture() []byte {
	if x != nil {
		return x.Picture
	}
	return nil
}

// Request message for inserting a pet
type InsertPetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Gender  string `protobuf:"bytes,2,opt,name=gender,proto3" json:"gender,omitempty"`
	Age     int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Breed   string `protobuf:"bytes,4,opt,name=breed,proto3" json:"breed,omitempty"`
	Picture []byte `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture,omitempty"`
}

func (x *InsertPetRequest) Reset() {
	*x = InsertPetRequest{}
	mi := &file_petAdoption_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertPetRequest) ProtoMessage() {}

func (x *InsertPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petAdoption_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertPetRequest.ProtoReflect.Descriptor instead.
func (*InsertPetRequest) Descriptor() ([]byte, []int) {
	return file_petAdoption_proto_rawDescGZIP(), []int{1}
}

func (x *InsertPetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InsertPetRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *InsertPetRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *InsertPetRequest) GetBreed() string {
	if x != nil {
		return x.Breed
	}
	return ""
}

func (x *InsertPetRequest) GetPicture() []byte {
	if x != nil {
		return x.Picture
	}
	return nil
}

// Response message for insert operations
type InsertPetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *InsertPetResponse) Reset() {
	*x = InsertPetResponse{}
	mi := &file_petAdoption_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertPetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertPetResponse) ProtoMessage() {}

func (x *InsertPetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petAdoption_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertPetResponse.ProtoReflect.Descriptor instead.
func (*InsertPetResponse) Descriptor() ([]byte, []int) {
	return file_petAdoption_proto_rawDescGZIP(), []int{2}
}

func (x *InsertPetResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Request message for searching pets
type SearchPetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchTerm string `protobuf:"bytes,1,opt,name=searchTerm,proto3" json:"searchTerm,omitempty"`
}

func (x *SearchPetRequest) Reset() {
	*x = SearchPetRequest{}
	mi := &file_petAdoption_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPetRequest) ProtoMessage() {}

func (x *SearchPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petAdoption_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPetRequest.ProtoReflect.Descriptor instead.
func (*SearchPetRequest) Descriptor() ([]byte, []int) {
	return file_petAdoption_proto_rawDescGZIP(), []int{3}
}

func (x *SearchPetRequest) GetSearchTerm() string {
	if x != nil {
		return x.SearchTerm
	}
	return ""
}

// Response message for search results
type SearchPetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pets []*Pet `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
}

func (x *SearchPetResponse) Reset() {
	*x = SearchPetResponse{}
	mi := &file_petAdoption_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchPetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPetResponse) ProtoMessage() {}

func (x *SearchPetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_petAdoption_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPetResponse.ProtoReflect.Descriptor instead.
func (*SearchPetResponse) Descriptor() ([]byte, []int) {
	return file_petAdoption_proto_rawDescGZIP(), []int{4}
}

func (x *SearchPetResponse) GetPets() []*Pet {
	if x != nil {
		return x.Pets
	}
	return nil
}

var File_petAdoption_proto protoreflect.FileDescriptor

var file_petAdoption_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x65, 0x74, 0x41, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x65, 0x74, 0x41, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x83, 0x01, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x65, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x65, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x65,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x65, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x32, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x39, 0x0a, 0x11,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x65, 0x74, 0x41, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x65,
	0x74, 0x52, 0x04, 0x70, 0x65, 0x74, 0x73, 0x32, 0xa5, 0x01, 0x0a, 0x0b, 0x50, 0x65, 0x74, 0x41,
	0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x50, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x65, 0x74, 0x41, 0x64, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x65, 0x74, 0x41, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x74,
	0x12, 0x1d, 0x2e, 0x70, 0x65, 0x74, 0x41, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x70, 0x65, 0x74, 0x41, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x65, 0x74, 0x41, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_petAdoption_proto_rawDescOnce sync.Once
	file_petAdoption_proto_rawDescData = file_petAdoption_proto_rawDesc
)

func file_petAdoption_proto_rawDescGZIP() []byte {
	file_petAdoption_proto_rawDescOnce.Do(func() {
		file_petAdoption_proto_rawDescData = protoimpl.X.CompressGZIP(file_petAdoption_proto_rawDescData)
	})
	return file_petAdoption_proto_rawDescData
}

var file_petAdoption_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_petAdoption_proto_goTypes = []any{
	(*Pet)(nil),               // 0: petAdoption.Pet
	(*InsertPetRequest)(nil),  // 1: petAdoption.InsertPetRequest
	(*InsertPetResponse)(nil), // 2: petAdoption.InsertPetResponse
	(*SearchPetRequest)(nil),  // 3: petAdoption.SearchPetRequest
	(*SearchPetResponse)(nil), // 4: petAdoption.SearchPetResponse
}
var file_petAdoption_proto_depIdxs = []int32{
	0, // 0: petAdoption.SearchPetResponse.pets:type_name -> petAdoption.Pet
	1, // 1: petAdoption.PetAdoption.InsertPet:input_type -> petAdoption.InsertPetRequest
	3, // 2: petAdoption.PetAdoption.SearchPet:input_type -> petAdoption.SearchPetRequest
	2, // 3: petAdoption.PetAdoption.InsertPet:output_type -> petAdoption.InsertPetResponse
	4, // 4: petAdoption.PetAdoption.SearchPet:output_type -> petAdoption.SearchPetResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_petAdoption_proto_init() }
func file_petAdoption_proto_init() {
	if File_petAdoption_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_petAdoption_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_petAdoption_proto_goTypes,
		DependencyIndexes: file_petAdoption_proto_depIdxs,
		MessageInfos:      file_petAdoption_proto_msgTypes,
	}.Build()
	File_petAdoption_proto = out.File
	file_petAdoption_proto_rawDesc = nil
	file_petAdoption_proto_goTypes = nil
	file_petAdoption_proto_depIdxs = nil
}
