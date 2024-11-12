package service

import (
	"context"

	pb "tx-builder/api/builder/v1"
	"tx-builder/internal/biz"
)

type TxService struct {
	pb.UnimplementedTxServer
	uc *biz.ChainDataUsecase
}

func NewTxService(uc *biz.ChainDataUsecase) *TxService {
	return &TxService{}
}

func (s *TxService) BuildTx(ctx context.Context, req *pb.TxInfoRequest) (*pb.BuildTxReply, error) {
	txInput, err := biz.BuildTxInput(req.Chain, req.From, req.To, req.Token, req.Amount, nil)
	return &pb.BuildTxReply{TxInput: txInput, Error: err.Error()}, nil
}
func (s *TxService) SendRawTx(ctx context.Context, req *pb.SendRawTxRequest) (*pb.SendRawTxReply, error) {
	txHash, err := biz.SendRawTx(req.Chain, req.RawTx, nil)
	return &pb.SendRawTxReply{TxHash: txHash, Error: err.Error()}, nil
}
func (s *TxService) SendTx(ctx context.Context, req *pb.TxInfoRequest) (*pb.SendRawTxReply, error) {
	txHash, err := biz.SendTx(req.Chain, req.From, req.To, req.Token, req.Amount, req.Passphrase)
	return &pb.SendRawTxReply{TxHash: txHash, Error: err.Error()}, nil
}
func (s *TxService) GetBalance(ctx context.Context, req *pb.BalanceRequest) (*pb.BalanceReply, error) {
	return &pb.BalanceReply{}, nil
}
