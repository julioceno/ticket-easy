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

create_queue "reduce-ticket"

echo "finished"
