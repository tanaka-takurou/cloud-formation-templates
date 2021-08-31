package main

import (
	"log"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	log.Fatal("Error")
	return Response {
		StatusCode: 200,
		Body: "success",
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
