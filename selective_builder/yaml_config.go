package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func newYamlConfig() *ymlData {
	return &ymlData{}
}

type ymlData struct {
}

// Function represents a function within a package
type Function struct {
	Route string `yaml:"route"`
}

// Package represents a package within lambda or http
type Package struct {
	Name      string     `yaml:"name"`
	Functions []Function `yaml:"functions"`
}

// Config represents the structure of our YAML file
type Config struct {
	Port   int `yaml:"port"`
	Lambda struct {
		Packages []Package `yaml:"packages"`
	} `yaml:"lambda"`
	HTTP struct {
		Packages []Package `yaml:"packages"`
	} `yaml:"http"`
}

func (*ymlData) pharse() (Config, error) {
	// Read the YAML file
	data, err := ioutil.ReadFile("lambda.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Parse the YAML file
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config, nil

}
