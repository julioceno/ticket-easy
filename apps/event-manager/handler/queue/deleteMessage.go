package queue

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

func DeleteMessage(queueUrl *sqs.GetQueueUrlOutput, receiptHandle *string) error {
	svc := sqs.New(awsSession)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueUrl.QueueUrl,
		ReceiptHandle: receiptHandle,
	})

	return err
}
