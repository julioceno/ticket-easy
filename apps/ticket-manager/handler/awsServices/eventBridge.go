package awsServices

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
)

var (
	eventBrigeService *eventbridge.EventBridge

	LAMBDA_ARN           string
	SCHEDULER_EXPRESSION = "rate(2 minutes)" // TODO: alterar pra 10
)

func initializeEventBridge(envs *Envs) {
	svc := eventbridge.New(awsSession)
	eventBrigeService = svc
	LAMBDA_ARN = envs.lambdaArn
}

func CreateEvent(ticketId string) error {
	// TODO: entender melhor essa parte
	eventBrigeService.PutRule(&eventbridge.PutRuleInput{
		Name:               aws.String(ticketId),
		ScheduleExpression: aws.String(SCHEDULER_EXPRESSION),
		State:              aws.String(eventbridge.RuleStateEnabled),
	})

	payload := map[string]string{
		"ticketId": ticketId,
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		logger.Error("failed to marshal payload, %v", err)
		return err
	}

	_, err = eventBrigeService.PutTargets(&eventbridge.PutTargetsInput{
		Rule: aws.String(ticketId),
		Targets: []*eventbridge.Target{
			{
				Id:    aws.String("1"),
				Arn:   aws.String(LAMBDA_ARN),
				Input: aws.String(string(payloadJson)),
			},
		},
	})

	if err != nil {
		logger.Error("failed to create target, %v", err)
		return err
	}

	return nil
}
