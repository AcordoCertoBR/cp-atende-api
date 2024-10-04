package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
)

func headers() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Credentials": "true",
		"Content-Type":                     "application/json",
	}
}

func SuccessResponse(value any) events.APIGatewayProxyResponse {
	body, err := generateBody(value)
	if err != nil {
		return ErrorResponse(err, http.StatusInternalServerError)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    headers(),
		Body:       body,
	}
}

func Response(statusCode int, value any) events.APIGatewayProxyResponse {
	body, err := generateBody(value)
	if err != nil {
		return ErrorResponse(err, http.StatusInternalServerError)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headers(),
		Body:       body,
	}
}

func CreatedResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers:    headers(),
	}
}

func NoContentResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
		Headers:    headers(),
	}
}

func NotFoundResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Headers:    headers(),
	}
}

func ForbiddenResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusForbidden,
		Headers:    headers(),
	}
}

func UnauthorizedResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusUnauthorized,
		Headers:    headers(),
	}
}

func ErrorResponse(err error, status int) events.APIGatewayProxyResponse {
	body := ""
	if err != nil {
		body = err.Error()
	}
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers:    headers(),
		Body:       body,
	}
}

func JSONErrorResponse(err error, status int) events.APIGatewayProxyResponse {
	body, err := generateBody(map[string]string{
		"message": err.Error(),
	})
	if err != nil {
		return ErrorResponse(err, http.StatusInternalServerError)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers:    headers(),
		Body:       body,
	}
}

func ValidationError(message string) events.APIGatewayProxyResponse {
	return JSONErrorResponse(errors.New(message), http.StatusBadRequest)
}

func InternalServerErrorResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers:    headers(),
	}
}

func generateBody(value any) (string, error) {
	body := ""
	if value == nil {
		return body, nil
	}

	if reflect.TypeOf(value).String() != "string" {
		byteOutput, err := json.Marshal(value)
		if err != nil {
			return "", errors.New("coudnt generate responde body")
		}

		var buf bytes.Buffer
		json.HTMLEscape(&buf, byteOutput)

		body = buf.String()
	} else {
		body = value.(string)
	}

	return body, nil
}
