// Package lambdalistener is a wrapper around AWS lambda to unify it with regular HTTP server
package lambdalistener

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"

	"handler"
)

// New creates the new listener
func New() Listener {
	return &listen{
		handler: handler.New(false),
	}
}

type pathRequest struct {
	Path string `json:"path"`
}

// Listener is the interface to make Lambda and HTTP unified
type Listener interface {
	Start(handlers ...HandlerDef) error
	Port(int)
}

type listen struct {
	handler handler.StructHandler
}

// HandlerDef is the structure how to pass a route and a handler
type HandlerDef struct {
	Route   string
	Handler handler.StructHandlerFunc
}

type lambdaHandlerFunc = func(context.Context, json.RawMessage) (string, error)

func (l *listen) Start(handlers ...HandlerDef) error {
	lambda.Start(l.middleware(handlers...))

	return nil
}

func (l *listen) Port(_ int) {
	// this function is not requred for Lambda
}

func (l *listen) middleware(handlers ...HandlerDef) lambdaHandlerFunc {
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
				res, err := l.handler.Process(fcHandler.Handler, rawMessage)
				if err != nil {
					return "hanler returned error:", err
				}
				return res, nil
			}
		}
		return "route not found", fmt.Errorf("route not found")
	}
}
