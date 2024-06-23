// Package snsMiddleware routes message from SNS queue to a string request, distingusing between HTTP and Lambda topic subscription
package snsmiddleware

import (
	"context"
	"fmt"
	"handler"
	"net/http"
	"sharedconfig"
)

// type SnsHandler func(*context.Context, sharedconfig.SharedConfiger, string) (string, error)

// TODO many of those values are not yet used
type request struct {
	requestItem
	Records []sNSReord `json:"Records"`
}

type requestItem struct {
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

type sNSReord struct {
	EventSource          string `json:"EventSource"`
	EventVersion         string `json:"EventVersion"`
	EventSubscriptionArn string `json:"EventSubscriptionArn"`
	Sns                  requestItem
}

type response struct {
	Response   []string `json:"respopnse"`
	ConfigType string   `json:"configType"`
}

func Middleware(handlerFunc any) handler.StructHandlerFunc {
	haldlerExecutor := handler.New(true)

	return func(ctx *context.Context, config sharedconfig.SharedConfiger, request *request) (*response, error) {
		var responses []string
		if request.Type == "SubscriptionConfirmation" {
			endpoint := config.GetSQSConfig().AWSConfig.Endpoint
			subscribeURL := fmt.Sprintf("%s?Action=ConfirmSubscription&TopicArn=%s&Token=%s", *endpoint, request.TopicArn, request.Token)
			_, err := http.Get(subscribeURL)
			if err != nil {
				return nil, err
			}
			return &response{Response: responses, ConfigType: config.GetConfigType()}, nil
		}

		if config.GetConfigType() == sharedconfig.TypeHttp {
			response, err := haldlerExecutor.Process(config, handlerFunc, asResponseItem(request).Message)
			// response, err := handlerFunc(ctx, config, asResponseItem(request).Message)
			if err != nil {
				return nil, err
			}

			responses = append(responses, response)
		}

		if config.GetConfigType() == sharedconfig.TypeLambda {
			for _, record := range request.Records {
				response, err := haldlerExecutor.Process(config, handlerFunc, record.Sns.Message)
				// response, err := handlerFunc(ctx, config, record.Sns.Message)
				if err != nil {
					return nil, err
				}

				responses = append(responses, response)
			}
		}

		return &response{Response: responses, ConfigType: config.GetConfigType()}, nil
	}
}

func asResponseItem(req *request) *requestItem {
	return &requestItem{
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
