// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.26.1
// source: builder/v1/tx.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TxInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain      string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	From       string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To         string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Token      string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Amount     string `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Data       string `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Memo       string `protobuf:"bytes,7,opt,name=memo,proto3" json:"memo,omitempty"`
	Passphrase string `protobuf:"bytes,8,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	MaxAmount  bool   `protobuf:"varint,9,opt,name=maxAmount,proto3" json:"maxAmount,omitempty"`
}

func (x *TxInfoRequest) Reset() {
	*x = TxInfoRequest{}
	mi := &file_builder_v1_tx_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TxInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxInfoRequest) ProtoMessage() {}

func (x *TxInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_builder_v1_tx_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxInfoRequest.ProtoReflect.Descriptor instead.
func (*TxInfoRequest) Descriptor() ([]byte, []int) {
	return file_builder_v1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *TxInfoRequest) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *TxInfoRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TxInfoRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TxInfoRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TxInfoRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *TxInfoRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *TxInfoRequest) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *TxInfoRequest) GetPassphrase() string {
	if x != nil {
		return x.Passphrase
	}
	return ""
}

func (x *TxInfoRequest) GetMaxAmount() bool {
	if x != nil {
		return x.MaxAmount
	}
	return false
}

type BuildTxReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxInput string `protobuf:"bytes,1,opt,name=txInput,proto3" json:"txInput,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BuildTxReply) Reset() {
	*x = BuildTxReply{}
	mi := &file_builder_v1_tx_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BuildTxReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildTxReply) ProtoMessage() {}

func (x *BuildTxReply) ProtoReflect() protoreflect.Message {
	mi := &file_builder_v1_tx_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildTxReply.ProtoReflect.Descriptor instead.
func (*BuildTxReply) Descriptor() ([]byte, []int) {
	return file_builder_v1_tx_proto_rawDescGZIP(), []int{1}
}

func (x *BuildTxReply) GetTxInput() string {
	if x != nil {
		return x.TxInput
	}
	return ""
}

func (x *BuildTxReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SignTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From       string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Passphrase string `protobuf:"bytes,2,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	TxInput    string `protobuf:"bytes,3,opt,name=txInput,proto3" json:"txInput,omitempty"`
}

func (x *SignTxRequest) Reset() {
	*x = SignTxRequest{}
	mi := &file_builder_v1_tx_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignTxRequest) ProtoMessage() {}

func (x *SignTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_builder_v1_tx_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignTxRequest.ProtoReflect.Descriptor instead.
func (*SignTxRequest) Descriptor() ([]byte, []int) {
	return file_builder_v1_tx_proto_rawDescGZIP(), []int{2}
}

func (x *SignTxRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SignTxRequest) GetPassphrase() string {
	if x != nil {
		return x.Passphrase
	}
	return ""
}

func (x *SignTxRequest) GetTxInput() string {
	if x != nil {
		return x.TxInput
	}
	return ""
}

type SignTxReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawTx string `protobuf:"bytes,1,opt,name=rawTx,proto3" json:"rawTx,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SignTxReply) Reset() {
	*x = SignTxReply{}
	mi := &file_builder_v1_tx_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignTxReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignTxReply) ProtoMessage() {}

func (x *SignTxReply) ProtoReflect() protoreflect.Message {
	mi := &file_builder_v1_tx_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignTxReply.ProtoReflect.Descriptor instead.
func (*SignTxReply) Descriptor() ([]byte, []int) {
	return file_builder_v1_tx_proto_rawDescGZIP(), []int{3}
}

func (x *SignTxReply) GetRawTx() string {
	if x != nil {
		return x.RawTx
	}
	return ""
}

func (x *SignTxReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SendRawTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	RawTx string `protobuf:"bytes,2,opt,name=rawTx,proto3" json:"rawTx,omitempty"`
}

func (x *SendRawTxRequest) Reset() {
	*x = SendRawTxRequest{}
	mi := &file_builder_v1_tx_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendRawTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRawTxRequest) ProtoMessage() {}

func (x *SendRawTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_builder_v1_tx_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRawTxRequest.ProtoReflect.Descriptor instead.
func (*SendRawTxRequest) Descriptor() ([]byte, []int) {
	return file_builder_v1_tx_proto_rawDescGZIP(), []int{4}
}

func (x *SendRawTxRequest) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *SendRawTxRequest) GetRawTx() string {
	if x != nil {
		return x.RawTx
	}
	return ""
}

type SendRawTxReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash string `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SendRawTxReply) Reset() {
	*x = SendRawTxReply{}
	mi := &file_builder_v1_tx_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendRawTxReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRawTxReply) ProtoMessage() {}

func (x *SendRawTxReply) ProtoReflect() protoreflect.Message {
	mi := &file_builder_v1_tx_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRawTxReply.ProtoReflect.Descriptor instead.
func (*SendRawTxReply) Descriptor() ([]byte, []int) {
	return file_builder_v1_tx_proto_rawDescGZIP(), []int{5}
}

func (x *SendRawTxReply) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *SendRawTxReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type AssetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Symbol   string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Balance  string `protobuf:"bytes,5,opt,name=balance,proto3" json:"balance,omitempty"`
	Decimals uint32 `protobuf:"varint,6,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (x *AssetInfo) Reset() {
	*x = AssetInfo{}
	mi := &file_builder_v1_tx_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetInfo) ProtoMessage() {}

func (x *AssetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_builder_v1_tx_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetInfo.ProtoReflect.Descriptor instead.
func (*AssetInfo) Descriptor() ([]byte, []int) {
	return file_builder_v1_tx_proto_rawDescGZIP(), []int{6}
}

func (x *AssetInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AssetInfo) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *AssetInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AssetInfo) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *AssetInfo) GetDecimals() uint32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

