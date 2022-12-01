package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

type GenericEvent struct {
	Headers map[string]interface{} `json:"headers"`
	Path    string                 `json:"path"`
}

func main() {
	lambda.Start(func(event GenericEvent) (interface{}, error) {
		spew.Dump(event)
		return map[string]string{
			"message": "hello, world",
		}, nil
	})
}
