package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	buildFolder  = "./build/packages/"
	sourceFolder = "../src/packages/"
	lambdaImport = "connector \"olbrichattila.co.uk/lambdawrapper\""
	httpImport   = "connector \"olbrichattila.co.uk/httpwrapper\""
)

func main() {
	fmt.Println("Selective builder")

	conf := newYamlConfig()
	config, err := conf.pharse()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = rmDir(buildFolder)
	if err != nil {
		fmt.Println(err)
	}

	// Print the parsed data
	fmt.Printf("Port: %d\n", config.Port)

	// do for lamdba or HTTP depending on command line arg
	imports := getImports(config.Lambda.Packages)
	handlers := getHandlers(config.Lambda.Packages)

	createMainFile(imports, handlers, config.Port)
}

func getImports(packages []Package) string {
	var builder strings.Builder
	// TODO decide lambda or http

	for _, packageInfo := range packages {
		builder.WriteString(fmt.Sprintf("	\"olbrichattila.co.uk/%s\"\n", packageInfo.Name))
	}

	builder.WriteString(fmt.Sprintf("	%s", lambdaImport))
	return builder.String()
}

func getHandlers(packages []Package) string {
	var builder strings.Builder
	for _, packageInfo := range packages {
		for _, handlerInfo := range packageInfo.Functions {
			handlerParts := strings.Split(handlerInfo.Route, ":")
			builder.WriteString(
				fmt.Sprintf("		connector.HandlerDef{Route: \"%s\", Handler: %s.%s},\n", handlerParts[0], packageInfo.Name, handlerParts[1]),
			)
		}
	}

	return builder.String()
}
func createMainFile(imports, handlers string, port int) {

	inputFile := "template.tmpl"
	outputFile := "build/main.go"

	// Read the content of the input file
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	newContent := string(content)
	newContent = strings.ReplaceAll(newContent, "<--imports-->", imports)
	newContent = strings.ReplaceAll(newContent, "<--handlers-->", handlers)
	newContent = strings.ReplaceAll(newContent, "<--port-->", strconv.Itoa(port)) // todo config

	// Write the modified content to the output file
	err = ioutil.WriteFile(outputFile, []byte(newContent), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}
}
