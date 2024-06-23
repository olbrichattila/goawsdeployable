package instructioner

import (
	"context"
	"sharedconfig"
	"sqseventdispatcher"
)

type request struct {
	EventName  string       `json:"eventName"`
	Id         int64        `json:"id"`
	Name       string       `json:"name"`
	Parameters createParams `json:"parameters"`
}

type sQSMessage struct {
	EventName string `json:"eventName"`
	Id        int64  `json:"id"`
}

type response struct{}

func Instruction(ctx *context.Context, config sharedconfig.SharedConfiger, request *request) (*response, error) {
	if request.EventName == "create" {
		return CreateInstruction(ctx, config, request)
	}

	if request.EventName == "processInstruction" {
		return ProcessInstruction(ctx, config, request)
	}

	return &response{}, nil
}

func CreateInstruction(_ *context.Context, config sharedconfig.SharedConfiger, request *request) (*response, error) {
	instructoner := newInstructioner()
	id, err := instructoner.create(config, request.Name, request.Parameters)
	if err != nil {
		return nil, err
	}

	event := &sQSMessage{
		EventName: "processInstruction",
		Id:        id,
	}

	dispatcher := sqseventdispatcher.NewDispatcher(config.GetSQSConfig())
	err = dispatcher.Send(*event)
	if err != nil {
		return nil, err
	}
	return &response{}, nil
}

func ProcessInstruction(_ *context.Context, config sharedconfig.SharedConfiger, request *request) (*response, error) {
	instructoner := newInstructioner()
	err := instructoner.process(config, request.Id)
	if err != nil {
		return nil, err
	}

	return &response{}, nil
}
