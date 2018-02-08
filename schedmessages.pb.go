// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schedmessages.proto

/*
Package scheduler is a generated protocol buffer package.

It is generated from these files:
	schedmessages.proto

It has these top-level messages:
	Quotum
	Organization
	ResourceVolume
	Decision
*/
package scheduler

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
	Q isQuotum_Q `protobuf_oneof:"Q"`
}

func (m *Quotum) Reset()                    { *m = Quotum{} }
func (m *Quotum) String() string            { return proto.CompactTextString(m) }
func (*Quotum) ProtoMessage()               {}
func (*Quotum) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
		n += proto.SizeVarint(10<<3 | proto.WireFixed32)
		n += 4
	case *Quotum_CpuHoursAbs:
		n += proto.SizeVarint(11<<3 | proto.WireFixed32)
		n += 4
	case *Quotum_RamHoursAbs:
		n += proto.SizeVarint(12<<3 | proto.WireFixed32)
		n += 4
	case *Quotum_RamHoursRatio:
		n += proto.SizeVarint(13<<3 | proto.WireFixed32)
		n += 4
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Organization struct {
	Name  string  `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Quota *Quotum `protobuf:"bytes,2,opt,name=Quota" json:"Quota,omitempty"`
}

func (m *Organization) Reset()                    { *m = Organization{} }
func (m *Organization) String() string            { return proto.CompactTextString(m) }
func (*Organization) ProtoMessage()               {}
func (*Organization) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
	Id         uint64        `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	CPU        uint32        `protobuf:"varint,2,opt,name=CPU" json:"CPU,omitempty"`
	GPU        uint32        `protobuf:"varint,3,opt,name=GPU" json:"GPU,omitempty"`
	RAMmb      uint32        `protobuf:"varint,4,opt,name=RAMmb" json:"RAMmb,omitempty"`
	TimePeriod uint64        `protobuf:"varint,5,opt,name=TimePeriod" json:"TimePeriod,omitempty"`
	Owner      *Organization `protobuf:"bytes,6,opt,name=Owner" json:"Owner,omitempty"`
}

func (m *ResourceVolume) Reset()                    { *m = ResourceVolume{} }
func (m *ResourceVolume) String() string            { return proto.CompactTextString(m) }
func (*ResourceVolume) ProtoMessage()               {}
func (*ResourceVolume) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

type Decision struct {
	// Pairs jobs[idx] -> workers[idx]
	JobIdx    uint64 `protobuf:"varint,1,opt,name=JobIdx" json:"JobIdx,omitempty"`
	WorkerIdx uint64 `protobuf:"varint,2,opt,name=WorkerIdx" json:"WorkerIdx,omitempty"`
}

func (m *Decision) Reset()                    { *m = Decision{} }
func (m *Decision) String() string            { return proto.CompactTextString(m) }
func (*Decision) ProtoMessage()               {}
func (*Decision) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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

func init() {
	proto.RegisterType((*Quotum)(nil), "scheduler.Quotum")
	proto.RegisterType((*Organization)(nil), "scheduler.Organization")
	proto.RegisterType((*ResourceVolume)(nil), "scheduler.ResourceVolume")
	proto.RegisterType((*Decision)(nil), "scheduler.Decision")
}

