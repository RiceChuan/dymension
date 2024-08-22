// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/dymns/dym_name.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// DymNameConfigType specifies the type of the Dym-Name configuration.
// Currently only supports Name, similar to DNS.
type DymNameConfigType int32

const (
	DymNameConfigType_DCT_UNKNOWN DymNameConfigType = 0
	DymNameConfigType_DCT_NAME    DymNameConfigType = 1
)

var DymNameConfigType_name = map[int32]string{
	0: "DCT_UNKNOWN",
	1: "DCT_NAME",
}

var DymNameConfigType_value = map[string]int32{
	"DCT_UNKNOWN": 0,
	"DCT_NAME":    1,
}

func (x DymNameConfigType) String() string {
	return proto.EnumName(DymNameConfigType_name, int32(x))
}

func (DymNameConfigType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_463436600bef60e6, []int{0}
}

// DymName defines a Dym-Name, the mainly purpose is to store ownership and resolution information.
// Dym-Name is similar to DNS. It is a human-readable name that maps to a chain address.
// One Dym-Name can have multiple configurations, each configuration is a resolution record.
// Dym-Name is owned by an account, and is able to grant permission to another account to control the Dym-Name.
type DymName struct {
	// name is the human-readable name of the Dym-Name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// owner is the account address that owns the Dym-Name. Owner has permission to transfer ownership.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// controller is the account address that has permission update configuration for the Dym-Name.
	// Default is the owner. Able to transfer control to another account by the owner.
	// Users can set Dym-Name owned by Cold-Wallet and controlled by Hot-Wallet.
	Controller string `protobuf:"bytes,3,opt,name=controller,proto3" json:"controller,omitempty"`
	// expire_at is the UTC epoch represent the last effective date of the Dym-Name,
	// after which the Dym-Name is no longer valid.
	// NOTE: Expired Dym-Names are not deleted from the store
	// because iterating through store is very expensive because expiry date must be checked every use.
	ExpireAt int64 `protobuf:"varint,4,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	// configs are configuration records for the Dym-Name.
	Configs []DymNameConfig `protobuf:"bytes,5,rep,name=configs,proto3" json:"configs"`
	// contact is an optional information for the Dym-Name.
	// Convenient for retails users.
	Contact string `protobuf:"bytes,6,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (m *DymName) Reset()         { *m = DymName{} }
func (m *DymName) String() string { return proto.CompactTextString(m) }
func (*DymName) ProtoMessage()    {}
func (*DymName) Descriptor() ([]byte, []int) {
	return fileDescriptor_463436600bef60e6, []int{0}
}
func (m *DymName) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DymName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DymName.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DymName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DymName.Merge(m, src)
}
func (m *DymName) XXX_Size() int {
	return m.Size()
}
func (m *DymName) XXX_DiscardUnknown() {
	xxx_messageInfo_DymName.DiscardUnknown(m)
}

var xxx_messageInfo_DymName proto.InternalMessageInfo

func (m *DymName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DymName) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DymName) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *DymName) GetExpireAt() int64 {
	if m != nil {
		return m.ExpireAt
	}
	return 0
}

func (m *DymName) GetConfigs() []DymNameConfig {
	if m != nil {
		return m.Configs
	}
	return nil
}

func (m *DymName) GetContact() string {
	if m != nil {
		return m.Contact
	}
	return ""
}

