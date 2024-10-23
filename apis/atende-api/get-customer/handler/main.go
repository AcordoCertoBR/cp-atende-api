package main

import (
	"context"
	"log/slog"

	"github.com/AcordoCertoBR/cp-atende-api/apis/atende-api/get-customer/service"
	"github.com/AcordoCertoBR/cp-atende-api/libs/auth"
	"github.com/AcordoCertoBR/cp-atende-api/libs/config"
	httpUtils "github.com/AcordoCertoBR/cp-atende-api/libs/http"
	"github.com/AcordoCertoBR/cp-atende-api/libs/logger"
	"github.com/AcordoCertoBR/streamsurfer"

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
	queue, err := streamsurfer.New("cp-atende-audit-stream")
	if err != nil {
		slog.Error("Error creating KinesisQueue", "error", err)
		return
	}

	defer func() {
		_, err := queue.Flush()
		if err != nil {
			slog.Error("Error flushing queue", "error", err)
		}
	}()

	token := req.Headers["Authorization"]
	claims, err := auth.ValidateJWT(token, cfg.Auth0.PublicCertificate)
	if err != nil {
		slog.Error(err.Error())
		return httpUtils.UnauthorizedResponse(), nil
	}

	data := map[string]interface{}{
		"event":    "atende.GetCustomer.v1",
		"ip":       req.RequestContext.Identity.SourceIP,
		"action":   "get-customer",
		"operator": claims.Sub,
	}
	err = queue.Enqueue(data)
	if err != nil {
		slog.Error("Error enqueuing data", "error", err)
		return
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
