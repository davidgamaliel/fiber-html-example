package domain

import (
	"github.com/fiber-html/config"
)

type MerchantUsecase struct {
	Repository MerchantPersister
	Logger     config.Logger
}

type MerchantExecutor interface {
	GetMerchants(req MerchantListRequest) ([]MerchantDataResponse, error)
	GetMerchantByID(id int) (*MerchantDataResponse, error)
}

func NewUseCase(rep MerchantPersister, logger config.Logger) MerchantExecutor {
	return &MerchantUsecase{
		Repository: rep,
		Logger:     logger,
	}
}

func (u *MerchantUsecase) GetMerchants(req MerchantListRequest) ([]MerchantDataResponse, error) {
	var results []MerchantDataResponse

	merchants, err := u.Repository.GetMerchants(req)
	if err != nil {
		return []MerchantDataResponse{}, err
	}

	for _, merch := range merchants {
		results = append(results, MerchantDataResponse{
			ID:        merch.ID,
			Name:      merch.Name,
			StoreName: merch.StoreName,
			Status:    merch.Status.String(),
			PhoneNo:   merch.PhoneNo,
			JoinedAt:  merch.JoinedAt,
		})
	}

	return results, nil
}

func (u *MerchantUsecase) GetMerchantByID(id int) (*MerchantDataResponse, error) {
	var res MerchantDataResponse
	merchant, err := u.Repository.GetMerchantByID(id)
	if err != nil {
		return nil, err
	}
	if merchant == nil {
		return nil, nil
	}

	res.ID = merchant.ID
	res.Name = merchant.Name
	res.StoreName = merchant.StoreName
	res.Status = merchant.Status.String()
	res.PhoneNo = merchant.PhoneNo
	res.JoinedAt = merchant.JoinedAt

	return &res, nil
}
