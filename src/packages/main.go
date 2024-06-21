// Package main is the entry point
package main

import (
	"fmt"

	// connector "lambdalistener"

	"example"
	"example2"
	connector "httplistener"
)

func main() {
	listener := connector.New()
	listener.Port(8080)
	err := listener.Start(
		connector.HandlerDef{Route: "/", Handler: example.TestHandler},
		connector.HandlerDef{Route: "/add", Handler: example2.TestHandler},
	)

	if err != nil {
		fmt.Println(err)
	}
}
