package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-jsonnet"
	"log"
	"os"
	"path/filepath"
)

type APIDefinition struct{}
type Config struct{}

type VM struct {
	vm *jsonnet.VM
}

func (v *VM) run(vm *jsonnet.VM) (APIDefinition, error) {
	var apiDefinition APIDefinition
	template := APITemplatePath("template.jsonnet")
	jsonStr, err := vm.EvaluateFile(template)
	if err != nil {
		return APIDefinition{}, err
	}
	err = json.Unmarshal([]byte(jsonStr), &apiDefinition)
	if err != nil {
		return apiDefinition, err
	}
	return apiDefinition, nil
}

func (v *VM) build(environment string, config *Config) {
	js := AppConfigJsonString(config)
	v.vm = jsonnet.MakeVM()
	v.vm.ExtVar("env", environment)
	v.vm.ExtCode("config", js)
}

func AppConfigJsonString(config *Config) string {
	c, err := json.Marshal(config)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return string(c)
}

func APITemplatePath(file string) string {
	absPath, err := filepath.Abs(fmt.Sprintf("../../%s", file))
	if err != nil {
		fmt.Printf("Failed to read file: %s, err: %s", "template.jsonnet", err)
		os.Exit(1)
	}
	return absPath
}
