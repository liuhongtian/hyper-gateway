package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Route struct {
	Path    string `yaml:"path"`
	Method  string `yaml:"method"`
	Handler string `yaml:"handler"`
}

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Routes []Route `yaml:"routes"`
}

var AppConfig Config

func LoadConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		return err
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Printf("Error parsing config file: %v", err)
		return err
	}

	return nil
}
