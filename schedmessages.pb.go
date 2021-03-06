// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schedmessages.proto

package schedulers

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Quotum struct {
	ProjectRatio float32 `protobuf:"fixed32,1,opt,name=ProjectRatio" json:"ProjectRatio,omitempty"`
	// Types that are valid to be assigned to Q:
	//	*Quotum_CpuHoursRatio
	//	*Quotum_CpuHoursAbs
	//	*Quotum_RamHoursAbs
	//	*Quotum_RamHoursRatio
	Q                    isQuotum_Q `protobuf_oneof:"Q"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Quotum) Reset()         { *m = Quotum{} }
func (m *Quotum) String() string { return proto.CompactTextString(m) }
func (*Quotum) ProtoMessage()    {}
func (*Quotum) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedmessages_885d9f7d21732d74, []int{0}
}
func (m *Quotum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quotum.Unmarshal(m, b)
}
func (m *Quotum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quotum.Marshal(b, m, deterministic)
}
func (dst *Quotum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quotum.Merge(dst, src)
}
func (m *Quotum) XXX_Size() int {
	return xxx_messageInfo_Quotum.Size(m)
}
func (m *Quotum) XXX_DiscardUnknown() {
	xxx_messageInfo_Quotum.DiscardUnknown(m)
}

var xxx_messageInfo_Quotum proto.InternalMessageInfo

type isQuotum_Q interface {
	isQuotum_Q()
}

type Quotum_CpuHoursRatio struct {
	CpuHoursRatio float32 `protobuf:"fixed32,10,opt,name=CpuHoursRatio,oneof"`
}
type Quotum_CpuHoursAbs struct {
	CpuHoursAbs float32 `protobuf:"fixed32,11,opt,name=CpuHoursAbs,oneof"`
}
type Quotum_RamHoursAbs struct {
	RamHoursAbs float32 `protobuf:"fixed32,12,opt,name=RamHoursAbs,oneof"`
}
type Quotum_RamHoursRatio struct {
	RamHoursRatio float32 `protobuf:"fixed32,13,opt,name=RamHoursRatio,oneof"`
}

func (*Quotum_CpuHoursRatio) isQuotum_Q() {}
func (*Quotum_CpuHoursAbs) isQuotum_Q()   {}
func (*Quotum_RamHoursAbs) isQuotum_Q()   {}
func (*Quotum_RamHoursRatio) isQuotum_Q() {}

func (m *Quotum) GetQ() isQuotum_Q {
	if m != nil {
		return m.Q
	}
	return nil
}

func (m *Quotum) GetProjectRatio() float32 {
	if m != nil {
		return m.ProjectRatio
	}
	return 0
}

func (m *Quotum) GetCpuHoursRatio() float32 {
	if x, ok := m.GetQ().(*Quotum_CpuHoursRatio); ok {
		return x.CpuHoursRatio
	}
	return 0
}

func (m *Quotum) GetCpuHoursAbs() float32 {
	if x, ok := m.GetQ().(*Quotum_CpuHoursAbs); ok {
		return x.CpuHoursAbs
	}
	return 0
}

func (m *Quotum) GetRamHoursAbs() float32 {
	if x, ok := m.GetQ().(*Quotum_RamHoursAbs); ok {
		return x.RamHoursAbs
	}
	return 0
}

func (m *Quotum) GetRamHoursRatio() float32 {
	if x, ok := m.GetQ().(*Quotum_RamHoursRatio); ok {
		return x.RamHoursRatio
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Quotum) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Quotum_OneofMarshaler, _Quotum_OneofUnmarshaler, _Quotum_OneofSizer, []interface{}{
		(*Quotum_CpuHoursRatio)(nil),
		(*Quotum_CpuHoursAbs)(nil),
		(*Quotum_RamHoursAbs)(nil),
		(*Quotum_RamHoursRatio)(nil),
	}
}

func _Quotum_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Quotum)
	// Q
	switch x := m.Q.(type) {
	case *Quotum_CpuHoursRatio:
		b.EncodeVarint(10<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.CpuHoursRatio)))
	case *Quotum_CpuHoursAbs:
		b.EncodeVarint(11<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.CpuHoursAbs)))
	case *Quotum_RamHoursAbs:
		b.EncodeVarint(12<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.RamHoursAbs)))
	case *Quotum_RamHoursRatio:
		b.EncodeVarint(13<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.RamHoursRatio)))
	case nil:
	default:
		return fmt.Errorf("Quotum.Q has unexpected type %T", x)
	}
	return nil
}

func _Quotum_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Quotum)
	switch tag {
	case 10: // Q.CpuHoursRatio
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Q = &Quotum_CpuHoursRatio{math.Float32frombits(uint32(x))}
		return true, err
	case 11: // Q.CpuHoursAbs
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Q = &Quotum_CpuHoursAbs{math.Float32frombits(uint32(x))}
		return true, err
	case 12: // Q.RamHoursAbs
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Q = &Quotum_RamHoursAbs{math.Float32frombits(uint32(x))}
		return true, err
	case 13: // Q.RamHoursRatio
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Q = &Quotum_RamHoursRatio{math.Float32frombits(uint32(x))}
		return true, err
	default:
		return false, nil
	}
}

func _Quotum_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Quotum)
	// Q
	switch x := m.Q.(type) {
	case *Quotum_CpuHoursRatio:
		n += 1 // tag and wire
		n += 4
	case *Quotum_CpuHoursAbs:
		n += 1 // tag and wire
		n += 4
	case *Quotum_RamHoursAbs:
		n += 1 // tag and wire
		n += 4
	case *Quotum_RamHoursRatio:
		n += 1 // tag and wire
		n += 4
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Organization struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Quota                *Quotum  `protobuf:"bytes,2,opt,name=Quota" json:"Quota,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Organization) Reset()         { *m = Organization{} }
func (m *Organization) String() string { return proto.CompactTextString(m) }
func (*Organization) ProtoMessage()    {}
func (*Organization) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedmessages_885d9f7d21732d74, []int{1}
}
func (m *Organization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Organization.Unmarshal(m, b)
}
func (m *Organization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Organization.Marshal(b, m, deterministic)
}
func (dst *Organization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Organization.Merge(dst, src)
}
func (m *Organization) XXX_Size() int {
	return xxx_messageInfo_Organization.Size(m)
}
func (m *Organization) XXX_DiscardUnknown() {
	xxx_messageInfo_Organization.DiscardUnknown(m)
}

