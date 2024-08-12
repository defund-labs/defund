// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: defund/dex/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	cosmossdk_io_store_types "cosmossdk.io/store/types"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the liquidity module.
type Params struct {
	BatchSize                       uint32                                   `protobuf:"varint,1,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	TickPrecision                   uint32                                   `protobuf:"varint,2,opt,name=tick_precision,json=tickPrecision,proto3" json:"tick_precision,omitempty"`
	FeeCollectorAddress             string                                   `protobuf:"bytes,3,opt,name=fee_collector_address,json=feeCollectorAddress,proto3" json:"fee_collector_address,omitempty"`
	DustCollectorAddress            string                                   `protobuf:"bytes,4,opt,name=dust_collector_address,json=dustCollectorAddress,proto3" json:"dust_collector_address,omitempty"`
	MinInitialPoolCoinSupply        cosmossdk_io_math.Int                    `protobuf:"bytes,5,opt,name=min_initial_pool_coin_supply,json=minInitialPoolCoinSupply,proto3,customtype=cosmossdk.io/math.Int" json:"min_initial_pool_coin_supply"`
	PairCreationFee                 github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=pair_creation_fee,json=pairCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pair_creation_fee"`
	PoolCreationFee                 github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=pool_creation_fee,json=poolCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pool_creation_fee"`
	MinInitialDepositAmount         cosmossdk_io_math.Int                    `protobuf:"bytes,8,opt,name=min_initial_deposit_amount,json=minInitialDepositAmount,proto3,customtype=cosmossdk.io/math.Int" json:"min_initial_deposit_amount"`
	MaxPriceLimitRatio              cosmossdk_io_math.LegacyDec              `protobuf:"bytes,9,opt,name=max_price_limit_ratio,json=maxPriceLimitRatio,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"max_price_limit_ratio"`
	MaxNumMarketMakingOrderTicks    uint32                                   `protobuf:"varint,10,opt,name=max_num_market_making_order_ticks,json=maxNumMarketMakingOrderTicks,proto3" json:"max_num_market_making_order_ticks,omitempty"`
	MaxNumMarketMakingOrdersPerPair uint32                                   `protobuf:"varint,11,opt,name=max_num_market_making_orders_per_pair,json=maxNumMarketMakingOrdersPerPair,proto3" json:"max_num_market_making_orders_per_pair,omitempty"`
	MaxOrderLifespan                time.Duration                            `protobuf:"bytes,12,opt,name=max_order_lifespan,json=maxOrderLifespan,proto3,stdduration" json:"max_order_lifespan"`
	SwapFeeRate                     cosmossdk_io_math.LegacyDec              `protobuf:"bytes,13,opt,name=swap_fee_rate,json=swapFeeRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"swap_fee_rate"`
	WithdrawFeeRate                 cosmossdk_io_math.LegacyDec              `protobuf:"bytes,14,opt,name=withdraw_fee_rate,json=withdrawFeeRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"withdraw_fee_rate"`
	DepositExtraGas                 cosmossdk_io_store_types.Gas             `protobuf:"varint,15,opt,name=deposit_extra_gas,json=depositExtraGas,proto3,customtype=cosmossdk.io/store/types.Gas" json:"deposit_extra_gas"`
	WithdrawExtraGas                cosmossdk_io_store_types.Gas             `protobuf:"varint,16,opt,name=withdraw_extra_gas,json=withdrawExtraGas,proto3,customtype=cosmossdk.io/store/types.Gas" json:"withdraw_extra_gas"`
	OrderExtraGas                   cosmossdk_io_store_types.Gas             `protobuf:"varint,17,opt,name=order_extra_gas,json=orderExtraGas,proto3,customtype=cosmossdk.io/store/types.Gas" json:"order_extra_gas"`
	MaxNumActivePoolsPerPair        uint32                                   `protobuf:"varint,18,opt,name=max_num_active_pools_per_pair,json=maxNumActivePoolsPerPair,proto3" json:"max_num_active_pools_per_pair,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_726cd32d4b7083e4, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetBatchSize() uint32 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *Params) GetTickPrecision() uint32 {
	if m != nil {
		return m.TickPrecision
	}
	return 0
}

func (m *Params) GetFeeCollectorAddress() string {
	if m != nil {
		return m.FeeCollectorAddress
	}
	return ""
}

func (m *Params) GetDustCollectorAddress() string {
	if m != nil {
		return m.DustCollectorAddress
	}
	return ""
}

func (m *Params) GetPairCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PairCreationFee
	}
	return nil
}

func (m *Params) GetPoolCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PoolCreationFee
	}
	return nil
}

func (m *Params) GetMaxNumMarketMakingOrderTicks() uint32 {
	if m != nil {
		return m.MaxNumMarketMakingOrderTicks
	}
	return 0
}

func (m *Params) GetMaxNumMarketMakingOrdersPerPair() uint32 {
	if m != nil {
		return m.MaxNumMarketMakingOrdersPerPair
	}
	return 0
}

func (m *Params) GetMaxOrderLifespan() time.Duration {
	if m != nil {
		return m.MaxOrderLifespan
	}
	return 0
}

func (m *Params) GetMaxNumActivePoolsPerPair() uint32 {
	if m != nil {
		return m.MaxNumActivePoolsPerPair
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "defund.dex.Params")
}

func init() { proto.RegisterFile("defund/dex/params.proto", fileDescriptor_726cd32d4b7083e4) }

var fileDescriptor_726cd32d4b7083e4 = []byte{
	// 822 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0x4f, 0x6f, 0xdb, 0x36,
	0x18, 0xc6, 0xa3, 0xb5, 0xcb, 0x1a, 0x66, 0xae, 0x63, 0xb6, 0x59, 0xd4, 0x2c, 0xb1, 0xbd, 0x3f,
	0x05, 0x8c, 0x61, 0x93, 0x96, 0x6e, 0xa7, 0x5d, 0x86, 0x38, 0x59, 0x83, 0x02, 0x69, 0xeb, 0xa9,
	0xc3, 0x0e, 0x05, 0x06, 0x82, 0x96, 0x5e, 0xcb, 0x84, 0x45, 0x51, 0x20, 0xa9, 0x58, 0xe9, 0x47,
	0xd8, 0x69, 0xd8, 0x69, 0x9f, 0x61, 0xa7, 0x7d, 0x8c, 0x1e, 0x7b, 0x1c, 0x76, 0x68, 0x87, 0xe4,
	0xb0, 0xaf, 0x31, 0x90, 0xa2, 0x62, 0x17, 0xdd, 0x86, 0xe4, 0xb0, 0x8b, 0x65, 0xf1, 0x7d, 0x9f,
	0xdf, 0x4b, 0xbe, 0x7c, 0x48, 0xa1, 0xad, 0x04, 0x26, 0x65, 0x9e, 0x84, 0x09, 0x54, 0x61, 0x41,
	0x25, 0xe5, 0x2a, 0x28, 0xa4, 0xd0, 0x02, 0xa3, 0x3a, 0x10, 0x24, 0x50, 0x6d, 0x77, 0x28, 0x67,
	0xb9, 0x08, 0xed, 0x6f, 0x1d, 0xde, 0xde, 0x8a, 0x85, 0xe2, 0x42, 0x85, 0x5c, 0xa5, 0xe1, 0xc9,
	0x9e, 0x79, 0xb8, 0xc0, 0x9d, 0x3a, 0x40, 0xec, 0x5b, 0x58, 0xbf, 0xb8, 0xd0, 0xed, 0x54, 0xa4,
	0xa2, 0x1e, 0x37, 0xff, 0xdc, 0x68, 0xd7, 0x91, 0xc6, 0x54, 0x41, 0x78, 0xb2, 0x37, 0x06, 0x4d,
	0xf7, 0xc2, 0x58, 0xb0, 0xbc, 0x89, 0xa7, 0x42, 0xa4, 0x19, 0x84, 0xf6, 0x6d, 0x5c, 0x4e, 0xc2,
	0xa4, 0x94, 0x54, 0x33, 0xe1, 0xe2, 0x1f, 0xfe, 0xbc, 0x8e, 0x56, 0x47, 0x76, 0xe6, 0x78, 0x17,
	0xa1, 0x31, 0xd5, 0xf1, 0x94, 0x28, 0xf6, 0x0c, 0x7c, 0xaf, 0xef, 0x0d, 0x5a, 0xd1, 0x9a, 0x1d,
	0x79, 0xc2, 0x9e, 0x01, 0xbe, 0x8b, 0x6e, 0x6a, 0x16, 0xcf, 0x48, 0x21, 0x21, 0x66, 0x8a, 0x89,
	0xdc, 0x7f, 0xcb, 0xa6, 0xb4, 0xcc, 0xe8, 0xa8, 0x19, 0xc4, 0xf7, 0xd0, 0xe6, 0x04, 0x80, 0xc4,
	0x22, 0xcb, 0x20, 0xd6, 0x42, 0x12, 0x9a, 0x24, 0x12, 0x94, 0xf2, 0xaf, 0xf5, 0xbd, 0xc1, 0x5a,
	0x74, 0x6b, 0x02, 0x70, 0xd0, 0xc4, 0xf6, 0xeb, 0x10, 0xfe, 0x12, 0xbd, 0x97, 0x94, 0x4a, 0xff,
	0x83, 0xe8, 0xba, 0x15, 0xdd, 0x36, 0xd1, 0x37, 0x54, 0x3f, 0xa0, 0x1d, 0xce, 0x72, 0xc2, 0x72,
	0xa6, 0x19, 0xcd, 0x48, 0x21, 0x44, 0x46, 0xcc, 0xca, 0x89, 0x2a, 0x8b, 0x22, 0x3b, 0xf5, 0xdf,
	0x36, 0xda, 0xe1, 0xee, 0xf3, 0x97, 0xbd, 0x95, 0x3f, 0x5e, 0xf6, 0x36, 0xeb, 0x46, 0xa9, 0x64,
	0x16, 0x30, 0x11, 0x72, 0xaa, 0xa7, 0xc1, 0x83, 0x5c, 0x47, 0x3e, 0x67, 0xf9, 0x83, 0x9a, 0x30,
	0x12, 0x22, 0x3b, 0x10, 0x2c, 0x7f, 0x62, 0xe5, 0x78, 0x8e, 0x3a, 0x05, 0x65, 0x92, 0xc4, 0x12,
	0x6c, 0xc3, 0xc8, 0x04, 0xc0, 0x5f, 0xed, 0x5f, 0x1b, 0xac, 0xdf, 0xbb, 0x13, 0xb8, 0x9d, 0x31,
	0x5d, 0x0f, 0x5c, 0xd7, 0x03, 0xa3, 0x1d, 0x7e, 0x6e, 0xca, 0xfd, 0xfa, 0xaa, 0x37, 0x48, 0x99,
	0x9e, 0x96, 0xe3, 0x20, 0x16, 0xdc, 0x6d, 0xa3, 0x7b, 0x7c, 0xa6, 0x92, 0x59, 0xa8, 0x4f, 0x0b,
	0x50, 0x56, 0xa0, 0xa2, 0xb6, 0xa9, 0x72, 0xe0, 0x8a, 0xdc, 0x07, 0xb0, 0x85, 0xed, 0x5a, 0x96,
	0x0b, 0xbf, 0xf3, 0x7f, 0x14, 0x36, 0x0b, 0x5e, 0x2a, 0xfc, 0x14, 0x6d, 0x2f, 0x37, 0x34, 0x81,
	0x42, 0x28, 0xa6, 0x09, 0xe5, 0xa2, 0xcc, 0xb5, 0x7f, 0xe3, 0x32, 0xed, 0xdc, 0x5a, 0xb4, 0xf3,
	0xb0, 0x96, 0xef, 0x5b, 0x35, 0xfe, 0x1e, 0x6d, 0x72, 0x5a, 0x91, 0x42, 0xb2, 0x18, 0x48, 0xc6,
	0x38, 0xd3, 0xc4, 0xfa, 0xd0, 0x5f, 0xb3, 0xd8, 0x8f, 0x1c, 0xf6, 0xfd, 0x37, 0xb1, 0xc7, 0x90,
	0xd2, 0xf8, 0xf4, 0x10, 0xe2, 0x08, 0x73, 0x5a, 0x8d, 0x0c, 0xe0, 0xd8, 0xe8, 0x23, 0x23, 0xc7,
	0x47, 0xe8, 0x03, 0xc3, 0xcd, 0x4b, 0x4e, 0x38, 0x95, 0x33, 0xd0, 0x84, 0xd3, 0x19, 0xcb, 0x53,
	0x22, 0x64, 0x02, 0x92, 0x18, 0x6f, 0x2a, 0x1f, 0x59, 0xa3, 0xee, 0x70, 0x5a, 0x3d, 0x2a, 0xf9,
	0x43, 0x9b, 0xf6, 0xd0, 0x66, 0x3d, 0x36, 0x49, 0xdf, 0x99, 0x1c, 0xfc, 0x08, 0xdd, 0xfd, 0x0f,
	0x90, 0x22, 0x05, 0x48, 0x62, 0x76, 0xca, 0x5f, 0xb7, 0xb0, 0xde, 0xbf, 0xc0, 0xd4, 0x08, 0xe4,
	0x88, 0x32, 0x89, 0xbf, 0x45, 0x66, 0xba, 0x6e, 0x1a, 0x19, 0x9b, 0x80, 0x2a, 0x68, 0xee, 0xbf,
	0xdb, 0xf7, 0xec, 0x36, 0xd6, 0xa7, 0x32, 0x68, 0x4e, 0x65, 0x70, 0xe8, 0x4e, 0xe5, 0xf0, 0x86,
	0x69, 0xc4, 0x2f, 0xaf, 0x7a, 0x5e, 0xb4, 0xc1, 0x69, 0x65, 0x91, 0xc7, 0x4e, 0x8c, 0x8f, 0x50,
	0x4b, 0xcd, 0x69, 0x61, 0xfc, 0x60, 0x9a, 0x07, 0x7e, 0xeb, 0xf2, 0xbd, 0x5b, 0x37, 0xca, 0xfb,
	0x00, 0x11, 0xd5, 0x80, 0x1f, 0xa3, 0xce, 0x9c, 0xe9, 0x69, 0x22, 0xe9, 0x7c, 0x01, 0xbb, 0x79,
	0x79, 0x58, 0xbb, 0x51, 0x37, 0xc0, 0x11, 0xea, 0x34, 0x6e, 0x81, 0x4a, 0x4b, 0x4a, 0x52, 0xaa,
	0xfc, 0x76, 0xdf, 0x1b, 0x5c, 0x1f, 0x7e, 0xec, 0x80, 0x3b, 0xaf, 0x01, 0x95, 0x16, 0x12, 0x9c,
	0x1b, 0x8f, 0xa8, 0x8a, 0xda, 0x4e, 0xfe, 0x8d, 0x51, 0x1f, 0x51, 0x85, 0x23, 0x84, 0x2f, 0xa6,
	0xb8, 0x40, 0x6e, 0x5c, 0x01, 0xb9, 0xd1, 0xe8, 0x2f, 0x98, 0xc7, 0xa8, 0x5d, 0x6f, 0xc7, 0x02,
	0xd8, 0xb9, 0x02, 0xb0, 0x65, 0xc5, 0x17, 0xb4, 0xaf, 0xd1, 0x6e, 0x63, 0x18, 0x1a, 0x6b, 0x76,
	0x02, 0xf6, 0x06, 0x5a, 0x32, 0x0a, 0xb6, 0x46, 0xf1, 0x6b, 0xa3, 0xec, 0xdb, 0x14, 0x73, 0xc5,
	0x34, 0x0e, 0xf9, 0xca, 0xff, 0xf1, 0xaf, 0xdf, 0x3e, 0xb9, 0xe5, 0xbe, 0x20, 0x95, 0xfd, 0x86,
	0xd4, 0x37, 0xf1, 0xf0, 0xd3, 0xe7, 0x67, 0x5d, 0xef, 0xc5, 0x59, 0xd7, 0xfb, 0xf3, 0xac, 0xeb,
	0xfd, 0x74, 0xde, 0x5d, 0x79, 0x71, 0xde, 0x5d, 0xf9, 0xfd, 0xbc, 0xbb, 0xf2, 0x14, 0xbf, 0x96,
	0x6e, 0xe7, 0x36, 0x5e, 0xb5, 0x2e, 0xfa, 0xe2, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x0f,
	0x7a, 0x9d, 0x8d, 0x06, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxNumActivePoolsPerPair != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxNumActivePoolsPerPair))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	if m.OrderExtraGas != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.OrderExtraGas))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.WithdrawExtraGas != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.WithdrawExtraGas))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.DepositExtraGas != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DepositExtraGas))
		i--
		dAtA[i] = 0x78
	}
	{
		size := m.WithdrawFeeRate.Size()
		i -= size
		if _, err := m.WithdrawFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	{
		size := m.SwapFeeRate.Size()
		i -= size
		if _, err := m.SwapFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.MaxOrderLifespan, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxOrderLifespan):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x62
	if m.MaxNumMarketMakingOrdersPerPair != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxNumMarketMakingOrdersPerPair))
		i--
		dAtA[i] = 0x58
	}
	if m.MaxNumMarketMakingOrderTicks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxNumMarketMakingOrderTicks))
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.MaxPriceLimitRatio.Size()
		i -= size
		if _, err := m.MaxPriceLimitRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.MinInitialDepositAmount.Size()
		i -= size
		if _, err := m.MinInitialDepositAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.PoolCreationFee) > 0 {
		for iNdEx := len(m.PoolCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.PairCreationFee) > 0 {
		for iNdEx := len(m.PairCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PairCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size := m.MinInitialPoolCoinSupply.Size()
		i -= size
		if _, err := m.MinInitialPoolCoinSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.DustCollectorAddress) > 0 {
		i -= len(m.DustCollectorAddress)
		copy(dAtA[i:], m.DustCollectorAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DustCollectorAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FeeCollectorAddress) > 0 {
		i -= len(m.FeeCollectorAddress)
		copy(dAtA[i:], m.FeeCollectorAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.FeeCollectorAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TickPrecision != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TickPrecision))
		i--
		dAtA[i] = 0x10
	}
	if m.BatchSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BatchSize))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BatchSize != 0 {
		n += 1 + sovParams(uint64(m.BatchSize))
	}
	if m.TickPrecision != 0 {
		n += 1 + sovParams(uint64(m.TickPrecision))
	}
	l = len(m.FeeCollectorAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DustCollectorAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.MinInitialPoolCoinSupply.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.PairCreationFee) > 0 {
		for _, e := range m.PairCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.PoolCreationFee) > 0 {
		for _, e := range m.PoolCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.MinInitialDepositAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxPriceLimitRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MaxNumMarketMakingOrderTicks != 0 {
		n += 1 + sovParams(uint64(m.MaxNumMarketMakingOrderTicks))
	}
	if m.MaxNumMarketMakingOrdersPerPair != 0 {
		n += 1 + sovParams(uint64(m.MaxNumMarketMakingOrdersPerPair))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxOrderLifespan)
	n += 1 + l + sovParams(uint64(l))
	l = m.SwapFeeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.WithdrawFeeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.DepositExtraGas != 0 {
		n += 1 + sovParams(uint64(m.DepositExtraGas))
	}
	if m.WithdrawExtraGas != 0 {
		n += 2 + sovParams(uint64(m.WithdrawExtraGas))
	}
	if m.OrderExtraGas != 0 {
		n += 2 + sovParams(uint64(m.OrderExtraGas))
	}
	if m.MaxNumActivePoolsPerPair != 0 {
		n += 2 + sovParams(uint64(m.MaxNumActivePoolsPerPair))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchSize", wireType)
			}
			m.BatchSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickPrecision", wireType)
			}
			m.TickPrecision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickPrecision |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeCollectorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeCollectorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DustCollectorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DustCollectorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitialPoolCoinSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinInitialPoolCoinSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairCreationFee = append(m.PairCreationFee, types.Coin{})
			if err := m.PairCreationFee[len(m.PairCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolCreationFee = append(m.PoolCreationFee, types.Coin{})
			if err := m.PoolCreationFee[len(m.PoolCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitialDepositAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinInitialDepositAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPriceLimitRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxPriceLimitRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumMarketMakingOrderTicks", wireType)
			}
			m.MaxNumMarketMakingOrderTicks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumMarketMakingOrderTicks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumMarketMakingOrdersPerPair", wireType)
			}
			m.MaxNumMarketMakingOrdersPerPair = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumMarketMakingOrdersPerPair |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxOrderLifespan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.MaxOrderLifespan, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WithdrawFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositExtraGas", wireType)
			}
			m.DepositExtraGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositExtraGas |= cosmossdk_io_store_types.Gas(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawExtraGas", wireType)
			}
			m.WithdrawExtraGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WithdrawExtraGas |= cosmossdk_io_store_types.Gas(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderExtraGas", wireType)
			}
			m.OrderExtraGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderExtraGas |= cosmossdk_io_store_types.Gas(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumActivePoolsPerPair", wireType)
			}
			m.MaxNumActivePoolsPerPair = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumActivePoolsPerPair |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)