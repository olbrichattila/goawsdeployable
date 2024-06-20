// Package example is just an example how to create a module accross AWS lambda or HTTP with the same code
package example

import (
	"context"
	"fmt"

	dispather "attilaolbrich.co.uk/sqs_eventdispatcher"
)

// Request data automaticall marhalled here
type Request struct {
	Name string `json:"name"`
}

// Response what we want to be returned as HTTP or Lambda
type Response struct {
	ResponseName string `json:"respopnseName"`
}

// TestHandler is the unfied entry point of the module
func TestHandler(_ *context.Context, request *Request) (*Response, error) {
	fmt.Println(request)
	str, err := dispather.NewDispatcher().Send(*request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	return &Response{
		ResponseName: "It is the response",
	}, nil
}
