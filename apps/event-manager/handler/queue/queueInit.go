package queue

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/julioceno/ticket-easy/apps/event-manager/config/logger"
)

type Envs struct {
	queueDecreaseTicket string
	queueRollbackTicket string

	awsUrl          string
	awsRegion       string
	awsAccessKeyID  string
	awsSecretKey    string
	awsSessionToken string
}

var (
	awsSession          *session.Session
	QueueDecreaseTicket *sqs.GetQueueUrlOutput
	QueueRollbackTicket *sqs.GetQueueUrlOutput
)

func init() {
	envs := getEnvs()

	awsSession = createSession(&envs)
	QueueDecreaseTicket = getQueueURL(&envs.queueDecreaseTicket)
	QueueRollbackTicket = getQueueURL(&envs.queueRollbackTicket)
}

func getEnvs() Envs {
	queueDecreaseTicket := os.Getenv("QUEUE_Decrease_TICKET_NAME")
	throwErrorIfEnvNotExists("QUEUE_DECREASE_TICKET_NAME", queueDecreaseTicket)

	queueRollbackTicket := os.Getenv("QUEUE_ROLLBACK_TICKET_NAME")
	throwErrorIfEnvNotExists("QUEUE_ROLLBACK_TICKET_NAME", queueRollbackTicket)

	awsUrl := os.Getenv("AWS_URL")
	throwErrorIfEnvNotExists("AWS_URL", awsUrl)

	awsRegion := os.Getenv("AWS_REGION")
	throwErrorIfEnvNotExists("AWS_REGION", awsRegion)

	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	throwErrorIfEnvNotExists("AWS_ACCESS_KEY_ID", awsAccessKeyID)

	awsSecretKey := os.Getenv("AWS_SECRET_KEY")
	throwErrorIfEnvNotExists("AWS_SECRET_KEY", awsSecretKey)

	awsSessionToken := os.Getenv("AWS_SESSION_TOKEN")
	throwErrorIfEnvNotExists("AWS_SESSION_TOKEN", awsSessionToken)

	return Envs{
		queueDecreaseTicket: queueDecreaseTicket,
		queueRollbackTicket: queueRollbackTicket,
		awsUrl:              awsUrl,
		awsRegion:           awsRegion,
		awsAccessKeyID:      awsAccessKeyID,
		awsSecretKey:        awsSecretKey,
		awsSessionToken:     awsSessionToken,
	}
}

func throwErrorIfEnvNotExists(key string, value string) {
	if value == "" {
		logger.Fatal(fmt.Sprintf("%s não existe", key), nil)
	}
}

func createSession(envs *Envs) *session.Session {
	logger.Info("Creating aws session")
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:    aws.String(envs.awsUrl),
		Region:      aws.String(envs.awsRegion),
		Credentials: credentials.NewStaticCredentials(envs.awsAccessKeyID, envs.awsSecretKey, envs.awsSessionToken),
	}))
	logger.Info("Session created, create queue instance")

	return sess
}

func getQueueURL(queueUrl *string) *sqs.GetQueueUrlOutput {
	svc := sqs.New(awsSession)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queueUrl,
	})
	if err != nil {
		logger.Fatal("Ocurred error when get queue connection", err)
	}

	logger.Info("Obtained queue connection")

	return result
}
