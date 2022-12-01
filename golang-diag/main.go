package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var sanitzeHeaders = []string{
	"authorization",
}

type DiagnosticResponse struct {
	EventType string      `json:"eventType"`
	Event     interface{} `json:"event"`
	Message   string      `json:"message"`
}

func handleLambdaEvent(event interface{}) (DiagnosticResponse, error) {
	response := DiagnosticResponse{}

	switch evt := discoverPayloadType(event).(type) {
	case events.APIGatewayProxyRequest:
		sanitze(evt.Headers)
		response.Event = evt
		response.EventType = fmt.Sprintf("%T", evt)
		response.Message = "looks like an apigw proxy request"
	case events.ALBTargetGroupRequest:
		sanitze(evt.Headers)
		response.Event = evt
		response.EventType = fmt.Sprintf("%T", evt)
		response.Message = "looks like an alb target group request"
	case map[string]interface{}:
		sanitze(evt["headers"])
		response.Event = evt
		response.EventType = fmt.Sprintf("%T", evt)
		response.Message = "looks like a generic map payload"
	case map[interface{}]interface{}:
		sanitze(evt["headers"])
		response.Event = evt
		response.EventType = fmt.Sprintf("%T", evt)
		response.Message = "looks like a generic map payload"
	}
	return response, nil
}

func discoverPayloadType(event interface{}) interface{} {
	// order matters - most generic last - if two types are ambiguous, add additional
	// validation logic as needed
	apigwEvt, err := couldBe[events.APIGatewayProxyRequest](event)
	if err == nil {
		if apigwEvt.RequestContext.RequestID != "" {
			return apigwEvt
		}
	} else {
		fmt.Println(err.Error())
	}

	albEvt, err := couldBe[events.ALBTargetGroupRequest](event)
	if err == nil {
		if albEvt.RequestContext.ELB.TargetGroupArn != "" {
			return albEvt
		}
	} else {
		fmt.Println(err.Error())
	}

	genericPayload, err := couldBe[map[string]interface{}](event)
	if err == nil {
		return genericPayload
	} else {
		fmt.Println(err.Error())
	}

	return nil
}

func couldBe[e any](event interface{}) (e, error) {
	var result e
	err := JSONRemarshalStrict(event, &result)
	return result, err
}

func sanitze(values interface{}) {
	if values == nil {
		return
	}

	switch vals := values.(type) {
	case map[string]interface{}:
		for k, v := range vals {
			for _, h := range sanitzeHeaders {
				if strings.EqualFold(k, h) {
					if _, ok := v.(string); ok {
						vals[k] = "****"
					} else if _, ok := v.([]string); ok {
						vals[k] = []string{"****"}
					}
				}
			}
		}
	case map[string]string:
		for k, _ := range vals {
			for _, h := range sanitzeHeaders {
				if strings.EqualFold(k, h) {
					vals[k] = "****"
				}
			}
		}
	case interface{}:
	}
}

func main() {
	if os.Getenv("AWS_LAMBDA_RUNTIME_API") != "" {
		lambda.Start(handleLambdaEvent)
		return
	}

	if os.Getenv("LAMBDA_TEST_PAYLOAD_FILE") != "" {
		data, err := ioutil.ReadFile(os.Getenv("LAMBDA_TEST_PAYLOAD_FILE"))
		if err != nil {
			fmt.Printf("failed to read %s: %s", "", err)
			os.Exit(1)
		}
		payload := map[string]interface{}{}
		err = json.Unmarshal(data, &payload)
		if err != nil {
			fmt.Printf("failed to unmarshal %s: %s", "", err)
			os.Exit(1)
		}

		result, err := handleLambdaEvent(payload)
		if err != nil {
			fmt.Printf("failed to handle request: %s", err)
			os.Exit(1)
		}

		resultString, _ := json.Marshal(&result)
		fmt.Println(string(resultString))
		return
	}

	fmt.Println("not lambda runtime (AWS_LAMBDA_RUNTIME_API not set) and LAMBDA_TEST_PAYLOAD_FILE not set - exiting without doing anything!")
}

func JSONRemarshal(obj interface{}, out interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, out)
}

func JSONRemarshalStrict(obj interface{}, out interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(bytes.NewBuffer(data))
	// decoder.DisallowUnknownFields()
	return decoder.Decode(out)
}
