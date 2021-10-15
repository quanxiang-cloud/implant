// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/proto/v1alpha1/api.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateReq struct {
	ObjectMeta           *ObjectMeta `protobuf:"bytes,1,opt,name=objectMeta,proto3" json:"objectMeta,omitempty"`
	State                *Status     `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateReq) Reset()         { *m = CreateReq{} }
func (m *CreateReq) String() string { return proto.CompactTextString(m) }
func (*CreateReq) ProtoMessage()    {}
func (*CreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d227e95d8b5d27, []int{0}
}

func (m *CreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReq.Unmarshal(m, b)
}
func (m *CreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReq.Marshal(b, m, deterministic)
}
func (m *CreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReq.Merge(m, src)
}
func (m *CreateReq) XXX_Size() int {
	return xxx_messageInfo_CreateReq.Size(m)
}
func (m *CreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReq proto.InternalMessageInfo

func (m *CreateReq) GetObjectMeta() *ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *CreateReq) GetState() *Status {
	if m != nil {
		return m.State
	}
	return nil
}

type ObjectMeta struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namepsace            string            `protobuf:"bytes,2,opt,name=namepsace,proto3" json:"namepsace,omitempty"`
	Uid                  string            `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	CreationTime         string            `protobuf:"bytes,4,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	Labels               map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OwnerReference       []*OwnerReference `protobuf:"bytes,6,rep,name=ownerReference,proto3" json:"ownerReference,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ObjectMeta) Reset()         { *m = ObjectMeta{} }
func (m *ObjectMeta) String() string { return proto.CompactTextString(m) }
func (*ObjectMeta) ProtoMessage()    {}
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d227e95d8b5d27, []int{1}
}

func (m *ObjectMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectMeta.Unmarshal(m, b)
}
func (m *ObjectMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectMeta.Marshal(b, m, deterministic)
}
func (m *ObjectMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMeta.Merge(m, src)
}
func (m *ObjectMeta) XXX_Size() int {
	return xxx_messageInfo_ObjectMeta.Size(m)
}
func (m *ObjectMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMeta proto.InternalMessageInfo

func (m *ObjectMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectMeta) GetNamepsace() string {
	if m != nil {
		return m.Namepsace
	}
	return ""
}

func (m *ObjectMeta) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ObjectMeta) GetCreationTime() string {
	if m != nil {
		return m.CreationTime
	}
	return ""
}

func (m *ObjectMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ObjectMeta) GetOwnerReference() []*OwnerReference {
	if m != nil {
		return m.OwnerReference
	}
	return nil
}

type OwnerReference struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OwnerReference) Reset()         { *m = OwnerReference{} }
func (m *OwnerReference) String() string { return proto.CompactTextString(m) }
func (*OwnerReference) ProtoMessage()    {}
func (*OwnerReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d227e95d8b5d27, []int{2}
}

func (m *OwnerReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnerReference.Unmarshal(m, b)
}
func (m *OwnerReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnerReference.Marshal(b, m, deterministic)
}
func (m *OwnerReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerReference.Merge(m, src)
}
func (m *OwnerReference) XXX_Size() int {
	return xxx_messageInfo_OwnerReference.Size(m)
}
func (m *OwnerReference) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerReference.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerReference proto.InternalMessageInfo

func (m *OwnerReference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OwnerReference) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type Status struct {
	Conditions           *Condition `protobuf:"bytes,1,opt,name=conditions,proto3" json:"conditions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d227e95d8b5d27, []int{3}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetConditions() *Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type Condition struct {
	Status               string                    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Reason               string                    `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message              string                    `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	ResourceRef          map[string]*StepCondition `protobuf:"bytes,4,rep,name=resourceRef,proto3" json:"resourceRef,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LastTransitionTime   string                    `protobuf:"bytes,5,opt,name=lastTransitionTime,proto3" json:"lastTransitionTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Condition) Reset()         { *m = Condition{} }
func (m *Condition) String() string { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()    {}
func (*Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d227e95d8b5d27, []int{4}
}

func (m *Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Condition.Unmarshal(m, b)
}
func (m *Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Condition.Marshal(b, m, deterministic)
}
func (m *Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Condition.Merge(m, src)
}
func (m *Condition) XXX_Size() int {
	return xxx_messageInfo_Condition.Size(m)
}
func (m *Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Condition proto.InternalMessageInfo

func (m *Condition) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Condition) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Condition) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Condition) GetResourceRef() map[string]*StepCondition {
	if m != nil {
		return m.ResourceRef
	}
	return nil
}

