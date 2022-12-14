package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

type GlooRequest struct {
	Headers     map[string]interface{} `json:"headers"`
	Path        string                 `json:"path"`
	HttpMethod  string                 `json:"httpMethod"`
	QueryString string                 `json:"queryString"`
	Body        string                 `json:"body"`
}

type GlooResponse struct {
	Headers    map[string]interface{} `json:"headers"`
	StatusCode int                    `json:"statusCode"`
	Body       string                 `json:"body"`
}

func main() {
	lambda.Start(func(event GlooRequest) (interface{}, error) {
		spew.Dump(event)
		return GlooResponse{
			StatusCode: 404,
			Body:       "can't find it",
		}, nil
	})
}
