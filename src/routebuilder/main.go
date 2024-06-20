// Package routebuilder will create the routes and dependin on the import, decide if it is lambda or http
package routebuilder

import (

	// connector "attilaolbrich.co.uk/lambdawrapper"

	connector "attilaolbrich.co.uk/deploymentwrapper/http"
	"attilaolbrich.co.uk/example"
	"attilaolbrich.co.uk/example2"
)

// RouteBuilder is a wrapper around ASW/HTTP build
type RouteBuilder interface {
	Port(port int)
	Start() error
}

type routeBuilt struct {
	listener connector.Listener
}

// New creates a new builder instance
func New() RouteBuilder {
	return &routeBuilt{
		listener: connector.New(),
	}
}

func (t *routeBuilt) Port(port int) {
	t.listener.Port(port)

}

func (t *routeBuilt) Start() error {
	return t.listener.Start(
		connector.HandlerDef{Route: "/", Handler: example.TestHandler},
		connector.HandlerDef{Route: "/add", Handler: example2.TestHandler},
	)
}
