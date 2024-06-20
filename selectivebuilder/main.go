// Package main is the entry point of the command line utility
package main

import (
	"fmt"
)

const (
	sourceFolder = "../src/packages/"
	buildFolder  = "../prebuild%s/"

	deploymentWrapperSourceFolder = "../src/deploymentwrapper/"
	deploymentWrapperBuildFolder  = "../prebuild%s/deploymentwrapper/"
	eventDispatcherSourceFolder   = "../src/eventdispatcher/"
	eventDispatcherBuildFolder    = "../prebuild%s/eventdispatcher/"
	handlerSourceFolder           = "../src/handler/"
	handlerBuildFolder            = "../prebuild%s/handler/"

	lambdaImport = "connector \"attilaolbrich.co.uk/lambdawrapper\""
	httpImport   = "connector \"attilaolbrich.co.uk/httpwrapper\""
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

	fmt.Printf(
		"Done, built into %s\n",
		fmt.Sprintf(buildFolder, app.buildType),
	)
}
