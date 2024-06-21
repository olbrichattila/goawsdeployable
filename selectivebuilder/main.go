// Package main is the entry point of the command line utility
package main

import (
	"fmt"
)

const (
	sourceFolder = "../src/packages/"
	buildFolder  = "../src/built/%s/"

	lambdaImport = "connector \"lambdalistener\""
	httpImport   = "connector \"httplistener\""
	typeLambda   = "lambda"
	typeHTTP     = "http"
)

// Add bash script to build, alternatively go mod tidy all subfolders and deployment wrapppers and so
// or this command add a recursive function to execute go mod tidy for all destination folder

func main() {
	fmt.Println("Selective builder")

	app := newApp()

	err := app.init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = app.validate()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = app.build()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Done.")
}
