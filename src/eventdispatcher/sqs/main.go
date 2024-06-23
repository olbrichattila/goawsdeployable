// Package sqseventdispatcher puths an item from a struct to the AWS SQS queue
package sqseventdispatcher

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sharedconfig"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Dispatcher event with Send(<your stryct>) error
type Dispatcher interface {
	Send(any) error
	SendString(string) error
}

type dispatch struct {
	awsConfig *sharedconfig.SQSConfig
}

// NewDispatcher creates a new dispatcher struct
func NewDispatcher(config *sharedconfig.SQSConfig) Dispatcher {
	return &dispatch{
		awsConfig: config,
	}
}

func (t *dispatch) Send(v any) error {

	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("parameter must be a struct")
	}

	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	messageBody := string(jsonBytes)
	err = t.SendString(messageBody)
	if err != nil {
		return err
	}

	return nil
}

func (t *dispatch) SendString(messageBody string) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  t.awsConfig.AWSConfig,
		Profile: "default",
	}))

	svc := sqs.New(sess)
	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(0),
		MessageBody:  aws.String(messageBody),
		QueueUrl:     aws.String(t.awsConfig.QueueURL),
	})

	if err != nil {
		return err
	}

	return nil
}
