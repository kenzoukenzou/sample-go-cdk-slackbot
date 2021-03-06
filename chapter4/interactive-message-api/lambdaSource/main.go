package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func main() {
	lambda.Start(interactiveMessageHandler)
}

func interactiveMessageHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.
	APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{}

	signingSecrets := os.Getenv("SIGNING_SECRETS")

	interactiveMessageUsecase := NewInteractionUsecase(signingSecrets)
	response, err := interactiveMessageUsecase.MakeSlackResponse(request)
	if err != nil {
		return response, err
	}
	return response, nil
}
