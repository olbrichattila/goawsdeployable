package example2

import (
	"context"
	"fmt"
)

func TestHandler(_ context.Context, payload string) (string, error) {
	return fmt.Sprintf("{\"RES\": \"It works with example 2%s\"}", fmt.Sprint(payload)), nil
}
