// Package example is just an example how to create a module accross AWS lambda or HTTP with the same code
package example

// TODO: refactor this, and make it an adapter between SNS and HTML to be usable generally
import (
	"context"
	"sharedconfig"
	"sqseventdispatcher"
)

// TestHandler is the unfied entry point of the module
func TestHandler(_ *context.Context, config sharedconfig.SharedConfiger, request string) (string, error) {
	dispatcher := sqseventdispatcher.NewDispatcher(config.GetSQSConfig())
	dispatcher.SendString(request)
	return request, nil
}