// DymNameConfig contains the resolution configuration for the Dym-Name.
// Each record is a resolution record, similar to DNS.
type DymNameConfig struct {
	// type is the type of the Dym-Name configuration (equals to Type in DNS).
	Type DymNameConfigType `protobuf:"varint,1,opt,name=type,proto3,enum=dymensionxyz.dymension.dymns.DymNameConfigType" json:"type,omitempty"`
	// chain_id is the chain-id of the Dym-Name configuration (equals to top-level-domain).
	// If empty, the configuration is for host chain (Dymension Hub).
	ChainId string `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// path of the Dym-Name configuration (equals to Host in DNS).
	// If the type of this config record is Name, it is the Sub-Name of the Dym-Name Address.
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// value of the Dym-Name configuration resolves to (equals to Value in DNS).
	// If the type of this config record is Name, it is the address which the Dym-Name Address resolves to.
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *DymNameConfig) Reset()         { *m = DymNameConfig{} }
func (m *DymNameConfig) String() string { return proto.CompactTextString(m) }
func (*DymNameConfig) ProtoMessage()    {}
func (*DymNameConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_463436600bef60e6, []int{1}
}
func (m *DymNameConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DymNameConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DymNameConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DymNameConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DymNameConfig.Merge(m, src)
}
func (m *DymNameConfig) XXX_Size() int {
	return m.Size()
}
func (m *DymNameConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DymNameConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DymNameConfig proto.InternalMessageInfo

func (m *DymNameConfig) GetType() DymNameConfigType {
	if m != nil {
		return m.Type
	}
	return DymNameConfigType_DCT_UNKNOWN
}

func (m *DymNameConfig) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *DymNameConfig) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *DymNameConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// ReverseLookupDymNames contains a list of Dym-Names for reverse lookup.
type ReverseLookupDymNames struct {
	// dym_names is a list of name of the Dym-Names linked to the reverse-lookup record.
	DymNames []string `protobuf:"bytes,1,rep,name=dym_names,json=dymNames,proto3" json:"dym_names,omitempty"`
}

func (m *ReverseLookupDymNames) Reset()         { *m = ReverseLookupDymNames{} }
func (m *ReverseLookupDymNames) String() string { return proto.CompactTextString(m) }
func (*ReverseLookupDymNames) ProtoMessage()    {}
func (*ReverseLookupDymNames) Descriptor() ([]byte, []int) {
	return fileDescriptor_463436600bef60e6, []int{2}
}
func (m *ReverseLookupDymNames) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReverseLookupDymNames) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReverseLookupDymNames.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReverseLookupDymNames) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReverseLookupDymNames.Merge(m, src)
}
func (m *ReverseLookupDymNames) XXX_Size() int {
	return m.Size()
}
func (m *ReverseLookupDymNames) XXX_DiscardUnknown() {
	xxx_messageInfo_ReverseLookupDymNames.DiscardUnknown(m)
}

var xxx_messageInfo_ReverseLookupDymNames proto.InternalMessageInfo

func (m *ReverseLookupDymNames) GetDymNames() []string {
	if m != nil {
		return m.DymNames
	}
	return nil
}

func init() {
	proto.RegisterEnum("dymensionxyz.dymension.dymns.DymNameConfigType", DymNameConfigType_name, DymNameConfigType_value)
	proto.RegisterType((*DymName)(nil), "dymensionxyz.dymension.dymns.DymName")
	proto.RegisterType((*DymNameConfig)(nil), "dymensionxyz.dymension.dymns.DymNameConfig")
	proto.RegisterType((*ReverseLookupDymNames)(nil), "dymensionxyz.dymension.dymns.ReverseLookupDymNames")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/dymns/dym_name.proto", fileDescriptor_463436600bef60e6)
}

var fileDescriptor_463436600bef60e6 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x8e, 0x69, 0xb7, 0xb6, 0x1e, 0x3f, 0xc3, 0x1a, 0x92, 0x19, 0xc8, 0x54, 0xbd, 0x8a, 0x98,
	0x14, 0x6b, 0x19, 0x2f, 0xb0, 0x75, 0x5c, 0xa0, 0x8d, 0x20, 0x45, 0x43, 0x48, 0xdc, 0x44, 0x4e,
	0x62, 0xd2, 0x88, 0xc6, 0x8e, 0x62, 0x37, 0x34, 0x3c, 0x05, 0xb7, 0xbc, 0xd1, 0x2e, 0x77, 0x07,
	0x57, 0x08, 0xb5, 0x2f, 0x82, 0xec, 0xa4, 0x53, 0x11, 0x02, 0x89, 0x9b, 0xe8, 0x7c, 0xdf, 0x39,
	0x9f, 0xcf, 0x39, 0x5f, 0x0e, 0x3c, 0x4a, 0x9b, 0x82, 0x0b, 0x95, 0x4b, 0xb1, 0x6c, 0x3e, 0xd3,
	0x5b, 0x60, 0x22, 0xa1, 0xcc, 0x37, 0x12, 0xac, 0xe0, 0x5e, 0x59, 0x49, 0x2d, 0xd1, 0xd3, 0xed,
	0x62, 0xef, 0x16, 0x78, 0xb6, 0xf8, 0xf0, 0x20, 0x93, 0x99, 0xb4, 0x85, 0xd4, 0x44, 0xad, 0xe6,
	0x90, 0x24, 0x52, 0x15, 0x52, 0xd1, 0x98, 0x29, 0x4e, 0xeb, 0xe3, 0x98, 0x6b, 0x76, 0x4c, 0x13,
	0x99, 0x8b, 0x36, 0x3f, 0xf9, 0x06, 0xe0, 0xe0, 0xbc, 0x29, 0x02, 0x56, 0x70, 0x84, 0x60, 0xdf,
	0x74, 0xc3, 0x60, 0x0c, 0xdc, 0x51, 0x68, 0x63, 0x74, 0x00, 0x77, 0xe4, 0x27, 0xc1, 0x2b, 0x7c,
	0xc7, 0x92, 0x2d, 0x40, 0x04, 0xc2, 0x44, 0x0a, 0x5d, 0xc9, 0xf9, 0x9c, 0x57, 0xb8, 0x67, 0x53,
	0x5b, 0x0c, 0x7a, 0x02, 0x47, 0x7c, 0x59, 0xe6, 0x15, 0x8f, 0x98, 0xc6, 0xfd, 0x31, 0x70, 0x7b,
	0xe1, 0xb0, 0x25, 0x4e, 0x35, 0xba, 0x80, 0x83, 0x44, 0x8a, 0x0f, 0x79, 0xa6, 0xf0, 0xce, 0xb8,
	0xe7, 0xee, 0xf9, 0x47, 0xde, 0xbf, 0x16, 0xf3, 0xba, 0xf1, 0xa6, 0x56, 0x73, 0xd6, 0xbf, 0xfe,
	0xf1, 0xcc, 0x09, 0x37, 0x2f, 0x20, 0x6c, 0x1f, 0xd3, 0x2c, 0xd1, 0x78, 0xd7, 0x8e, 0xb1, 0x81,
	0x93, 0xaf, 0x00, 0xde, 0xfb, 0x4d, 0x8a, 0xa6, 0xb0, 0xaf, 0x9b, 0xb2, 0xdd, 0xef, 0xbe, 0x4f,
	0xff, 0xa3, 0xeb, 0x55, 0x53, 0xf2, 0xd0, 0x8a, 0xd1, 0x63, 0x38, 0x4c, 0x66, 0x2c, 0x17, 0x51,
	0x9e, 0x76, 0x9e, 0x0c, 0x2c, 0x7e, 0x95, 0x1a, 0xff, 0x4a, 0xa6, 0x67, 0x9d, 0x1f, 0x36, 0x36,
	0xfe, 0xd5, 0x6c, 0xbe, 0xe0, 0xd6, 0x85, 0x51, 0xd8, 0x82, 0xc9, 0x0b, 0xf8, 0x28, 0xe4, 0x35,
	0xaf, 0x14, 0xbf, 0x94, 0xf2, 0xe3, 0xa2, 0xec, 0x9a, 0x29, 0x63, 0xdc, 0xe6, 0xa7, 0x2b, 0x0c,
	0xc6, 0x3d, 0x77, 0x14, 0x0e, 0xd3, 0x2e, 0xf9, 0xdc, 0x87, 0x0f, 0xff, 0x98, 0x0a, 0x3d, 0x80,
	0x7b, 0xe7, 0xd3, 0xab, 0xe8, 0x6d, 0x70, 0x11, 0xbc, 0x79, 0x17, 0xec, 0x3b, 0xe8, 0x2e, 0x1c,
	0x1a, 0x22, 0x38, 0x7d, 0xfd, 0x72, 0x1f, 0x9c, 0x5d, 0x5e, 0xaf, 0x08, 0xb8, 0x59, 0x11, 0xf0,
	0x73, 0x45, 0xc0, 0x97, 0x35, 0x71, 0x6e, 0xd6, 0xc4, 0xf9, 0xbe, 0x26, 0xce, 0x7b, 0x3f, 0xcb,
	0xf5, 0x6c, 0x11, 0x7b, 0x89, 0x2c, 0xe8, 0x5f, 0xae, 0xb0, 0x3e, 0xa1, 0xcb, 0xee, 0x14, 0xcd,
	0xee, 0x2a, 0xde, 0xb5, 0x47, 0x73, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x28, 0x32, 0x1e, 0x91,
	0xb7, 0x02, 0x00, 0x00,
}

func (m *DymName) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DymName) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DymName) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contact) > 0 {
		i -= len(m.Contact)
		copy(dAtA[i:], m.Contact)
		i = encodeVarintDymName(dAtA, i, uint64(len(m.Contact)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Configs) > 0 {
		for iNdEx := len(m.Configs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Configs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDymName(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.ExpireAt != 0 {
		i = encodeVarintDymName(dAtA, i, uint64(m.ExpireAt))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintDymName(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintDymName(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDymName(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DymNameConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DymNameConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DymNameConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintDymName(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintDymName(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintDymName(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintDymName(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReverseLookupDymNames) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReverseLookupDymNames) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReverseLookupDymNames) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DymNames) > 0 {
		for iNdEx := len(m.DymNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DymNames[iNdEx])
			copy(dAtA[i:], m.DymNames[iNdEx])
			i = encodeVarintDymName(dAtA, i, uint64(len(m.DymNames[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDymName(dAtA []byte, offset int, v uint64) int {
	offset -= sovDymName(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DymName) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDymName(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDymName(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovDymName(uint64(l))
	}
	if m.ExpireAt != 0 {
		n += 1 + sovDymName(uint64(m.ExpireAt))
	}
	if len(m.Configs) > 0 {
		for _, e := range m.Configs {
			l = e.Size()
			n += 1 + l + sovDymName(uint64(l))
		}
	}
	l = len(m.Contact)
	if l > 0 {
		n += 1 + l + sovDymName(uint64(l))
	}
	return n
}

func (m *DymNameConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovDymName(uint64(m.Type))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovDymName(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovDymName(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovDymName(uint64(l))
	}
	return n
}

func (m *ReverseLookupDymNames) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DymNames) > 0 {
		for _, s := range m.DymNames {
			l = len(s)
			n += 1 + l + sovDymName(uint64(l))
		}
	}
	return n
}

func sovDymName(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDymName(x uint64) (n int) {
	return sovDymName(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DymName) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDymName
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DymName: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DymName: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDymName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDymName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDymName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDymName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDymName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDymName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireAt", wireType)
			}
			m.ExpireAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpireAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Configs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDymName
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDymName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Configs = append(m.Configs, DymNameConfig{})
			if err := m.Configs[len(m.Configs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDymName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDymName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contact = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDymName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDymName
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DymNameConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDymName
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DymNameConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DymNameConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= DymNameConfigType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDymName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDymName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDymName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDymName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDymName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDymName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDymName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDymName
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReverseLookupDymNames) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDymName
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReverseLookupDymNames: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReverseLookupDymNames: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DymNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDymName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDymName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DymNames = append(m.DymNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDymName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDymName
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDymName(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDymName
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDymName
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDymName
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDymName
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDymName
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDymName        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDymName          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDymName = fmt.Errorf("proto: unexpected end of group")
)