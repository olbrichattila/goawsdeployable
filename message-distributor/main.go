package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Replace these with your actual SNS topic ARN and SQS queue URL
const (
	snsTopicArn = "arn:aws:sns:us-east-1:000000000000:my-topic"
	sqsQueueUrl = "http://localstack:4566/000000000000/test"
)

func main() {

	cfg := aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localstack:4566"),
		Credentials: credentials.NewStaticCredentials(
			"your-access-key-id",     // Replace with your actual AWS Access Key ID
			"your-secret-access-key", // Replace with your actual AWS Secret Access Key
			"",
		),
	}

	sess, err := session.NewSession(&cfg)
	if err != nil {
		log.Printf("error creating session: %vzn", err)
	}

	// Create SQS and SNS clients
	sqsClient := sqs.New(sess)
	snsClient := sns.New(sess)

	// Receive messages from SQS
	for {
		receiveMessageOutput, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(sqsQueueUrl),
			MaxNumberOfMessages: aws.Int64(10), // Adjust as needed
			WaitTimeSeconds:     aws.Int64(10), // Long polling
		})
		if err != nil {
			log.Printf("Error receiving message from SQS: %v\n", err)
		}

		for _, message := range receiveMessageOutput.Messages {
			// Publish each SQS message to the SNS topic
			_, err := snsClient.Publish(&sns.PublishInput{
				Message:  message.Body,
				TopicArn: aws.String(snsTopicArn),
			})
			if err != nil {
				log.Printf("error publishing message to SNS: %v\n", err)
			}

			fmt.Printf("Message %s sent to SNS topic\n", *message.MessageId)

			// Delete the message from SQS after successfully publishing to SNS
			_, err = sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(sqsQueueUrl),
				ReceiptHandle: message.ReceiptHandle,
			})
			if err != nil {
				log.Printf("error deleting message from SQS: %v\n", err)
			}

			fmt.Printf("Message %s deleted from SQS queue\n", *message.MessageId)
		}
		// time.Sleep(time.Second)
	}
}
