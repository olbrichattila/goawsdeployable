// Package example is just an example how to create a module accross AWS lambda or HTTP with the same code
package example

// TODO: refactor this, and make it an adapter between SNS and HTML to be usable generally
import (
	"context"
	"sharedconfig"
	"sqseventdispatcher"
)

type request struct {
	Name string `json:"name"`
}

type response struct {
	Request request `json:"request"`
}

// TestHandler is the unfied entry point of the module
func TestHandler(_ *context.Context, config sharedconfig.SharedConfiger, request *request) (*response, error) {
	dispatcher := sqseventdispatcher.NewDispatcher(config.GetSQSConfig())
	err := dispatcher.Send(*request)
	if err != nil {
		return nil, err
	}

	return &response{
		Request: *request,
	}, nil
}
