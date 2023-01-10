# golang-apigw

Simple go lambda to demonstrate how Gloo can replace AWS API Gateway without lambda code changes.

If the request succeeds, the response will be formatted in the AWS API Gateway Proxy integration format, and the body will be a JSON payload that contains the current time.