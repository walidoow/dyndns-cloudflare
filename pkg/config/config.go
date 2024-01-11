package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"reflect"
)

func ReadConfigFile(configPath string) []byte {
	body, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func DecodeYAMLConfig(yamlData []byte, out interface{}) {
	if reflect.ValueOf(out).Kind() != reflect.Pointer || reflect.ValueOf(out).IsNil() {
		log.Fatal("Output parameter must be a non-nil pointer")
	}

	err := yaml.Unmarshal(yamlData, out)
	if err != nil {
		log.Fatal(err)
	}
}
