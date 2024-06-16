package main

import (
	"fmt"
	"io/ioutil"
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

	environments, err := newEnvLister().load()
	if err != nil {
		fmt.Println(err)
	}

	err = rmDir(buildFolder)
	if err != nil {
		fmt.Println(err)
	}

	for _, env := range environments {
		src := sourceFolder + env.name
		dest := buildFolder + env.name
		fmt.Printf("- copy %s to %s\n", src, dest)

		err = copyDir(src, dest)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	imports := getImports(environments)
	handlers := getHandlers(environments)

	createMainFile(imports, handlers)
}

func getImports(packages []*packageInfo) string {
	var builder strings.Builder
	// TODO decide lambda or http

	for _, packageInfo := range packages {
		builder.WriteString(fmt.Sprintf("	\"olbrichattila.co.uk/%s\"\n", packageInfo.name))
	}

	builder.WriteString(fmt.Sprintf("	%s", lambdaImport))
	return builder.String()
}

func getHandlers(packages []*packageInfo) string {
	var builder strings.Builder
	for _, packageInfo := range packages {
		fmt.Println(packageInfo.functions)
		for _, handlerInfo := range packageInfo.functions {
			handlerParts := strings.Split(handlerInfo, "=")
			builder.WriteString(
				fmt.Sprintf("		connector.HandlerDef{Route: \"%s\", Handler: %s.%s},\n", handlerParts[0], packageInfo.name, handlerParts[1]),
			)
		}
	}

	return builder.String()
}
func createMainFile(imports, handlers string) {

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
	newContent = strings.ReplaceAll(newContent, "<--port-->", "8000") // todo config

	// Write the modified content to the output file
	err = ioutil.WriteFile(outputFile, []byte(newContent), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}
}
