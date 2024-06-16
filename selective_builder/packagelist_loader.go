package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func newEnvLister() *buildList {
	return &buildList{}
}

type packageInfo struct {
	name      string
	functions []string
}

type buildList struct {
}

func (b *buildList) load() ([]*packageInfo, error) {
	var packages []*packageInfo

	file, err := os.Open("./packagelist.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Get the current line
		if line != "" {
			if b.isFunctionName(line) {
				packageLen := len(packages)
				if packageLen == 0 {
					return nil, fmt.Errorf("the definition cannot start with function name")
				}

				lastPackage := packages[len(packages)-1]
				lastPackage.functions = append(lastPackage.functions, strings.TrimSpace(line))
			} else {
				packages = append(packages, &packageInfo{name: line})
			}
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return packages, nil
}

func (b *buildList) isFunctionName(s string) bool {
	if len(s) == 0 {
		return false
	}
	firstRune := rune(s[0])
	return unicode.IsSpace(firstRune) || firstRune == '\t'
}