func init() { proto.RegisterFile("schedmessages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xcb, 0x4e, 0xf2, 0x40,
	0x14, 0x80, 0xff, 0x16, 0xda, 0xfc, 0x1c, 0x2e, 0xd1, 0xd1, 0x68, 0x17, 0xc6, 0x90, 0x2e, 0x94,
	0x8d, 0x2c, 0xf4, 0x05, 0x44, 0x4c, 0x00, 0x8d, 0x52, 0x26, 0xa2, 0xeb, 0x5e, 0x4e, 0xb0, 0xca,
	0x30, 0x64, 0xa6, 0x13, 0x8d, 0x6f, 0xe4, 0xa3, 0xf8, 0x56, 0x66, 0x66, 0x28, 0x96, 0xdd, 0x9c,
	0xef, 0x7c, 0x39, 0xb7, 0x0c, 0x1c, 0xc8, 0xf4, 0x15, 0x33, 0x86, 0x52, 0xc6, 0x0b, 0x94, 0xfd,
	0xb5, 0xe0, 0x05, 0x27, 0x0d, 0x03, 0xd5, 0x12, 0x45, 0xf8, 0xe3, 0x80, 0x3f, 0x53, 0xbc, 0x50,
	0x8c, 0x84, 0xd0, 0x8a, 0x04, 0x7f, 0xc3, 0xb4, 0xa0, 0x71, 0x91, 0xf3, 0xc0, 0xe9, 0x3a, 0x3d,
	0x97, 0xee, 0x30, 0x72, 0x06, 0xed, 0xe1, 0x5a, 0x8d, 0xb9, 0x12, 0xd2, 0x4a, 0xa0, 0xa5, 0xf1,
	0x3f, 0xba, 0x8b, 0x49, 0x08, 0xcd, 0x12, 0x0c, 0x12, 0x19, 0x34, 0x37, 0x56, 0x15, 0x6a, 0x87,
	0xc6, 0x6c, 0xeb, 0xb4, 0x4a, 0xa7, 0x02, 0x75, 0xbf, 0x32, 0xb4, 0xfd, 0xda, 0x65, 0xbf, 0x1d,
	0x7c, 0x53, 0x03, 0x67, 0x16, 0xde, 0x43, 0x6b, 0x2a, 0x16, 0xf1, 0x2a, 0xff, 0xd2, 0x6c, 0x45,
	0x08, 0xd4, 0x1f, 0x63, 0x86, 0x66, 0x91, 0x06, 0x35, 0x6f, 0x72, 0x0e, 0x9e, 0x5e, 0x37, 0x0e,
	0xdc, 0xae, 0xd3, 0x6b, 0x5e, 0xee, 0xf7, 0xb7, 0xa7, 0xe8, 0xdb, 0x33, 0x50, 0x9b, 0x0f, 0xbf,
	0x1d, 0xe8, 0x50, 0x94, 0x5c, 0x89, 0x14, 0x9f, 0xf9, 0x52, 0x31, 0x24, 0x1d, 0x70, 0x27, 0x99,
	0xa9, 0x56, 0xa7, 0xee, 0x24, 0x23, 0x7b, 0x50, 0x1b, 0x46, 0x73, 0x53, 0xa9, 0x4d, 0xf5, 0x53,
	0x93, 0x51, 0x34, 0x0f, 0x6a, 0x96, 0x8c, 0xa2, 0x39, 0x39, 0x04, 0x8f, 0x0e, 0x1e, 0x58, 0x12,
	0xd4, 0x0d, 0xb3, 0x01, 0x39, 0x05, 0x78, 0xca, 0x19, 0x46, 0x28, 0x72, 0x9e, 0x05, 0x9e, 0xa9,
	0x58, 0x21, 0xe4, 0x02, 0xbc, 0xe9, 0xc7, 0x0a, 0x45, 0xe0, 0x9b, 0x29, 0x8f, 0x2b, 0x53, 0x56,
	0x37, 0xa4, 0xd6, 0x0a, 0xaf, 0xe1, 0xff, 0x2d, 0xa6, 0xb9, 0xd4, 0x4b, 0x1f, 0x81, 0x7f, 0xc7,
	0x93, 0x49, 0xf6, 0xb9, 0x19, 0x74, 0x13, 0x91, 0x13, 0x68, 0xbc, 0x70, 0xf1, 0x8e, 0x42, 0xa7,
	0x5c, 0x93, 0xfa, 0x03, 0x89, 0x6f, 0x3e, 0xc6, 0xd5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfb,
	0xbb, 0xfa, 0x5a, 0x2f, 0x02, 0x00, 0x00,
}