package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func newApp() *application {
	return &application{}
}

type application struct {
	buildPackages []Package
	config        Config
	buildType     string
}

func (t *application) init() error {
	var err error
	t.buildType, err = t.arg()
	if err != nil {
		return err
	}

	conf := newYamlConfig()
	t.config, err = conf.pharse()
	if err != nil {
		return err
	}

	if t.buildType == typeLambda {
		t.buildPackages = t.config.Lambda.Packages
	} else {
		t.buildPackages = t.config.HTTP.Packages
	}

	return nil
}

func (t *application) validate() error {
	buildValidator := newBuildValidator()
	err := buildValidator.validate(t.buildPackages)
	if err != nil {
		return err
	}

	return nil
}

func (t *application) build() error {

	err := t.createMainFile(
		t.getImports(t.buildPackages, t.buildType),
		t.getHandlers(t.buildPackages),
		t.config.Port,
	)
	if err != nil {
		return err
	}

	return nil
}

func (t *application) getImports(packages []Package, buildType string) string {
	needSNSMiddlewareImport := false
	var builder strings.Builder
	for _, packageInfo := range packages {
		if !needSNSMiddlewareImport {
			for _, functions := range packageInfo.Functions {
				sNSHandlerParts := strings.Split(functions.SNSRoute, ":")
				if len(sNSHandlerParts) == 2 {
					needSNSMiddlewareImport = true
				}
			}

		}

		builder.WriteString(fmt.Sprintf("	\"%s\"\n", packageInfo.Name))
	}

	if needSNSMiddlewareImport {
		builder.WriteString(fmt.Sprintf("	%s\n", sNSMiddlewareImport))
	}

	if buildType == typeLambda {
		builder.WriteString(fmt.Sprintf("	%s\n", lambdaImport))
		builder.WriteString(fmt.Sprintf("	%s", lambdaConfigImport))
	} else {
		builder.WriteString(fmt.Sprintf("	%s\n", httpImport))
		builder.WriteString(fmt.Sprintf("	%s", httpConfigImport))
	}
	return builder.String()
}

func (t *application) getHandlers(packages []Package) string {
	var builder strings.Builder
	for _, packageInfo := range packages {
		for _, handlerInfo := range packageInfo.Functions {
			handlerParts := strings.Split(handlerInfo.Route, ":")

			if len(handlerParts) == 2 {
				builder.WriteString(
					fmt.Sprintf("		deploymentwrapper.HandlerDef{Route: \"%s\", Handler: %s.%s},\n", handlerParts[0], packageInfo.Name, handlerParts[1]),
				)
			}

			sNSHandlerParts := strings.Split(handlerInfo.SNSRoute, ":")
			if len(sNSHandlerParts) == 2 {
				builder.WriteString(
					fmt.Sprintf("		deploymentwrapper.HandlerDef{Route: \"%s\", Handler: snsmiddleware.Middleware(%s.%s)},\n", sNSHandlerParts[0], packageInfo.Name, sNSHandlerParts[1]),
				)
			}
		}
	}

	return builder.String()
}

func (t *application) createMainFile(imports, handlers string, port int) error {
	replacements := map[string]string{
		"<--imports-->":  imports,
		"<--handlers-->": handlers,
		"<--port-->":     strconv.Itoa(port),
	}

	return t.replaceInFile("template.tmpl", fmt.Sprintf(buildFolder+"main.go", t.buildType), replacements)
}

func (t *application) replaceInFile(sourceFile, targetFile string, replacements map[string]string) error {
	content, err := os.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	newContent := string(content)

	for rFrom, rTo := range replacements {
		newContent = strings.ReplaceAll(newContent, rFrom, rTo)
	}

	err = os.WriteFile(targetFile, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (t *application) arg() (string, error) {
	args := os.Args
	errorMessage := fmt.Errorf("usage selective_builder <lambda|http>")

	if len(os.Args) < 2 {
		return "", errorMessage
	}

	arg := strings.ToLower(args[1])
	if arg == typeLambda || arg == typeHTTP {
		return arg, nil
	}

	return "", errorMessage
}
