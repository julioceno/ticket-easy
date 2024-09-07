package awsServices

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
)

var (
	queueUrl *sqs.GetQueueUrlOutput
)

func initializeQueue(queueName *string) *sqs.GetQueueUrlOutput {
	queueService := sqs.New(awsSession)
	result, err := queueService.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queueName,
	})
	if err != nil {
		logger.Fatal("Ocurred error when get queue connection", err)
	}

	logger.Info("Obtained queue connection")
	return result
}

func SendMessage(queueUrl *sqs.GetQueueUrlOutput, message string) error {
	svc := sqs.New(awsSession)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(message),
		QueueUrl:     queueUrl.QueueUrl,
	})

	if err != nil {
		return err
	}

	return nil
}

func ReceveidMessage(queueUrl *sqs.GetQueueUrlOutput) (*sqs.ReceiveMessageOutput, error) {
	svc := sqs.New(awsSession)

	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            queueUrl.QueueUrl,
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(10),
	})

	if err != nil {
		return nil, err
	}

	return msgResult, nil
}

func DeleteMessage(queueUrl *sqs.GetQueueUrlOutput, receiptHandle *string) error {
	svc := sqs.New(awsSession)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueUrl.QueueUrl,
		ReceiptHandle: receiptHandle,
	})

	return err
}
