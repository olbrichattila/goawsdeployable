package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func newBuildValidator() *buildValidator {
	return &buildValidator{}
}

type funcLookupResult struct {
	packageName string
	name        string
	found       bool
}

type buildValidator struct {
}

func (t *buildValidator) validate(packages []Package) error {

	var errors []error
	for _, buildPackage := range packages {
		err := t.validatePackage(buildPackage)
		if err != nil {
			errors = append(errors, err...)
		}
	}

	return t.convertErrorsToError(errors)
}

func (t *buildValidator) init(buildPackage Package) ([]string, []*funcLookupResult, *[]fileInfo, error) {

	packageDirectory := sourceFolder + buildPackage.Name
	if _, err := os.Stat(packageDirectory); os.IsNotExist(err) {
		return nil, nil, nil, fmt.Errorf("the %s package does not exists", buildPackage.Name)
	}

	routePaths := make([]string, len(buildPackage.Functions))
	functionLookups := make([]*funcLookupResult, len(buildPackage.Functions))

	for i, fn := range buildPackage.Functions {
		routeParts := strings.Split(fn.Route, ":")
		routePaths[i] = routeParts[0]
		functionLookups[i] = &funcLookupResult{name: routeParts[1], packageName: buildPackage.Name}
	}

	files, err := readDir(packageDirectory)
	if err != nil {
		return nil, nil, nil, err
	}
	return routePaths, functionLookups, files, nil

}

func (t *buildValidator) convertErrorsToError(errors []error) error {
	if len(errors) == 0 {
		return nil
	}

	var builder strings.Builder
	for _, e := range errors {
		builder.WriteString(fmt.Sprintln(e.Error()))
	}
	return fmt.Errorf(builder.String())
}

func (t *buildValidator) validatePackage(buildPackage Package) []error {
	routePaths, functionLookups, files, err := t.init(buildPackage)
	if err != nil {
		return []error{err}
	}

	for _, f := range *files {
		if filepath.Ext(f.path) == ".go" {
			content, err := t.loadFile(f.path)
			if err != nil {
				return []error{err}
			}

			err = t.valudateFunctions(functionLookups, content)
			if err != nil {
				return []error{err}
			}
		}
	}

	return t.getCombinerErrors(functionLookups, routePaths)
}

func (t *buildValidator) valudateFunctions(functionLookups []*funcLookupResult, content string) error {
	for _, funcLookup := range functionLookups {
		if !funcLookup.found {
			match, err := t.validateFunctionsExistsInGoFile(funcLookup.name, content)
			if err != nil {
				return err
			}
			funcLookup.found = match
		}
	}

	return nil
}

func (t *buildValidator) getCombinerErrors(functionLookups []*funcLookupResult, routePaths []string) []error {
	lookupErrors := t.functinLookupErrors(functionLookups)
	duplicateErrors := t.duplicateErrors(routePaths)

	return append(lookupErrors, duplicateErrors...)

}

func (t *buildValidator) functinLookupErrors(functionLookups []*funcLookupResult) []error {
	var errors []error
	for _, funcLookup := range functionLookups {
		if !funcLookup.found {
			errors = append(errors, fmt.Errorf("the %s hanler not implemented in package %s", funcLookup.name, funcLookup.packageName))
		}
	}

	return errors
}

func (t *buildValidator) duplicateErrors(routePaths []string) []error {
	var errors []error
	duplicates := t.findRepeatedElements(routePaths)
	for _, duplicate := range duplicates {
		errors = append(errors, fmt.Errorf("the route %s is duplicate accross your packages", duplicate))
	}

	return errors
}

func (t *buildValidator) loadFile(sourceFile string) (string, error) {
	content, err := os.ReadFile(sourceFile)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func (t *buildValidator) validateFunctionsExistsInGoFile(funcname, content string) (bool, error) {
	escapedStr := regexp.QuoteMeta(funcname)
	match, err := regexp.MatchString(fmt.Sprintf("func.*%s", escapedStr), string(content))
	if err != nil {
		return false, err
	}

	return match, nil
}

func (t *buildValidator) findRepeatedElements(elements []string) []string {
	var duplicates []string
	seen := make(map[string]bool)

	for _, element := range elements {
		if seen[element] {
			duplicates = append(duplicates, element)
		} else {
			seen[element] = true
		}
	}
	return duplicates
}
