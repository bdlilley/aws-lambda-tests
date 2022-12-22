#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
NAME="$(basename $SCRIPT_DIR)-latest"

GOOS=linux GOARCH=amd64 go build -o main

zip lambda.zip main

aws s3 cp lambda.zip s3://solo-io-terraform-931713665590/lambda/${NAME}.zip

aws lambda create-function \
    --function-name ${NAME} \
    --runtime go1.x \
    --code=S3Bucket=solo-io-terraform-931713665590,S3Key=lambda/${NAME}.zip \
    --handler main \
    --role arn:aws:iam::931713665590:role/lambda-basic \
    || true

aws lambda update-function-code \
    --function-name ${NAME} \
    --s3-bucket solo-io-terraform-931713665590 \
    --s3-key lambda/${NAME}.zip

# kubectl apply -f - <<EOF
# apiVersion: lambda.aws.crossplane.io/v1beta1
# kind: Function
# metadata:
#   name: ${NAME}
#   namespace: crossplane-system
# spec:
#   forProvider:
#     region: us-east-2
#     description: ${NAME}
#     runtime: go1.x
#     code:
#       s3Bucket: solo-io-terraform-931713665590
#       s3Key: lambda/${NAME}.zip
# EOF