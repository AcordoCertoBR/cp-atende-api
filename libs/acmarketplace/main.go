package acmarketplace

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AcordoCertoBR/cp-atende-api/libs/config"
	"github.com/AcordoCertoBR/cp-atende-api/libs/entities"
	"github.com/AcordoCertoBR/cp-atende-api/libs/errors"
	httpUtils "github.com/AcordoCertoBR/cp-atende-api/libs/http"
)

type ACMarketplace struct {
	httpClient *httpUtils.Http
	cfg        *config.Config
}

func NewACMarkeplace(httpClient *httpUtils.Http, cfg *config.Config) *ACMarketplace {
	return &ACMarketplace{
		httpClient: httpClient,
		cfg:        cfg,
	}
}

func (a *ACMarketplace) GetCustomer(document string) (retVal entities.ACGetCustomerResponse, err error) {
	url := fmt.Sprintf("%s/marketplace/v1/internal/customer/%s", a.cfg.ACMarketplaceApiConfig.Host, document)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return retVal, errors.Wrap(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", a.cfg.ACMarketplaceApiConfig.ApiKey)

	res, err := a.httpClient.Do(req)
	if err != nil {
		return retVal, errors.Wrap(err)
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNotFound {
		return retVal, errors.New(fmt.Sprintf("acmarketplace -> GetCustomer -> unexpected status code %d", res.StatusCode))
	}

	err = json.NewDecoder(res.Body).Decode(&retVal)
	if err != nil {
		return retVal, errors.Wrap(err)
	}

	return retVal, nil
}
