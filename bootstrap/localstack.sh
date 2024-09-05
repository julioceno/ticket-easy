#!/usr/bin/env bash

set -euo pipefail

# enable debug
# set -x

echo "configuring services"
echo "==================="

create_queue() {
    local QUEUE_NAME_TO_CREATE=$1
    awslocal sqs create-queue --queue-name ${QUEUE_NAME_TO_CREATE}
}

create_lambda() {
    local LAMBDA_NAME_TO_CREATE=$1
    local LAMBDA_DIR=$2

    cd ${LAMBDA_DIR}
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${LAMBDA_NAME_TO_CREATE} main.go
    ~/Go/Bin/build-lambda-zip.exe -o ${LAMBDA_NAME_TO_CREATE}.zip ${LAMBDA_NAME_TO_CREATE}
    cd ..

    local LAMBDA_ZIP_PATH="${LAMBDA_DIR}/${LAMBDA_NAME_TO_CREATE}.zip"

    echo ${LAMBDA_ZIP_PATH}

    awslocal lambda create-function \
        --function-name ${LAMBDA_NAME_TO_CREATE} \
        --runtime go1.x \
        --zip-file fileb://${LAMBDA_ZIP_PATH} \
        --handler ${LAMBDA_NAME_TO_CREATE} \
        --role arn:aws:iam::000000000000:role/cool-stacklifter
}


create_queue "reduce-ticket"
echo "Queue created"

create_lambda "verify-payment" "./verify-payment" 
echo "lambda created"

echo "finished"
