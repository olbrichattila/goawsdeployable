package example2

import (
	"context"
)

type Request struct {
	Name string `json:"name"`
}

func TestHandler(_ *context.Context, request *Request) (*Request, error) {
	return request, nil
}