func (m *Condition) GetLastTransitionTime() string {
	if m != nil {
		return m.LastTransitionTime
	}
	return ""
}

type StepCondition struct {
	State                string   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StepCondition) Reset()         { *m = StepCondition{} }
func (m *StepCondition) String() string { return proto.CompactTextString(m) }
func (*StepCondition) ProtoMessage()    {}
func (*StepCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d227e95d8b5d27, []int{5}
}

func (m *StepCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepCondition.Unmarshal(m, b)
}
func (m *StepCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepCondition.Marshal(b, m, deterministic)
}
func (m *StepCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepCondition.Merge(m, src)
}
func (m *StepCondition) XXX_Size() int {
	return xxx_messageInfo_StepCondition.Size(m)
}
func (m *StepCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_StepCondition.DiscardUnknown(m)
}

var xxx_messageInfo_StepCondition proto.InternalMessageInfo

func (m *StepCondition) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *StepCondition) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *StepCondition) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UpdateReq struct {
	ObjectMeta           *ObjectMeta `protobuf:"bytes,1,opt,name=objectMeta,proto3" json:"objectMeta,omitempty"`
	State                *Status     `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d227e95d8b5d27, []int{6}
}

func (m *UpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReq.Unmarshal(m, b)
}
func (m *UpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReq.Marshal(b, m, deterministic)
}
func (m *UpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReq.Merge(m, src)
}
func (m *UpdateReq) XXX_Size() int {
	return xxx_messageInfo_UpdateReq.Size(m)
}
func (m *UpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReq proto.InternalMessageInfo

func (m *UpdateReq) GetObjectMeta() *ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *UpdateReq) GetState() *Status {
	if m != nil {
		return m.State
	}
	return nil
}

type DeleteReq struct {
	ObjectMeta           *ObjectMeta `protobuf:"bytes,1,opt,name=objectMeta,proto3" json:"objectMeta,omitempty"`
	State                *Status     `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeleteReq) Reset()         { *m = DeleteReq{} }
func (m *DeleteReq) String() string { return proto.CompactTextString(m) }
func (*DeleteReq) ProtoMessage()    {}
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d227e95d8b5d27, []int{7}
}

func (m *DeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReq.Unmarshal(m, b)
}
func (m *DeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReq.Marshal(b, m, deterministic)
}
func (m *DeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReq.Merge(m, src)
}
func (m *DeleteReq) XXX_Size() int {
	return xxx_messageInfo_DeleteReq.Size(m)
}
func (m *DeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReq proto.InternalMessageInfo

func (m *DeleteReq) GetObjectMeta() *ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *DeleteReq) GetState() *Status {
	if m != nil {
		return m.State
	}
	return nil
}

type Resp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_94d227e95d8b5d27, []int{8}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateReq)(nil), "v1alpha1.CreateReq")
	proto.RegisterType((*ObjectMeta)(nil), "v1alpha1.ObjectMeta")
	proto.RegisterMapType((map[string]string)(nil), "v1alpha1.ObjectMeta.LabelsEntry")
	proto.RegisterType((*OwnerReference)(nil), "v1alpha1.OwnerReference")
	proto.RegisterType((*Status)(nil), "v1alpha1.Status")
	proto.RegisterType((*Condition)(nil), "v1alpha1.Condition")
	proto.RegisterMapType((map[string]*StepCondition)(nil), "v1alpha1.Condition.ResourceRefEntry")
	proto.RegisterType((*StepCondition)(nil), "v1alpha1.StepCondition")
	proto.RegisterType((*UpdateReq)(nil), "v1alpha1.UpdateReq")
	proto.RegisterType((*DeleteReq)(nil), "v1alpha1.DeleteReq")
	proto.RegisterType((*Resp)(nil), "v1alpha1.Resp")
}

func init() { proto.RegisterFile("pkg/proto/v1alpha1/api.proto", fileDescriptor_94d227e95d8b5d27) }

