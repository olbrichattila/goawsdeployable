package main

import (
	"fmt"
	"io/ioutil"
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
	_ = rmDir("../prebuild")

	// TODO: do them by slice, or separeate everithing into one source and dest folder
	err := copyDir(deploymentWrapperSourceFolder, fmt.Sprintf(deploymentWrapperBuildFolder, t.buildType))
	if err != nil {
		return err
	}

	err = copyDir(eventDispatcherSourceFolder, fmt.Sprintf(eventDispatcherBuildFolder, t.buildType))
	if err != nil {
		return err
	}

	err = copyDir(handlerSourceFolder, fmt.Sprintf(handlerBuildFolder, t.buildType))
	if err != nil {
		return err
	}

	err = t.copyPackages(t.buildPackages)
	if err != nil {
		return err
	}

	err = t.createMainFile(
		t.getImports(t.buildPackages, t.buildType),
		t.getHandlers(t.buildPackages),
		t.config.Port,
	)
	if err != nil {
		return err
	}

	err = t.createTemplateFile(
		t.getModReplaces(t.buildPackages),
	)
	if err != nil {
		return err
	}

	return nil
}

func (t *application) copyPackages(packages []Package) error {
	for _, packageInfo := range packages {
		sourceFile := sourceFolder + packageInfo.Name
		destinationFile := fmt.Sprintf(buildFolder, t.buildType) + packageInfo.Name
		fmt.Printf(" -copy %s -> %s \n", sourceFile, destinationFile)
		err := copyDir(sourceFile, destinationFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *application) getImports(packages []Package, buildType string) string {
	var builder strings.Builder
	for _, packageInfo := range packages {
		builder.WriteString(fmt.Sprintf("	\"attilaolbrich.co.uk/%s\"\n", packageInfo.Name))
	}

	if buildType == typeLambda {
		builder.WriteString(fmt.Sprintf("	%s", lambdaImport))
	} else {
		builder.WriteString(fmt.Sprintf("	%s", httpImport))
	}
	return builder.String()
}

func (t *application) getModReplaces(packages []Package) string {
	var builder strings.Builder
	for _, packageInfo := range packages {
		builder.WriteString(fmt.Sprintf("replace attilaolbrich.co.uk/%s => ./%s\n\n", packageInfo.Name, packageInfo.Name))
	}

	return builder.String()
}

func (t *application) getHandlers(packages []Package) string {
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

func (t *application) createMainFile(imports, handlers string, port int) error {
	replacements := map[string]string{
		"<--imports-->":  imports,
		"<--handlers-->": handlers,
		"<--port-->":     strconv.Itoa(port),
	}

	return t.replaceInFile("template.tmpl", fmt.Sprintf(buildFolder+"main.go", t.buildType), replacements)
}

func (t *application) createTemplateFile(repacements string) error {
	replacements := map[string]string{
		"<--replacements->": repacements,
	}

	return t.replaceInFile("mod-template.tmpl", fmt.Sprintf(buildFolder+"go.mod", t.buildType), replacements)
}

func (t *application) replaceInFile(sourceFile, targetFile string, replacements map[string]string) error {
	content, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	newContent := string(content)

	for rFrom, rTo := range replacements {
		newContent = strings.ReplaceAll(newContent, rFrom, rTo)
	}

	err = ioutil.WriteFile(targetFile, []byte(newContent), 0644)
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