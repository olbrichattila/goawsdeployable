package example

import (
	"context"
	"fmt"

	dispather "olbrichattila.co.uk/sqs_event_dispatcher"
)

type Event struct {
	Name string `json:"name"`
}

func TestHandler(_ context.Context, payload string) (string, error) {
	fmt.Println(payload)
	str, err := dispather.NewDispatcher().Send(Event{Name: payload})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	return fmt.Sprintf("{\"RES\": \"It works%s\"}", fmt.Sprint(payload)), nil
}
