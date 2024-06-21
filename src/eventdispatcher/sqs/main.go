// Package sqseventdispatcher puths an item from a struct to the AWS SQS queue
package sqseventdispatcher

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Dispatcher event with Send(<your stryct>) error
type Dispatcher interface {
	Send(any) (string, error)
}

type dispatch struct {
}

// NewDispatcher creates a new dispatcher struct
func NewDispatcher() Dispatcher {
	return &dispatch{}
}

func (t dispatch) Send(v any) (string, error) {

	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Struct {
		return "", fmt.Errorf("parameter must be a struct")
	}

	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	messageBody := string(jsonBytes)
	err = t.sendToSqs(messageBody)
	if err != nil {
		return "", err
	}

	return messageBody, nil
}

func (t dispatch) sendToSqs(messageBody string) error {
	// Load the Shared AWS Configuration (~/.aws/config) //TODO: do it a different way
	// TODO get full config from file, not default
	// endpoint := "http://localhost:4566"
	endpoint := "http://localstack:4566"
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   aws.String("us-east-1"),
			Endpoint: &endpoint,
		},
		Profile: "default",
	}))

	// queueURL := "http://localhost:4566/000000000000/test"
	queueURL := "http://localstack:4566/000000000000/test"
	svc := sqs.New(sess)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(0),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("The Whistler"),
			},
			"Author": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("John Grisham"),
			},
			"WeeksOn": &sqs.MessageAttributeValue{
				DataType:    aws.String("Number"),
				StringValue: aws.String("6"),
			},
		},
		MessageBody: aws.String(messageBody),
		QueueUrl:    &queueURL,
	})

	if err != nil {
		return err
	}

	return nil
}
