package service

import (
	context "context"
	v1 "wallet-service/pb/wallet-service"
	"wallet-service/utils"
)

type WalletService struct {
}

func (s *WalletService) QueryByUserId(ctx context.Context, userId *v1.Id) (*v1.Wallet, error) {
	println("UserId:", userId.Value)
	if userId.Value != "123" {
		return nil, utils.NewError(100, "用户不存在")
	} else {
		return &v1.Wallet{
			Id:               "123",
			UserId:           userId.Value,
			CoinTypeId:       "1",
			AvailableBalance: 799,
			FreezeBalance:    1,
			IsLocking:        false,
		}, nil
	}
}
