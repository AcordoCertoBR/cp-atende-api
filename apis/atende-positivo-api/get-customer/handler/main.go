package main

import (
	"context"

	"github.com/AcordoCertoBR/ac-atende-positivo-api/libs/config"
	httpUtils "github.com/AcordoCertoBR/ac-atende-positivo-api/libs/http"

	"github.com/AcordoCertoBR/ac-atende-positivo-api/apis/atende-positivo-api/get-customer/service"
	"github.com/AcordoCertoBR/ac-atende-positivo-api/libs/acmarketplace"
	"github.com/AcordoCertoBR/ac-atende-positivo-api/libs/errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var ACMarketplace *acmarketplace.ACMarketplace
var cfg *config.Config
var httpClient *httpUtils.Http

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	document := req.PathParameters["document"]
	if document == "" {
		return httpUtils.ValidationError("document is required"), nil
	}

	customerSvc := service.NewGetCustomerService(ACMarketplace)

	response, err := customerSvc.GetCustomer(document)
	if err != nil {
		return res, errors.Wrap(err)
	}

	if response.Data.User.Documento == "" {
		return httpUtils.NotFoundResponse(), nil
	}

	return httpUtils.SuccessResponse(response), nil
}

func initDependencies() (err error) {
	cfg = config.NewConfig()

	httpClient = httpUtils.New("cp-atende-api", 10)

	ACMarketplace = acmarketplace.NewACMarkeplace(httpClient, cfg)

	return nil
}

func main() {
	err := initDependencies()
	if err != nil {
		panic(err.Error())
	}

	lambda.Start(Handler)
}
