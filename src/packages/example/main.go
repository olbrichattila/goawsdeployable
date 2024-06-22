// Package example is just an example how to create a module accross AWS lambda or HTTP with the same code
package example

// TODO: refactor this, and make it an adapter between SNS and HTML to be usable generally
import (
	"context"
	"fmt"
	"net/http"
	"sharedconfig"
	"sqseventdispatcher"
)

// Request data lambda sends full message(s)
type Request struct {
	RequestItem
	Records []SnsReord `json:"Records"`
}

// One request item, only this is sent if SNS sub is HTTP
type RequestItem struct {
	Name             string `json:"name"`
	Type             string `json:"Type"`
	MessageId        string `json:"MessageId"`
	TopicArn         string `json:"TopicArn"`
	Token            string `json:"Token"`
	Subject          string `json:"Subject"`
	Message          string `json:"Message"`
	Timestamp        string `json:"Timestamp"`
	SignatureVersion string `json:"SignatureVersion"`
	Signature        string `json:"Signature"`
	SigningCertURL   string `json:"SigningCertURL"`
	SubscribeURL     string `json:"SubscribeURL"`
	UnsubscribeURL   string `json:"UnsubscribeURL"`
}

// Onse SNS message record
type SnsReord struct {
	EventSource          string `json:"EventSource"`
	EventVersion         string `json:"EventVersion"`
	EventSubscriptionArn string `json:"EventSubscriptionArn"`
	Sns                  RequestItem
}

// Response what we want to be returned as HTTP or Lambda
type Response struct {
	ResponseName string `json:"respopnseName"`
	ConfigType   string `json:"configType"`
}

// TestHandler is the unfied entry point of the module
func TestHandler(_ *context.Context, config sharedconfig.SharedConfiger, request *Request) (*Response, error) {
	if request.Type == "SubscriptionConfirmation" {
		endpoint := config.GetSQSConfig().AWSConfig.Endpoint
		subscribeURL := fmt.Sprintf("%s?Action=ConfirmSubscription&TopicArn=%s&Token=%s", *endpoint, request.TopicArn, request.Token)
		_, err := http.Get(subscribeURL)
		if err != nil {
			return nil, err
		}
	}

	if config.GetConfigType() == sharedconfig.TypeHttp {
		fmt.Println(request)
		err := process(config, asResponseItem(request))
		if err != nil {
			return nil, err
		}
	}

	if config.GetConfigType() == sharedconfig.TypeLambda {
		for _, record := range request.Records {
			err := process(config, &record.Sns)
			if err != nil {
				return nil, err
			}
		}
	}

	return &Response{
		ResponseName: request.Name,
		ConfigType:   config.GetConfigType(),
	}, nil
}

func asResponseItem(req *Request) *RequestItem {

	return &RequestItem{
		Name:             req.Name,
		Type:             req.Type,
		MessageId:        req.MessageId,
		TopicArn:         req.TopicArn,
		Token:            req.Token,
		Subject:          req.Subject,
		Message:          req.Message,
		Timestamp:        req.Timestamp,
		SignatureVersion: req.SignatureVersion,
		Signature:        req.Signature,
		SigningCertURL:   req.SigningCertURL,
		SubscribeURL:     req.SubscribeURL,
		UnsubscribeURL:   req.UnsubscribeURL,
	}
}

func process(config sharedconfig.SharedConfiger, item *RequestItem) error {
	awsConfig := config.GetSQSConfig()
	return sqseventdispatcher.NewDispatcher(awsConfig).Send(*item)
}