var fileDescriptor_94d227e95d8b5d27 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xdf, 0x6b, 0xd4, 0x40,
	0x10, 0xe6, 0xd2, 0xbb, 0xd4, 0x4c, 0xf4, 0x38, 0xc6, 0xa2, 0xe1, 0xe8, 0xc3, 0x11, 0x44, 0xfa,
	0xd2, 0x1c, 0xbd, 0x8a, 0x54, 0x41, 0x10, 0xaa, 0x82, 0xa0, 0x08, 0x6b, 0xa5, 0xcf, 0x7b, 0xb9,
	0x69, 0x8d, 0xcd, 0x25, 0x31, 0xbb, 0x57, 0xe9, 0x3f, 0xe2, 0x9f, 0xe0, 0x8b, 0xff, 0xa4, 0xec,
	0x8f, 0x24, 0x7b, 0x47, 0x5e, 0x7c, 0xe9, 0x53, 0x76, 0x66, 0xbe, 0xfd, 0x66, 0xe7, 0xfb, 0xb2,
	0x0b, 0x87, 0xd5, 0xcd, 0xf5, 0xbc, 0xaa, 0x4b, 0x59, 0xce, 0x6f, 0x4f, 0x78, 0x5e, 0x7d, 0xe7,
	0x27, 0x73, 0x5e, 0x65, 0x89, 0x4e, 0xe1, 0x83, 0x26, 0x17, 0x67, 0x10, 0x9c, 0xd7, 0xc4, 0x25,
	0x31, 0xfa, 0x89, 0x2f, 0x00, 0xca, 0xe5, 0x0f, 0x4a, 0xe5, 0x67, 0x92, 0x3c, 0x1a, 0xcc, 0x06,
	0x47, 0xe1, 0xe2, 0x20, 0x69, 0xb0, 0xc9, 0x97, 0xb6, 0xc6, 0x1c, 0x1c, 0x3e, 0x87, 0x91, 0x90,
	0x5c, 0x52, 0xe4, 0xe9, 0x0d, 0x93, 0x6e, 0xc3, 0x57, 0xc9, 0xe5, 0x46, 0x30, 0x53, 0x8e, 0xff,
	0x7a, 0x00, 0x1d, 0x05, 0x22, 0x0c, 0x0b, 0xbe, 0x26, 0xdd, 0x26, 0x60, 0x7a, 0x8d, 0x87, 0x10,
	0xa8, 0x6f, 0x25, 0x78, 0x6a, 0xe8, 0x02, 0xd6, 0x25, 0x70, 0x02, 0x7b, 0x9b, 0x6c, 0x15, 0xed,
	0xe9, 0xbc, 0x5a, 0x62, 0x0c, 0x0f, 0x53, 0x75, 0xfa, 0xac, 0x2c, 0x2e, 0xb2, 0x35, 0x45, 0x43,
	0x5d, 0xda, 0xca, 0xe1, 0x19, 0xf8, 0x39, 0x5f, 0x52, 0x2e, 0xa2, 0xd1, 0x6c, 0xef, 0x28, 0x5c,
	0xcc, 0xfa, 0x06, 0x4a, 0x3e, 0x69, 0xc8, 0xfb, 0x42, 0xd6, 0x77, 0xcc, 0xe2, 0xf1, 0x2d, 0x8c,
	0xcb, 0x5f, 0x05, 0xd5, 0x8c, 0xae, 0xa8, 0xa6, 0x22, 0xa5, 0xc8, 0xd7, 0x0c, 0x91, 0xc3, 0xb0,
	0x55, 0x67, 0x3b, 0xf8, 0xe9, 0x2b, 0x08, 0x1d, 0x62, 0x35, 0xc0, 0x0d, 0xdd, 0xd9, 0x89, 0xd5,
	0x12, 0x0f, 0x60, 0x74, 0xcb, 0xf3, 0x4d, 0x33, 0xac, 0x09, 0x5e, 0x7b, 0x67, 0x83, 0xf8, 0x25,
	0x8c, 0xb7, 0xc9, 0x7b, 0x05, 0xb3, 0x92, 0x78, 0xad, 0x24, 0xf1, 0x1b, 0xf0, 0x8d, 0xec, 0x78,
	0x0a, 0x90, 0x96, 0xc5, 0x2a, 0x53, 0x4a, 0x08, 0xeb, 0xe6, 0xe3, 0xee, 0xe8, 0xe7, 0x4d, 0x8d,
	0x39, 0xb0, 0xf8, 0x8f, 0x07, 0x41, 0x5b, 0xc1, 0x27, 0xe0, 0x0b, 0x4d, 0x66, 0x9b, 0xda, 0x48,
	0xe5, 0x6b, 0xe2, 0xa2, 0x2c, 0x6c, 0x67, 0x1b, 0x61, 0x04, 0xfb, 0x6b, 0x12, 0x82, 0x5f, 0x93,
	0x75, 0xa9, 0x09, 0xf1, 0x03, 0x84, 0x35, 0x89, 0x72, 0x53, 0xa7, 0xc4, 0xe8, 0x2a, 0x1a, 0x6a,
	0x21, 0x9f, 0xf5, 0x9c, 0x26, 0x61, 0x1d, 0xcc, 0xd8, 0xe1, 0x6e, 0xc4, 0x04, 0x30, 0xe7, 0x42,
	0x5e, 0xd4, 0xbc, 0x10, 0x59, 0xeb, 0xfb, 0x48, 0x37, 0xeb, 0xa9, 0x4c, 0x2f, 0x61, 0xb2, 0x4b,
	0xd8, 0x63, 0xc3, 0xb1, 0x6b, 0x43, 0xb8, 0x78, 0xea, 0xfe, 0xc2, 0x54, 0x75, 0x4a, 0x39, 0xfe,
	0x5c, 0xc2, 0xa3, 0xad, 0x9a, 0xb2, 0xd2, 0x5c, 0x03, 0xc3, 0x6b, 0x82, 0xff, 0x57, 0x4a, 0xdd,
	0xc8, 0x6f, 0xd5, 0xea, 0x5e, 0x6e, 0x64, 0x06, 0xc1, 0x3b, 0xca, 0xe9, 0x3e, 0x5a, 0xf9, 0x30,
	0x64, 0x24, 0xaa, 0xc5, 0xef, 0x01, 0xec, 0x7f, 0x5c, 0x57, 0x39, 0x2f, 0x24, 0x1e, 0x83, 0x6f,
	0xde, 0x1e, 0x74, 0x7f, 0xcb, 0xe6, 0x35, 0x9a, 0x8e, 0xbb, 0xa4, 0xda, 0xaa, 0xe0, 0x46, 0x18,
	0x17, 0xde, 0x4a, 0xd5, 0x07, 0x37, 0xc3, 0xb9, 0xf0, 0x76, 0xdc, 0x5d, 0xf8, 0xd2, 0xd7, 0x2f,
	0xe3, 0xe9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0xd7, 0x2b, 0xa6, 0x39, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImplantClient is the client API for Implant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImplantClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*Resp, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*Resp, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Resp, error)
}

