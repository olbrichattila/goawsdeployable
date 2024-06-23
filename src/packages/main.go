// Package main is the entry point
package main

import (
	"deploymentwrapper"
	"fmt"
	"instructioner"
	"snsmiddleware"

	// config "httpconfig"
	config "localconfig"

	// connector "lambdalistener"

	"example2"
	connector "httplistener"
)

func main() {
	listener := connector.New()
	listener.Config(config.New())
	listener.Port(8082)
	err := listener.Start(
		// deploymentwrapper.HandlerDef{Route: "/", Handler: snsmiddleware.Middleware(example.TestHandler)},
		deploymentwrapper.HandlerDef{Route: "/", Handler: snsmiddleware.Middleware(instructioner.Instruction)},
		deploymentwrapper.HandlerDef{Route: "/add", Handler: example2.TestHandler},
	)

	if err != nil {
		fmt.Println(err)
	}
}
