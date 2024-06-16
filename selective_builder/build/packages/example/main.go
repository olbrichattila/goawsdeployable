package example

import (
	"context"
	"fmt"
)

func TestHandler(_ context.Context, payload string) (string, error) {
	return fmt.Sprintf("{\"RES\": \"It works%s\"}", fmt.Sprint(payload)), nil
}
