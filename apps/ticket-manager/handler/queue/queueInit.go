package queue

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
)

type Envs struct {
	queueName       string
	awsUrl          string
	awsRegion       string
	awsAccessKeyID  string
	awsSecretKey    string
	awsSessionToken string
}

var (
	awsSession *session.Session
	queueUrl   *sqs.GetQueueUrlOutput
)

func init() {
	envs := getEnvs()

	awsSession = createSession(&envs)
	queueUrl = getQueueURL(&envs)
}

func getEnvs() Envs {
	queueName := os.Getenv("QUEUE_REDUCE_TICKET_NAME")
	throwErrorIfEnvNotExists("QUEUE_REDUCE_TICKET_NAME", queueName)

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
		queueName:       queueName,
		awsUrl:          awsUrl,
		awsRegion:       awsRegion,
		awsAccessKeyID:  awsAccessKeyID,
		awsSecretKey:    awsSecretKey,
		awsSessionToken: awsSessionToken,
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

func getQueueURL(envs *Envs) *sqs.GetQueueUrlOutput {
	queueName := envs.queueName
	svc := sqs.New(awsSession)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		logger.Fatal("Ocurred error when get queue connection", err)
	}

	logger.Info("Obtained queue connection")

	return result
}
