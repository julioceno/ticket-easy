package awsServices

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
)

type Envs struct {
	awsUrl          string
	awsRegion       string
	awsAccessKeyID  string
	awsSecretKey    string
	awsSessionToken string

	queueName string
	lambdaArn string
}

var (
	awsSession *session.Session
)

func Initialize() {
	envs := getEnvs()
	awsSession = createSession(&envs)

	initializeQueue(&envs)
	initializeEventBridge(&envs)
}

func getEnvs() Envs {
	queueName := os.Getenv("QUEUE_REDUCE_TICKET_NAME")
	throwErrorIfEnvNotExists("QUEUE_REDUCE_TICKET_NAME", queueName)

	// TODO: alterar o nome da variavel de ambiente
	lambdaArn := os.Getenv("LAMBDA_ARN")
	throwErrorIfEnvNotExists("LAMBDA_ARN", lambdaArn)

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
		lambdaArn:       lambdaArn,
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
