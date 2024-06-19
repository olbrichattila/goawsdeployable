// Package handler will get a handler func like func(*context.Context, any) (any, error)
// where any can only be a struct pointer
// Example:
//
//	  func testHandler(ctx *context.Context, par2 *request) (*response, error) {
//		   return &response{Responsed: "OK response"}, nil
//	  }
//
// and a string wihth a json, which will be rendered to the second parameter of the func
package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
)

// StructHandler will get ahandler func, and a json string
// Using this as a middleware layer between http, and similar packages to unify and simplify
// parameter passing and receiving from JSON format
type StructHandler interface {
	Process(handleFunc StructHandlerFunc, jsonStr string) (string, error)
}

// New creates a new struct handler instanc
func New(validateJSONTags bool) StructHandler {
	return &handle{
		validateJSONTags: validateJSONTags,
	}
}

// StructHandlerFunc is an interface which will be validated to
// func testHandler(*context.Context, any) (any, error)
// where any can be a struct only
type StructHandlerFunc interface{}

type handle struct {
	validateJSONTags bool
}

func (t handle) Process(handleFunc StructHandlerFunc, jsonStr string) (string, error) {
	err := t.verify(handleFunc)
	if err != nil {
		return "", err
	}

	return t.processHandleFunc(handleFunc, jsonStr)
}

func (t handle) processHandleFunc(handlerFunc StructHandlerFunc, jsonStr string) (string, error) {
	ctx := context.Background()
	funcValue := reflect.ValueOf(handlerFunc)
	funcType := funcValue.Type()

	ctxValue := reflect.ValueOf(&ctx)

	dataType := funcType.In(1).Elem()
	dataPtr := reflect.New(dataType)
	data := dataPtr.Elem()

	err := json.Unmarshal([]byte(jsonStr), data.Addr().Interface())
	if err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	result := funcValue.Call([]reflect.Value{ctxValue, dataPtr})
	if !result[1].IsNil() {
		errValue := result[1].Interface()
		if err, ok := errValue.(error); ok {
			return "", err
		}

		return "", fmt.Errorf("unknown error in handler parser")

	}

	m, err := json.Marshal(result[0].Elem().Interface())
	if err != nil {
		return "", err
	}
	return string(m), nil
}

func (t handle) verify(hFunc StructHandlerFunc) error {
	funcValue := reflect.ValueOf(hFunc)
	if funcValue.Kind() != reflect.Func {
		return fmt.Errorf("handler parameter is not a function")
	}

	funcType := funcValue.Type()
	if funcType.NumIn() != 2 {
		return fmt.Errorf("handler parameter function must have two parameters")
	}

	if funcType.In(0).Kind() != reflect.Ptr || funcType.In(0).Elem() != reflect.TypeOf((*context.Context)(nil)).Elem() {
		return fmt.Errorf("handler parameter first parameter must be context.Context")
	}

	if funcType.In(1).Kind() != reflect.Ptr || funcType.In(1).Elem().Kind() != reflect.Struct {
		return fmt.Errorf("handler parameter second parameter must be a struct")
	}

	if funcType.NumOut() != 2 {
		return fmt.Errorf("handler return parameter count sould be 2")
	}

	if funcType.Out(0).Kind() != reflect.Ptr || funcType.Out(0).Elem().Kind() != reflect.Struct {
		return fmt.Errorf("handler return parameter first parameter should be a struct")
	}

	if funcType.Out(1) != reflect.TypeOf((*error)(nil)).Elem() {
		return fmt.Errorf("handler return parameter second parameter should be nil or error")
	}

	if t.validateJSONTags {
		dataType := funcType.In(1).Elem()
		hasJSONTag := true
		for i := 0; i < dataType.NumField(); i++ {
			field := dataType.Field(i)
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				hasJSONTag = false
				break
			}
		}

		if !hasJSONTag {
			return fmt.Errorf("handler parameter struct must `json` must contain tags")
		}
	}

	return nil
}
