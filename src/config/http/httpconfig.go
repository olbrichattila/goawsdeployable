// Package httpconfig will be built if you are building for http
// Please make sure you are using the same structure for both
package httpconfig

import (
	"sharedconfig"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

// New is a function to create a new shared config.
// This is a config, with shared valeus, and lambda config specific values
// The interface ensures you will use the same strucure for both enfironment
// First implement the getter into the shared config interface above. Add getters only!
func New() sharedconfig.SharedConfiger {
	return &config{}
}

type config struct {
	// todo your config here
}

// GetSQSConfig implements sharedconfig.SharedConfiger.
func (c *config) GetSQSConfig() *sharedconfig.SQSConfig {
	return &sharedconfig.SQSConfig{
		QueueURL: "http://localhost:4566/000000000000/test",
		AWSConfig: aws.Config{
			Region:   aws.String("us-east-1"),
			Endpoint: aws.String("http://localhost:4566"),
			Credentials: credentials.NewStaticCredentials(
				"your-access-key-id",     // Replace with your actual AWS Access Key ID
				"your-secret-access-key", // Replace with your actual AWS Secret Access Key
				"",
			),
		},
	}
}

func (c *config) GetConfigType() string {
	return sharedconfig.TypeHttp
}
