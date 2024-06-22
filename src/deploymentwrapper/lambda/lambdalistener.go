// Package lambdalistener is a wrapper around AWS lambda to unify it with regular HTTP server
package lambdalistener

import (
	"context"
	"deploymentwrapper"
	"encoding/json"
	"fmt"
	"sharedconfig"

	"github.com/aws/aws-lambda-go/lambda"

	"handler"
)

// New creates the new listener
func New() deploymentwrapper.Listener {
	return &listen{
		handler: handler.New(false),
	}
}

type pathRequest struct {
	Path string `json:"path"`
}

type listen struct {
	handler handler.StructHandler
	config  sharedconfig.SharedConfiger
}

type lambdaHandlerFunc = func(context.Context, json.RawMessage) (string, error)

func (l *listen) Start(handlers ...deploymentwrapper.HandlerDef) error {
	lambda.Start(l.middleware(handlers...))

	return nil
}

func (l *listen) Port(_ int) {
	// this function is not requred for Lambda
}

func (l *listen) Config(config sharedconfig.SharedConfiger) {
	l.config = config
}

func (l *listen) middleware(handlers ...deploymentwrapper.HandlerDef) lambdaHandlerFunc {
	return func(_ context.Context, rawEvent json.RawMessage) (string, error) {
		rawMessage := string(rawEvent)
		var path pathRequest
		err := json.Unmarshal([]byte(rawMessage), &path)
		if err != nil {
			return "Error unmarshal to path request", err
		}

		requestPath := "/"
		if path.Path != "" {
			requestPath = path.Path
		}

		for _, fcHandler := range handlers {
			if fcHandler.Route == requestPath {
				res, err := l.handler.Process(l.config, fcHandler.Handler, rawMessage)
				if err != nil {
					return "hanler returned error:", err
				}
				return res, nil
			}
		}
		return "route not found", fmt.Errorf("route not found")
	}
}
