package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")

	echoLambda *echoadapter.EchoLambda
)

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("echo cold start")
	e := echo.New()
	e.GET("/nft", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "nft"})
	})
	e.GET("/nft/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "nft hello"})
	})
	echoLambda = echoadapter.New(e)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
