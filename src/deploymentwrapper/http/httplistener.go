// Package httplistener is a wrapper around http server to unify it with lambda
package httplistener

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"handler"
)

// New creates a new listener
func New() Listener {
	return &listen{
		handler: handler.New(false),
	}
}

// Listener is the interface to make Lambda and HTTP unified
type Listener interface {
	Start(handlers ...HandlerDef) error
	Port(int)
}

type listen struct {
	handler handler.StructHandler
	port    int
}

// HandlerDef is the structure how to pass a route and a handler
type HandlerDef struct {
	Route   string
	Handler handler.StructHandlerFunc
}

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

func (l *listen) middleware(structHandlerFunc handler.StructHandlerFunc) httpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading request body %s", err.Error()), http.StatusInternalServerError)
			return
		}

		response, err := l.handler.Process(structHandlerFunc, string(body))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing handler func %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(response))
	}
}
