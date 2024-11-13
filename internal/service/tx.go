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
	txInput, err := biz.BuildTxInput(req.Chain, req.From, req.To, req.Token, req.Amount, req.MaxAmount, nil)
	if err != nil {
		return &pb.BuildTxReply{Error: err.Error()}, nil
	}
	return &pb.BuildTxReply{TxInput: txInput}, nil
}
func (s *TxService) SignTx(ctx context.Context, req *pb.SignTxRequest) (*pb.SignTxReply, error) {
	rawTx, err := biz.SignTx(req.From, req.TxInput, req.Passphrase)
	if err != nil {
		return &pb.SignTxReply{Error: err.Error()}, nil
	}
	return &pb.SignTxReply{RawTx: rawTx}, nil
}
func (s *TxService) SendRawTx(ctx context.Context, req *pb.SendRawTxRequest) (*pb.SendRawTxReply, error) {
	txHash, err := biz.SendRawTx(req.Chain, req.RawTx, nil)
	if err != nil {
		return &pb.SendRawTxReply{Error: err.Error()}, nil
	}
	return &pb.SendRawTxReply{TxHash: txHash}, nil
}
func (s *TxService) SendTx(ctx context.Context, req *pb.TxInfoRequest) (*pb.SendRawTxReply, error) {
	txHash, err := biz.SendTx(req.Chain, req.From, req.To, req.Token, req.Amount, req.Passphrase, req.MaxAmount)
	if err != nil {
		return &pb.SendRawTxReply{Error: err.Error()}, nil
	}
	return &pb.SendRawTxReply{TxHash: txHash}, nil
}
func (s *TxService) GetBalance(ctx context.Context, req *pb.BalanceRequest) (*pb.BalanceReply, error) {
	assets, err := biz.GetBalance(req.Chain, req.AssetGroup)
	if err != nil {
		return &pb.BalanceReply{}, nil
	}
	return &pb.BalanceReply{AssetGroup: assets}, nil
}
