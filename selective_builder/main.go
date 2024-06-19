package main

import (
	"fmt"
)

const (
	sourceFolder = "../src/packages/"
	buildFolder  = "../prebuild_%s/"

	deploymentWrapperSourceFolder = "../src/deployment_wrapper/"
	deploymentWrapperBuildFolder  = "../prebuild_%s/deployment_wrapper/"
	eventDispatcherSourceFolder   = "../src/event_dispatcher/"
	eventDispatcherBuildFolder    = "../prebuild_%s/event_dispatcher/"
	handlerSourceFolder           = "../src/handler/"
	handlerBuildFolder            = "../prebuild_%s/handler/"

	lambdaImport = "connector \"attilaolbrich.co.uk/lambdawrapper\""
	httpImport   = "connector \"attilaolbrich.co.uk/httpwrapper\""
	typeLambda   = "lambda"
	typeHttp     = "http"
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
