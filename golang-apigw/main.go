package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

func handleLambdaEvent(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// the event should be an AWS APIGatewayProxyRequest
	// use wrapAsApiGateway in gloo to automatically convert a http request to
	// this format
	spew.Dump(event)

	// to use with real AWS API GW, we just return this format
	// in gloo use unwrapAsApiGateway to convert to a "regular" http response
	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		Body:            `{"message": "hello, apigw"}`,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handleLambdaEvent)
}
