package service

import (
	"github.com/AcordoCertoBR/ac-atende-positivo-api/libs/acmarketplace"
	"github.com/AcordoCertoBR/ac-atende-positivo-api/libs/entities"
	"github.com/AcordoCertoBR/ac-atende-positivo-api/libs/errors"
)

type GetCustomerService struct {
	ACMarketplace *acmarketplace.ACMarketplace
}

func NewGetCustomerService(ACMarketplace *acmarketplace.ACMarketplace) *GetCustomerService {
	return &GetCustomerService{
		ACMarketplace: ACMarketplace,
	}
}

func (s *GetCustomerService) GetCustomer(document string) (retVal entities.ACGetCustomerResponse, err error) {
	retVal, err = s.ACMarketplace.GetCustomer(document)
	if err != nil {
		return retVal, errors.Wrap(err)
	}

	return retVal, nil
}
