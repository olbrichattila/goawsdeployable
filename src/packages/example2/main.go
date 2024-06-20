// Package example2 is just an example how to create a module accross AWS lambda or HTTP with the same code
package example2

import (
	"context"
)

// The Request automatically marshalled here
type Request struct {
	Name string `json:"name"`
}

// TestHandler is the unfied entry point of the module
func TestHandler(_ *context.Context, request *Request) (*Request, error) {
	return request, nil
}
