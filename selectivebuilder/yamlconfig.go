package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	configYmlFileName = "config.yaml"
)

func newYamlConfig() *ymlData {
	return &ymlData{}
}

type ymlData struct {
}

type function struct {
	Route    string `yaml:"route"`
	SNSRoute string `yaml:"snsroute"`
}

// Package is the structure from yaml package
type Package struct {
	Name      string     `yaml:"name"`
	Functions []function `yaml:"functions"`
}

// Config is the structure from yaml package
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
	data, err := os.ReadFile(configYmlFileName)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config, nil
}
