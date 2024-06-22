// Package example is just an example how to create a module accross AWS lambda or HTTP with the same code
package example

import (
	"context"
	"sharedconfig"

	dispather "sqseventdispatcher"
)

// Request data automaticall marhalled here
type Request struct {
	Name string `json:"name"`
}

// Response what we want to be returned as HTTP or Lambda
type Response struct {
	DispactherResult string `json:"dispatherResult"`
	ResponseName     string `json:"respopnseName"`
	ConfigType       string `json:"configType"`
}

// TestHandler is the unfied entry point of the module
func TestHandler(_ *context.Context, config sharedconfig.SharedConfiger, request *Request) (*Response, error) {
	dispatcherResult, err := dispather.NewDispatcher().Send(*request)
	if err != nil {
		return nil, err
	}

	return &Response{
		ResponseName:     request.Name,
		DispactherResult: dispatcherResult,
		ConfigType:       config.GetConfigType(),
	}, nil
}
