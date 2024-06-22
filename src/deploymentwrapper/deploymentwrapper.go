package deploymentwrapper

import (
	"handler"
	"sharedconfig"
)

// Listener is the interface to make Lambda and HTTP unified
type Listener interface {
	Config(sharedconfig.SharedConfiger)
	Start(handlers ...HandlerDef) error
	Port(int)
}

// HandlerDef is the structure how to pass a route and a handler
type HandlerDef struct {
	Route   string
	Handler handler.StructHandlerFunc
}
