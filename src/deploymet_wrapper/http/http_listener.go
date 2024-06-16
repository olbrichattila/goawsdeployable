package http_listener

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func New() Listener {
	return &listen{}
}

type Listener interface {
	Start(handlers ...HandlerDef) error
	Port(int)
}

type listen struct {
	port int
}

type HandlerDef struct {
	Route   string
	Handler HandlerFunc
}
type HandlerFunc = func(ctx context.Context, payload string) (string, error)
type httpHandlerFunc = func(w http.ResponseWriter, r *http.Request)

func (l *listen) Port(port int) {
	l.port = port
}

func (l *listen) Start(handlers ...HandlerDef) error {
	for _, handler := range handlers {
		http.HandleFunc(
			handler.Route,
			l.middleware(handler.Handler),
		)
	}

	port := ":" + strconv.Itoa(l.port)
	if err := http.ListenAndServe(port, nil); err != nil {
		return err
	}

	return nil
}

func (l *listen) middleware(handler HandlerFunc) httpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context := context.Background()
		res, err := handler(context, "")
		if err != nil {
			http.Error(w, fmt.Sprintf("500 %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(res))
	}
}
