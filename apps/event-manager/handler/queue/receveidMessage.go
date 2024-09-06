package queue

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ReceveidMessage() (*sqs.ReceiveMessageOutput, error) {
	svc := sqs.New(awsSession)

	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            queueRollbackTicket.QueueUrl,
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(10),
	})

	if err != nil {
		return nil, err
	}

	return msgResult, nil
}
