package awsServices

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/julioceno/ticket-easy/apps/ticket-manager/config/logger"
)

var (
	eventBridgeService *eventbridge.EventBridge

	LAMBDA_ARN           string
	SCHEDULER_EXPRESSION = "rate(2 minutes)" // TODO: alterar pra 10
)

func initializeEventBridge(envs *Envs) {
	svc := eventbridge.New(awsSession)
	eventBridgeService = svc
	LAMBDA_ARN = envs.lambdaArn
}

func CreateEvent(ticketId string) error {
	// TODO: entender melhor essa parte
	eventBridgeService.PutRule(&eventbridge.PutRuleInput{
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

	_, err = eventBridgeService.PutTargets(&eventbridge.PutTargetsInput{
		Rule: aws.String(ticketId),
		Targets: []*eventbridge.Target{
			{
				Id:    aws.String("lambda"),
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

func DeleteEvent(ticketId string) error {
	_, err := eventBridgeService.DescribeRule(&eventbridge.DescribeRuleInput{
		Name: aws.String(ticketId),
	})
	if err != nil {
		return nil
	}

	ids := []string{"lambda"}
	_, err = eventBridgeService.RemoveTargets(&eventbridge.RemoveTargetsInput{
		Rule:  aws.String(ticketId),
		Ids:   aws.StringSlice(ids),
		Force: aws.Bool(true),
	})
	if err != nil {
		logger.Error("failed remove targets from lambda, %v", err)
		return err
	}

	_, err = eventBridgeService.DeleteRule(&eventbridge.DeleteRuleInput{
		Name: aws.String(ticketId),
	})
	if err != nil {
		logger.Error("failed to delete rule, %v", err)
		return err
	}

	return nil
}
