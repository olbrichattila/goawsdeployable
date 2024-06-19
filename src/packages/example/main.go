package example

import (
	"context"
	"fmt"

	dispather "attilaolbrich.co.uk/sqs_event_dispatcher"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	ResponseName string `json:"respopnseName"`
}

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
