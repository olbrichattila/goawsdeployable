// Package example2 is just an example how to create a module accross AWS lambda or HTTP with the same code
package example2

import (
	"context"
	"fmt"
	"sharedconfig"
)

// The Request automatically marshalled here
type Request struct {
	Name string `json:"name"`
}

// The response, which is unmarshalled automatically and returned in http or lambda response
type Response struct {
	ConfigType  string `json:"configType"`
	RequestName string `json:"requestName"`
}

// TestHandler is the unfied entry point of the module
func TestHandler(_ *context.Context, config sharedconfig.SharedConfiger, request *Request) (*Response, error) {
	fmt.Println()

	return &Response{
		ConfigType:  config.GetConfigType(),
		RequestName: request.Name,
	}, nil
}
