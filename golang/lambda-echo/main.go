package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type DiagnosticResponse struct {
	Message string `json:"message"`
}

func HandleLambdaEvent(event interface{}) (DiagnosticResponse, error) {
	return DiagnosticResponse{Message: "todo"}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
