package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

type DiagnosticResponse struct {
	Message string `json:"message"`
}

func HandleLambdaEvent(event interface{}) (DiagnosticResponse, error) {
	spew.Dump(event)
	return DiagnosticResponse{Message: "todo"}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
