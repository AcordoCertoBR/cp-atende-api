package main

import (
	"context"
	"log/slog"

	"github.com/AcordoCertoBR/cp-atende-api/apis/atende-api/get-customer/service"
	"github.com/AcordoCertoBR/cp-atende-api/libs/auth"
	"github.com/AcordoCertoBR/cp-atende-api/libs/config"
	httpUtils "github.com/AcordoCertoBR/cp-atende-api/libs/http"
	"github.com/AcordoCertoBR/cp-atende-api/libs/logger"

	"github.com/AcordoCertoBR/cp-atende-api/libs/acmarketplace"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var acMarketplace *acmarketplace.ACMarketplace
var cfg *config.Config
var httpClient *httpUtils.Http

/*
TODO:
- Integrate datadog
- Integrate redline
*/
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	logger.SetupLogger(cfg)

	token := req.Headers["Authorization"]
	_, err = auth.ValidateJWT(token, cfg.Auth0.PublicCertificate)
	if err != nil {
		slog.Error(err.Error())
		return httpUtils.UnauthorizedResponse(), nil
	}

	document := req.PathParameters["document"]
	if document == "" {
		slog.Error("document is required")
		return httpUtils.ValidationError("document is required"), nil
	}

	customerSvc := service.NewGetCustomerService(acMarketplace)

	response, err := customerSvc.GetCustomer(document)
	if err != nil {
		slog.Error(err.Error())
		return httpUtils.InternalServerErrorResponse(), nil
	}

	if response.Data.User.Documento == "" {
		return httpUtils.NotFoundResponse(), nil
	}

	return httpUtils.SuccessResponse(response), nil
}

func initDependencies() (err error) {
	cfg = config.NewConfig()

	httpClient = httpUtils.New("cp-atende-api", 10)

	acMarketplace = acmarketplace.NewACMarkeplace(httpClient, cfg)

	return nil
}

func main() {
	err := initDependencies()
	if err != nil {
		panic(err.Error())
	}

	lambda.Start(Handler)
}
