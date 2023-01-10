package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

type GlooRequest struct {
	Headers     map[string]string `json:"headers"`
	Path        string            `json:"path"`
	HttpMethod  string            `json:"httpMethod"`
	QueryString string            `json:"queryString"`
	Body        string            `json:"body"`
}

type GlooResponse struct {
	Headers    map[string]string `json:"headers"`
	StatusCode string            `json:"statusCode"`
	Body       interface{}       `json:"body"`
}

// AWS_ACCOUNT_ID="931713665590" AWS_REGION=us-east-2 AWS_PROFILE=product ./deploy.sh
// AWS_ACCOUNT_ID="410461945957" AWS_REGION=us-west-2 AWS_PROFILE=default ./deploy.sh
// istioctl proxy-config all istio-ingressgateway-848d7fbb8-jq5bk -n istio-gateway-ns -o json
//
// kubectl port-forward -n istio-gateway-ns istio-ingressgateway-d74c76c66-drzv4 15000
// localhost:15000/config_dump
func main() {
	lambda.Start(func(event GlooRequest) (interface{}, error) {
		spew.Dump(event)
		headersOut := make(map[string]string)
		for k, v := range event.Headers {
			if !strings.HasPrefix(k, ":") {
				headersOut[k] = v
			}
		}

		// xid is populated with the correct request id
		// xid := event.Headers["x-request-id"]

		// if I add it to the response with the same value as the request
		// it does not appear in gloo response:
		// headersOut["x-request-id"] = xid

		sBody, _ := json.Marshal(event)

		// but if I set it to any other value it does appear
		headersOut["x-request-id"] = "d5a16d04"

		genericPayload := map[string]interface{}{
			"statusCode": 202,
			"body":       fmt.Sprintf(`{"hello": "world", "originalEvent": %s}`, string(sBody)),
			"headers":    headersOut,
		}

		return genericPayload, nil
		// return events.APIGatewayProxyResponse{
		// 	StatusCode: 202,
		// 	Body:       fmt.Sprintf(`{"hello": "world", "originalEvent": %s}`, string(sBody)),
		// 	Headers:    headersOut,
		// 	MultiValueHeaders: map[string][]string{
		// 		"foo": {"baz", "bar"},
		// 	},
		// 	IsBase64Encoded: false,
		// }, nil
	})
}
