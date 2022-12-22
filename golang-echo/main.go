package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	lambda.Start(func(event interface{}) (interface{}, error) {
		spew.Dump(event)
		return event, nil
	})
}
