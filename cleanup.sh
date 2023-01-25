#!/bin/bash

aws lambda list-functions --region us-east-2 | jq -r '.Functions | .[] | .FunctionName' |
while read uname1; do
echo "Deleting $uname1";
aws lambda delete-function --region us-east-2 --function-name $uname1;
done
