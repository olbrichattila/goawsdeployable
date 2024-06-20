package main

import "attilaolbrich.co.uk/routebuilder"

func main() {
	rb := routebuilder.New()
	rb.Port(8080)
	rb.Start()
}