type AssetList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner  string       `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Assets []*AssetInfo `protobuf:"bytes,2,rep,name=assets,proto3" json:"assets,omitempty"`
}

func (x *AssetList) Reset() {
	*x = AssetList{}
	mi := &file_builder_v1_tx_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetList) ProtoMessage() {}

func (x *AssetList) ProtoReflect() protoreflect.Message {
	mi := &file_builder_v1_tx_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetList.ProtoReflect.Descriptor instead.
func (*AssetList) Descriptor() ([]byte, []int) {
	return file_builder_v1_tx_proto_rawDescGZIP(), []int{7}
}

func (x *AssetList) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *AssetList) GetAssets() []*AssetInfo {
	if x != nil {
		return x.Assets
	}
	return nil
}

type BalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain      string       `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	AssetGroup []*AssetList `protobuf:"bytes,2,rep,name=assetGroup,proto3" json:"assetGroup,omitempty"`
}

func (x *BalanceRequest) Reset() {
	*x = BalanceRequest{}
	mi := &file_builder_v1_tx_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceRequest) ProtoMessage() {}

func (x *BalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_builder_v1_tx_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceRequest.ProtoReflect.Descriptor instead.
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return file_builder_v1_tx_proto_rawDescGZIP(), []int{8}
}

func (x *BalanceRequest) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *BalanceRequest) GetAssetGroup() []*AssetList {
	if x != nil {
		return x.AssetGroup
	}
	return nil
}

type BalanceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetGroup []*AssetList `protobuf:"bytes,1,rep,name=assetGroup,proto3" json:"assetGroup,omitempty"`
}

func (x *BalanceReply) Reset() {
	*x = BalanceReply{}
	mi := &file_builder_v1_tx_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalanceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceReply) ProtoMessage() {}

func (x *BalanceReply) ProtoReflect() protoreflect.Message {
	mi := &file_builder_v1_tx_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceReply.ProtoReflect.Descriptor instead.
func (*BalanceReply) Descriptor() ([]byte, []int) {
	return file_builder_v1_tx_proto_rawDescGZIP(), []int{9}
}

func (x *BalanceReply) GetAssetGroup() []*AssetList {
	if x != nil {
		return x.AssetGroup
	}
	return nil
}

var File_builder_v1_tx_proto protoreflect.FileDescriptor

var file_builder_v1_tx_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x0d, 0x54, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72,
	0x61, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x3e, 0x0a, 0x0c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x78, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x78, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x5d, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x70, 0x68,
	0x72, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x73, 0x73,
	0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x78, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x22, 0x39, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x61, 0x77, 0x54, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x72, 0x61, 0x77, 0x54, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3e, 0x0a, 0x10, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x61, 0x77, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x77, 0x54, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x77, 0x54, 0x78, 0x22, 0x3e, 0x0a, 0x0e, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x61, 0x77, 0x54, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x83, 0x01, 0x0a, 0x09,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x73, 0x22, 0x54, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x22, 0x61, 0x0a, 0x0e, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12,
	0x39, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x49, 0x0a, 0x0c, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x32, 0xe7, 0x06, 0x0a, 0x02, 0x54, 0x78, 0x12, 0x5c, 0x0a, 0x07,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x78, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x78, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22,
	0x09, 0x2f, 0x74, 0x78, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x59, 0x0a, 0x06, 0x53, 0x69,
	0x67, 0x6e, 0x54, 0x78, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x74, 0x78,
	0x2f, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x66, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x61, 0x77,
	0x54, 0x78, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x61, 0x77, 0x54, 0x78, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x61, 0x77, 0x54, 0x78, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22,
	0x0c, 0x2f, 0x74, 0x78, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x61, 0x77, 0x12, 0x5c, 0x0a,
	0x06, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x61, 0x77, 0x54,
	0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01,
	0x2a, 0x22, 0x08, 0x2f, 0x74, 0x78, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x62, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a,
	0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x74, 0x78, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x86, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x62, 0x79, 0x68, 0x61, 0x73, 0x68, 0x12, 0x75, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x7e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x42,
	0x30, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x1c, 0x74, 0x78, 0x2d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_builder_v1_tx_proto_rawDescOnce sync.Once
	file_builder_v1_tx_proto_rawDescData = file_builder_v1_tx_proto_rawDesc
)

