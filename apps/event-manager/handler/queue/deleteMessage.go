package queue

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

func DeleteMessage(receiptHandle *string) error {
	svc := sqs.New(awsSession)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueRollbackTicket.QueueUrl,
		ReceiptHandle: receiptHandle,
	})

	return err
}
