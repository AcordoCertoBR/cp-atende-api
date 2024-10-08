package main

import (
	"context"
	"log/slog"

	"github.com/AcordoCertoBR/cp-atende-api/apis/atende-api/get-customer/service"
	"github.com/AcordoCertoBR/cp-atende-api/libs/auth"
	"github.com/AcordoCertoBR/cp-atende-api/libs/config"
	httpUtils "github.com/AcordoCertoBR/cp-atende-api/libs/http"
	"github.com/AcordoCertoBR/cp-atende-api/libs/logger"
	"github.com/golang-jwt/jwt"

	"github.com/AcordoCertoBR/cp-atende-api/libs/acmarketplace"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var ACMarketplace *acmarketplace.ACMarketplace
var cfg *config.Config
var httpClient *httpUtils.Http

/*
TODO:
- Add ip whitelist
- Integrate with auth0
- Integrate datadog
- Integrate redline
*/
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	logger.SetupLogger(cfg)

	token := req.Headers["Authorization"]
	valid, err := auth.ValidateJWT(cfg.Auth0.JwtSecret, token, jwt.SigningMethodHS256.Name)
	if err != nil {
		slog.Error(err.Error())
		return httpUtils.UnauthorizedResponse(), nil
	}

	if !valid {
		slog.Error("Invalid token")
		return httpUtils.UnauthorizedResponse(), nil
	}

	document := req.PathParameters["document"]
	if document == "" {
		slog.Error("document is required")
		return httpUtils.ValidationError("document is required"), nil
	}

	customerSvc := service.NewGetCustomerService(ACMarketplace)

	response, err := customerSvc.GetCustomer(document)
	if err != nil {
		slog.Error(err.Error())
		return httpUtils.InternalServerErrorResponse(), nil
	}

	slog.Info("document", "document", response.Data.User.Documento)

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
