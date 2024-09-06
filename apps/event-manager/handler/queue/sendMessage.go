package queue

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SendMessage(message string) error {
	svc := sqs.New(awsSession)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(message),
		QueueUrl:     queueReduceTicket.QueueUrl,
	})

	if err != nil {
		return err
	}

	return nil
}