var xxx_messageInfo_Organization proto.InternalMessageInfo

func (m *Organization) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Organization) GetQuota() *Quotum {
	if m != nil {
		return m.Quota
	}
	return nil
}

type ResourceVolume struct {
	Id                   uint64        `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	CPU                  uint32        `protobuf:"varint,2,opt,name=CPU" json:"CPU,omitempty"`
	GPU                  uint32        `protobuf:"varint,3,opt,name=GPU" json:"GPU,omitempty"`
	RAMmb                uint32        `protobuf:"varint,4,opt,name=RAMmb" json:"RAMmb,omitempty"`
	TimePeriod           uint64        `protobuf:"varint,5,opt,name=TimePeriod" json:"TimePeriod,omitempty"`
	Owner                *Organization `protobuf:"bytes,6,opt,name=Owner" json:"Owner,omitempty"`
	Param                uint64        `protobuf:"varint,7,opt,name=Param" json:"Param,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ResourceVolume) Reset()         { *m = ResourceVolume{} }
func (m *ResourceVolume) String() string { return proto.CompactTextString(m) }
func (*ResourceVolume) ProtoMessage()    {}
func (*ResourceVolume) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedmessages_885d9f7d21732d74, []int{2}
}
func (m *ResourceVolume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceVolume.Unmarshal(m, b)
}
func (m *ResourceVolume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceVolume.Marshal(b, m, deterministic)
}
func (dst *ResourceVolume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceVolume.Merge(dst, src)
}
func (m *ResourceVolume) XXX_Size() int {
	return xxx_messageInfo_ResourceVolume.Size(m)
}
func (m *ResourceVolume) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceVolume.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceVolume proto.InternalMessageInfo

