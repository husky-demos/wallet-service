package service

import (
	context "context"
	v1 "wallet-service/pb/wallet-service"
)

type WalletService struct {
}

func (s *WalletService) QueryByUserId(ctx context.Context, userId *v1.Id) (*v1.Wallet, error) {
	println("UserId:", userId.Value)
	return &v1.Wallet{
		Id:               "123",
		UserId:           userId.Value,
		CoinTypeId:       "1",
		AvailableBalance: 799,
		FreezeBalance:    1,
		IsLocking:        false,
	}, nil
}
