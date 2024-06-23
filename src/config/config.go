// Package sharedconfig is an interface to be used across both config files. The config have to use getters
// to ensure they follow the same structure
package sharedconfig

import "github.com/aws/aws-sdk-go/aws"

const (
	TypeLambda = "lambda"
	TypeHttp   = "http"
)

// SharedConfiger is the interface to define the config values across http/lamda env
type SharedConfiger interface {
	GetConfigType() string
	GetSQSConfig() *SQSConfig
	GetDBConfig() *DBConfig
}

type SQSConfig struct {
	QueueURL  string
	AWSConfig aws.Config
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}
