package utils

import (
	"fmt"
	"io/ioutil"
)

func LoadJson(filePath string) (string, error) {
	jsonBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("can't read file at %s... %w", filePath, err)
	}
	return string(jsonBytes), nil
}