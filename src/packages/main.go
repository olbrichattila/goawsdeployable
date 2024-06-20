// Package main is the entry point
package main

import (
	"fmt"

	// connector "attilaolbrich.co.uk/lambdawrapper"

	connector "attilaolbrich.co.uk/deploymentwrapper/http"
	"attilaolbrich.co.uk/example"
	"attilaolbrich.co.uk/example2"
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
