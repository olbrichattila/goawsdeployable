package main

import (
	"fmt"

	"olbrichattila.co.uk/example"
	"olbrichattila.co.uk/example2"
	connector "olbrichattila.co.uk/lambdawrapper"
)

func main() {
	listener := connector.New()
	listener.Port(8000)
	err := listener.Start(
		connector.HandlerDef{Route: "/", Handler: example.TestHandler},
		// connector.HandlerDef{Route: "/del", Handler: example.TestHandler2},
		connector.HandlerDef{Route: "/add", Handler: example2.TestHandler},
	)

	if err != nil {
		fmt.Println(err)
	}
}