func (m *ResourceVolume) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ResourceVolume) GetCPU() uint32 {
	if m != nil {
		return m.CPU
	}
	return 0
}

func (m *ResourceVolume) GetGPU() uint32 {
	if m != nil {
		return m.GPU
	}
	return 0
}

func (m *ResourceVolume) GetRAMmb() uint32 {
	if m != nil {
		return m.RAMmb
	}
	return 0
}

func (m *ResourceVolume) GetTimePeriod() uint64 {
	if m != nil {
		return m.TimePeriod
	}
	return 0
}

func (m *ResourceVolume) GetOwner() *Organization {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ResourceVolume) GetParam() uint64 {
	if m != nil {
		return m.Param
	}
	return 0
}

type Decision struct {
	// Pairs jobs[idx] -> workers[idx]
	JobIdx               uint64   `protobuf:"varint,1,opt,name=JobIdx" json:"JobIdx,omitempty"`
	WorkerIdx            uint64   `protobuf:"varint,2,opt,name=WorkerIdx" json:"WorkerIdx,omitempty"`
	CoresNum             uint32   `protobuf:"varint,3,opt,name=CoresNum" json:"CoresNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Decision) Reset()         { *m = Decision{} }
func (m *Decision) String() string { return proto.CompactTextString(m) }
func (*Decision) ProtoMessage()    {}
func (*Decision) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedmessages_885d9f7d21732d74, []int{3}
}
func (m *Decision) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Decision.Unmarshal(m, b)
}
func (m *Decision) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Decision.Marshal(b, m, deterministic)
}
func (dst *Decision) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decision.Merge(dst, src)
}
func (m *Decision) XXX_Size() int {
	return xxx_messageInfo_Decision.Size(m)
}
func (m *Decision) XXX_DiscardUnknown() {
	xxx_messageInfo_Decision.DiscardUnknown(m)
}

var xxx_messageInfo_Decision proto.InternalMessageInfo

func (m *Decision) GetJobIdx() uint64 {
	if m != nil {
		return m.JobIdx
	}
	return 0
}

func (m *Decision) GetWorkerIdx() uint64 {
	if m != nil {
		return m.WorkerIdx
	}
	return 0
}

func (m *Decision) GetCoresNum() uint32 {
	if m != nil {
		return m.CoresNum
	}
	return 0
}

func init() {
	proto.RegisterType((*Quotum)(nil), "schedulers.Quotum")
	proto.RegisterType((*Organization)(nil), "schedulers.Organization")
	proto.RegisterType((*ResourceVolume)(nil), "schedulers.ResourceVolume")
	proto.RegisterType((*Decision)(nil), "schedulers.Decision")
}

func init() { proto.RegisterFile("schedmessages.proto", fileDescriptor_schedmessages_885d9f7d21732d74) }

var fileDescriptor_schedmessages_885d9f7d21732d74 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xcf, 0x6e, 0xe2, 0x30,
	0x10, 0xc6, 0x37, 0x81, 0x64, 0x61, 0xf8, 0xa3, 0x95, 0x77, 0xb5, 0xb2, 0x56, 0xab, 0x15, 0xca,
	0x61, 0xc5, 0x29, 0x87, 0xdd, 0x27, 0xa0, 0x54, 0x02, 0xaa, 0x16, 0x82, 0x55, 0xda, 0x4b, 0x2f,
	0x4e, 0x62, 0xd1, 0xb4, 0x18, 0x23, 0x3b, 0x56, 0xab, 0x3e, 0x5d, 0xaf, 0x7d, 0xab, 0xca, 0x76,
	0x02, 0xe1, 0xe6, 0xef, 0x37, 0x9f, 0x66, 0xe6, 0x1b, 0x19, 0xbe, 0xab, 0xec, 0x91, 0xe5, 0x9c,
	0x29, 0x45, 0xb7, 0x4c, 0xc5, 0x07, 0x29, 0x4a, 0x81, 0xc0, 0x42, 0xbd, 0x63, 0x52, 0x45, 0x1f,
	0x1e, 0x84, 0x6b, 0x2d, 0x4a, 0xcd, 0x51, 0x04, 0xfd, 0x44, 0x8a, 0x27, 0x96, 0x95, 0x84, 0x96,
	0x85, 0xc0, 0xde, 0xc8, 0x1b, 0xfb, 0xe4, 0x8c, 0xa1, 0xbf, 0x30, 0x98, 0x1e, 0xf4, 0x5c, 0x68,
	0xa9, 0x9c, 0x09, 0x8c, 0x69, 0xfe, 0x85, 0x9c, 0x63, 0x14, 0x41, 0xaf, 0x06, 0x93, 0x54, 0xe1,
	0x5e, 0xe5, 0x6a, 0x42, 0xe3, 0x21, 0x94, 0x1f, 0x3d, 0xfd, 0xda, 0xd3, 0x80, 0x66, 0x5e, 0x2d,
	0xdd, 0xbc, 0x41, 0x3d, 0xef, 0x0c, 0x5f, 0xb4, 0xc0, 0x5b, 0x47, 0xd7, 0xd0, 0x5f, 0xc9, 0x2d,
	0xdd, 0x17, 0x6f, 0x86, 0xed, 0x11, 0x82, 0xf6, 0x92, 0x72, 0x66, 0x83, 0x74, 0x89, 0x7d, 0xa3,
	0x31, 0x04, 0x26, 0x2e, 0xc5, 0xfe, 0xc8, 0x1b, 0xf7, 0xfe, 0xa1, 0xf8, 0x74, 0x8b, 0xd8, 0xdd,
	0x81, 0x38, 0x43, 0xf4, 0xee, 0xc1, 0x90, 0x30, 0x25, 0xb4, 0xcc, 0xd8, 0x9d, 0xd8, 0x69, 0xce,
	0xd0, 0x10, 0xfc, 0x45, 0x6e, 0xdb, 0xb5, 0x89, 0xbf, 0xc8, 0xd1, 0x37, 0x68, 0x4d, 0x93, 0x8d,
	0x6d, 0x35, 0x20, 0xe6, 0x69, 0xc8, 0x2c, 0xd9, 0xe0, 0x96, 0x23, 0xb3, 0x64, 0x83, 0x7e, 0x40,
	0x40, 0x26, 0x37, 0x3c, 0xc5, 0x6d, 0xcb, 0x9c, 0x40, 0x7f, 0x00, 0x6e, 0x0b, 0xce, 0x12, 0x26,
	0x0b, 0x91, 0xe3, 0xc0, 0x76, 0x6c, 0x10, 0x14, 0x43, 0xb0, 0x7a, 0xd9, 0x33, 0x89, 0x43, 0xbb,
	0x26, 0x6e, 0xae, 0xd9, 0xcc, 0x48, 0x9c, 0xcd, 0x4c, 0x49, 0xa8, 0xa4, 0x1c, 0x7f, 0xb5, 0xad,
	0x9c, 0x88, 0x1e, 0xa0, 0x73, 0xc9, 0xb2, 0x42, 0x99, 0x63, 0xfc, 0x84, 0xf0, 0x4a, 0xa4, 0x8b,
	0xfc, 0xb5, 0xda, 0xbf, 0x52, 0xe8, 0x37, 0x74, 0xef, 0x85, 0x7c, 0x66, 0xd2, 0x94, 0x7c, 0x5b,
	0x3a, 0x01, 0xf4, 0x0b, 0x3a, 0x53, 0x21, 0x99, 0x5a, 0x6a, 0x5e, 0x85, 0x3a, 0xea, 0x34, 0xb4,
	0xbf, 0xe9, 0xff, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0x37, 0x21, 0x2c, 0x64, 0x02, 0x00,
	0x00,
}
