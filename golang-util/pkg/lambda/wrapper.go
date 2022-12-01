package lambda

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func WrappedH(lambda func(handler interface{}), handler func(event interface{})) (interface{}, error) {
	if os.Getenv("AWS_LAMBDA_RUNTIME_API") != "" {
		lambda(handler)
		return nil, nil
	}

	if os.Getenv("LAMBDA_TEST_PAYLOAD_B64") != "" {
		data, err := ioutil.ReadFile(os.Getenv("LAMBDA_TEST_PAYLOAD_B64"))
		if err != nil {
			return nil, fmt.Errorf("failed to read %s: %s", "", err)
		}
		payload := map[string]interface{}{}
		err = json.Unmarshal(data, &payload)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal %s: %s", "", err)
		}
		handler(payload)
	}

	return nil, nil
}