type implantClient struct {
	cc *grpc.ClientConn
}

func NewImplantClient(cc *grpc.ClientConn) ImplantClient {
	return &implantClient{cc}
}

func (c *implantClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/v1alpha1.Implant/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *implantClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/v1alpha1.Implant/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *implantClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/v1alpha1.Implant/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImplantServer is the server API for Implant service.
type ImplantServer interface {
	Create(context.Context, *CreateReq) (*Resp, error)
	Update(context.Context, *UpdateReq) (*Resp, error)
	Delete(context.Context, *DeleteReq) (*Resp, error)
}

// UnimplementedImplantServer can be embedded to have forward compatible implementations.
type UnimplementedImplantServer struct {
}

func (*UnimplementedImplantServer) Create(ctx context.Context, req *CreateReq) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedImplantServer) Update(ctx context.Context, req *UpdateReq) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedImplantServer) Delete(ctx context.Context, req *DeleteReq) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterImplantServer(s *grpc.Server, srv ImplantServer) {
	s.RegisterService(&_Implant_serviceDesc, srv)
}

func _Implant_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImplantServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Implant/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImplantServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Implant_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImplantServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Implant/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImplantServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Implant_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImplantServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Implant/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImplantServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Implant_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.Implant",
	HandlerType: (*ImplantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Implant_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Implant_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Implant_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/v1alpha1/api.proto",
}