func file_builder_v1_tx_proto_rawDescGZIP() []byte {
	file_builder_v1_tx_proto_rawDescOnce.Do(func() {
		file_builder_v1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_builder_v1_tx_proto_rawDescData)
	})
	return file_builder_v1_tx_proto_rawDescData
}

var file_builder_v1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_builder_v1_tx_proto_goTypes = []any{
	(*TxInfoRequest)(nil),               // 0: api.builder.v1.TxInfoRequest
	(*BuildTxReply)(nil),                // 1: api.builder.v1.BuildTxReply
	(*SignTxRequest)(nil),               // 2: api.builder.v1.SignTxRequest
	(*SignTxReply)(nil),                 // 3: api.builder.v1.SignTxReply
	(*SendRawTxRequest)(nil),            // 4: api.builder.v1.SendRawTxRequest
	(*SendRawTxReply)(nil),              // 5: api.builder.v1.SendRawTxReply
	(*AssetInfo)(nil),                   // 6: api.builder.v1.AssetInfo
	(*AssetList)(nil),                   // 7: api.builder.v1.AssetList
	(*BalanceRequest)(nil),              // 8: api.builder.v1.BalanceRequest
	(*BalanceReply)(nil),                // 9: api.builder.v1.BalanceReply
	(*GetTransactionByHashRequest)(nil), // 10: api.builder.v1.GetTransactionByHashRequest
	(*PageListRequest)(nil),             // 11: api.builder.v1.PageListRequest
	(*PageListAssetRequest)(nil),        // 12: api.builder.v1.PageListAssetRequest
	(*TransactionRecord)(nil),           // 13: api.builder.v1.TransactionRecord
	(*PageListResponse)(nil),            // 14: api.builder.v1.PageListResponse
	(*PageListAssetResponse)(nil),       // 15: api.builder.v1.PageListAssetResponse
}
var file_builder_v1_tx_proto_depIdxs = []int32{
	6,  // 0: api.builder.v1.AssetList.assets:type_name -> api.builder.v1.AssetInfo
	7,  // 1: api.builder.v1.BalanceRequest.assetGroup:type_name -> api.builder.v1.AssetList
	7,  // 2: api.builder.v1.BalanceReply.assetGroup:type_name -> api.builder.v1.AssetList
	0,  // 3: api.builder.v1.Tx.BuildTx:input_type -> api.builder.v1.TxInfoRequest
	2,  // 4: api.builder.v1.Tx.SignTx:input_type -> api.builder.v1.SignTxRequest
	4,  // 5: api.builder.v1.Tx.SendRawTx:input_type -> api.builder.v1.SendRawTxRequest
	0,  // 6: api.builder.v1.Tx.SendTx:input_type -> api.builder.v1.TxInfoRequest
	8,  // 7: api.builder.v1.Tx.GetBalance:input_type -> api.builder.v1.BalanceRequest
	10, // 8: api.builder.v1.Tx.GetTransactionByHash:input_type -> api.builder.v1.GetTransactionByHashRequest
	11, // 9: api.builder.v1.Tx.GetTransactionList:input_type -> api.builder.v1.PageListRequest
	12, // 10: api.builder.v1.Tx.GetAssetList:input_type -> api.builder.v1.PageListAssetRequest
	1,  // 11: api.builder.v1.Tx.BuildTx:output_type -> api.builder.v1.BuildTxReply
	3,  // 12: api.builder.v1.Tx.SignTx:output_type -> api.builder.v1.SignTxReply
	5,  // 13: api.builder.v1.Tx.SendRawTx:output_type -> api.builder.v1.SendRawTxReply
	5,  // 14: api.builder.v1.Tx.SendTx:output_type -> api.builder.v1.SendRawTxReply
	9,  // 15: api.builder.v1.Tx.GetBalance:output_type -> api.builder.v1.BalanceReply
	13, // 16: api.builder.v1.Tx.GetTransactionByHash:output_type -> api.builder.v1.TransactionRecord
	14, // 17: api.builder.v1.Tx.GetTransactionList:output_type -> api.builder.v1.PageListResponse
	15, // 18: api.builder.v1.Tx.GetAssetList:output_type -> api.builder.v1.PageListAssetResponse
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_builder_v1_tx_proto_init() }
func file_builder_v1_tx_proto_init() {
	if File_builder_v1_tx_proto != nil {
		return
	}
	file_builder_v1_transaction_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_builder_v1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_builder_v1_tx_proto_goTypes,
		DependencyIndexes: file_builder_v1_tx_proto_depIdxs,
		MessageInfos:      file_builder_v1_tx_proto_msgTypes,
	}.Build()
	File_builder_v1_tx_proto = out.File
	file_builder_v1_tx_proto_rawDesc = nil
	file_builder_v1_tx_proto_goTypes = nil
	file_builder_v1_tx_proto_depIdxs = nil
}
