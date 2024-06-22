// Package main is the entry point
package main

import (
	"deploymentwrapper"
	"fmt"
	"httpconfig"

	// connector "lambdalistener"

	"example"
	"example2"
	connector "httplistener"
)

func main() {
	listener := connector.New()
	listener.Config(httpconfig.New())
	listener.Port(8080)
	err := listener.Start(
		deploymentwrapper.HandlerDef{Route: "/", Handler: example.TestHandler},
		deploymentwrapper.HandlerDef{Route: "/add", Handler: example2.TestHandler},
	)

	if err != nil {
		fmt.Println(err)
	}
}
