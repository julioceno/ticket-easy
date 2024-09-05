#!/usr/bin/env bash

set -euo pipefail

# enable debug
# set -x

echo "configuring sqs"
echo "==================="
LOCALSTACK_HOST=localhost
AWS_REGION=eu-central-1

create_queue() {
    local QUEUE_NAME_TO_CREATE=$1
    awslocal sqs create-queue --queue-name ${QUEUE_NAME_TO_CREATE}
}

create_lambda() {
    local LAMBDA_NAME_TO_CREATE=$1
    local LAMBDA_NAME_FILE=$2
    
    awslocal lambda create-function \
        --function-name ${LAMBDA_NAME_TO_CREATE} \
        --runtime nodejs16.x \
        --zip-file ${LAMBDA_NAME_FILE} \
        --handler function.handler \
        --role arn:aws:iam::000000000000:role/cool-stacklifter
}

create_queue "reduce-ticket"

create_lambda "payment"


echo "finished"
