package main

import (
	"io/ioutil"
	"log"

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

type Function struct {
	Route string `yaml:"route"`
}

type Package struct {
	Name      string     `yaml:"name"`
	Functions []Function `yaml:"functions"`
}

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
	data, err := ioutil.ReadFile(configYmlFileName)
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
