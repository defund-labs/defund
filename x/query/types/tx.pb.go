// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgCreateInterqueryResult struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Storeid string `protobuf:"bytes,2,opt,name=storeid,proto3" json:"storeid,omitempty"`
	// data is submitted as a base64 encoded string but is broken down to bytes to be stored
	Data   string           `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Height *types.Height    `protobuf:"bytes,4,opt,name=height,proto3" json:"height,omitempty"`
	Proof  *crypto.ProofOps `protobuf:"bytes,5,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *MsgCreateInterqueryResult) Reset()         { *m = MsgCreateInterqueryResult{} }
func (m *MsgCreateInterqueryResult) String() string { return proto.CompactTextString(m) }
func (*MsgCreateInterqueryResult) ProtoMessage()    {}
func (*MsgCreateInterqueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_1defee906b3ee117, []int{0}
}
func (m *MsgCreateInterqueryResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateInterqueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateInterqueryResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateInterqueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateInterqueryResult.Merge(m, src)
}
func (m *MsgCreateInterqueryResult) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateInterqueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateInterqueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateInterqueryResult proto.InternalMessageInfo

func (m *MsgCreateInterqueryResult) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateInterqueryResult) GetStoreid() string {
	if m != nil {
		return m.Storeid
	}
	return ""
}

func (m *MsgCreateInterqueryResult) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *MsgCreateInterqueryResult) GetHeight() *types.Height {
	if m != nil {
		return m.Height
	}
	return nil
}

func (m *MsgCreateInterqueryResult) GetProof() *crypto.ProofOps {
	if m != nil {
		return m.Proof
	}
	return nil
}

type MsgCreateInterqueryResultResponse struct {
}

func (m *MsgCreateInterqueryResultResponse) Reset()         { *m = MsgCreateInterqueryResultResponse{} }
func (m *MsgCreateInterqueryResultResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateInterqueryResultResponse) ProtoMessage()    {}
func (*MsgCreateInterqueryResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1defee906b3ee117, []int{1}
}
func (m *MsgCreateInterqueryResultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateInterqueryResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateInterqueryResultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateInterqueryResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateInterqueryResultResponse.Merge(m, src)
}
func (m *MsgCreateInterqueryResultResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateInterqueryResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateInterqueryResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateInterqueryResultResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateInterqueryResult)(nil), "defundlabs.defund.query.MsgCreateInterqueryResult")
	proto.RegisterType((*MsgCreateInterqueryResultResponse)(nil), "defundlabs.defund.query.MsgCreateInterqueryResultResponse")
}

func init() { proto.RegisterFile("query/tx.proto", fileDescriptor_1defee906b3ee117) }

