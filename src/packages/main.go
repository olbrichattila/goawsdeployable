package main

import (
	"fmt"

	// connector "olbrichattila.co.uk/lambdawrapper"

	"olbrichattila.co.uk/example"
	"olbrichattila.co.uk/example2"
	connector "olbrichattila.co.uk/httpwrapper"
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
