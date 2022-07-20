package utils

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func AsJsonString(obj any) string {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	}

	return string(jsonBytes)
}

func AsYamlString(obj any) string {
	yamlStr, err := yaml.Marshal(obj)
	if err != nil {
		return err.Error()
	}

	return string(yamlStr)
}