var fileDescriptor_1defee906b3ee117 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xbd, 0x4e, 0xeb, 0x30,
	0x18, 0xad, 0x6f, 0x7f, 0xae, 0xae, 0xaf, 0xc4, 0xe0, 0x01, 0x42, 0x10, 0xa1, 0x94, 0xa5, 0x12,
	0xc2, 0x56, 0xc3, 0xc6, 0x08, 0x42, 0x82, 0xa1, 0x02, 0x65, 0x64, 0x4b, 0x9c, 0xaf, 0xa9, 0xa5,
	0x36, 0x0e, 0xb6, 0x53, 0xb5, 0x6f, 0xc0, 0xd8, 0xc7, 0x42, 0x4c, 0x1d, 0x19, 0x51, 0xfb, 0x22,
	0x28, 0x76, 0x2a, 0xa6, 0x0c, 0x6c, 0xe7, 0xcb, 0x39, 0xdf, 0xc9, 0x39, 0xb6, 0xf1, 0xc1, 0x6b,
	0x09, 0x6a, 0xc5, 0xcc, 0x92, 0x16, 0x4a, 0x1a, 0x49, 0x8e, 0x52, 0x98, 0x94, 0x79, 0x3a, 0x8b,
	0x13, 0x4d, 0x1d, 0xa4, 0x56, 0xe1, 0x9f, 0x1a, 0xc8, 0x53, 0x50, 0x73, 0x91, 0x1b, 0xc6, 0xd5,
	0xaa, 0x30, 0x92, 0x15, 0x4a, 0xca, 0x89, 0xdb, 0xf3, 0xcf, 0x44, 0xc2, 0x19, 0x97, 0x0a, 0x18,
	0x9f, 0x09, 0xc8, 0x0d, 0x5b, 0x8c, 0x6a, 0xe4, 0x04, 0x83, 0x0f, 0x84, 0x8f, 0xc7, 0x3a, 0xbb,
	0x53, 0x10, 0x1b, 0x78, 0xcc, 0x0d, 0x28, 0xeb, 0x1b, 0x81, 0x2e, 0x67, 0x86, 0x78, 0xf8, 0x2f,
	0xaf, 0x18, 0xa9, 0x3c, 0xd4, 0x47, 0xc3, 0x7f, 0xd1, 0x7e, 0xac, 0x18, 0x6d, 0xa4, 0x02, 0x91,
	0x7a, 0x7f, 0x1c, 0x53, 0x8f, 0x84, 0xe0, 0x4e, 0x1a, 0x9b, 0xd8, 0x6b, 0xdb, 0xcf, 0x16, 0x93,
	0x10, 0xf7, 0xa6, 0x20, 0xb2, 0xa9, 0xf1, 0x3a, 0x7d, 0x34, 0xfc, 0x1f, 0xfa, 0x54, 0x24, 0x9c,
	0x56, 0xb9, 0x68, 0x9d, 0x66, 0x31, 0xa2, 0x0f, 0x56, 0x11, 0xd5, 0x4a, 0x32, 0xc2, 0x5d, 0xdb,
	0xc4, 0xeb, 0xda, 0x95, 0x13, 0xfa, 0xd3, 0x94, 0xba, 0xa6, 0xf4, 0xb9, 0xe2, 0x9f, 0x0a, 0x1d,
	0x39, 0xe5, 0xe0, 0x02, 0x9f, 0x37, 0x76, 0x89, 0x40, 0x17, 0x32, 0xd7, 0x10, 0xae, 0x11, 0x6e,
	0x8f, 0x75, 0x46, 0xde, 0x10, 0x3e, 0x6c, 0xa8, 0x1d, 0xd2, 0x86, 0xe3, 0xa6, 0x8d, 0xf6, 0xfe,
	0xcd, 0xef, 0x77, 0xf6, 0x91, 0x6e, 0xef, 0xdf, 0xb7, 0x01, 0xda, 0x6c, 0x03, 0xf4, 0xb5, 0x0d,
	0xd0, 0x7a, 0x17, 0xb4, 0x36, 0xbb, 0xa0, 0xf5, 0xb9, 0x0b, 0x5a, 0x2f, 0x97, 0x99, 0x30, 0xd3,
	0x32, 0xa1, 0x5c, 0xce, 0x99, 0x33, 0xbd, 0xaa, 0x7e, 0x50, 0x63, 0xb6, 0x64, 0xf5, 0x3b, 0x59,
	0x15, 0xa0, 0x93, 0x9e, 0xbd, 0xd2, 0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x33, 0xbc,
	0x1c, 0x3d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateInterqueryResult(ctx context.Context, in *MsgCreateInterqueryResult, opts ...grpc.CallOption) (*MsgCreateInterqueryResultResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateInterqueryResult(ctx context.Context, in *MsgCreateInterqueryResult, opts ...grpc.CallOption) (*MsgCreateInterqueryResultResponse, error) {
	out := new(MsgCreateInterqueryResultResponse)
	err := c.cc.Invoke(ctx, "/defundlabs.defund.query.Msg/CreateInterqueryResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateInterqueryResult(context.Context, *MsgCreateInterqueryResult) (*MsgCreateInterqueryResultResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateInterqueryResult(ctx context.Context, req *MsgCreateInterqueryResult) (*MsgCreateInterqueryResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInterqueryResult not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateInterqueryResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateInterqueryResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateInterqueryResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/defundlabs.defund.query.Msg/CreateInterqueryResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateInterqueryResult(ctx, req.(*MsgCreateInterqueryResult))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "defundlabs.defund.query.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInterqueryResult",
			Handler:    _Msg_CreateInterqueryResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query/tx.proto",
}

func (m *MsgCreateInterqueryResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateInterqueryResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateInterqueryResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Height != nil {
		{
			size, err := m.Height.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Storeid) > 0 {
		i -= len(m.Storeid)
		copy(dAtA[i:], m.Storeid)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Storeid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateInterqueryResultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateInterqueryResultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateInterqueryResultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateInterqueryResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Storeid)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Height != nil {
		l = m.Height.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateInterqueryResultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateInterqueryResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateInterqueryResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateInterqueryResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storeid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storeid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Height == nil {
				m.Height = &types.Height{}
			}
			if err := m.Height.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &crypto.ProofOps{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateInterqueryResultResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateInterqueryResultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateInterqueryResultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
