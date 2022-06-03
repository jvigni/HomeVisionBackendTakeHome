package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func LoadJson(filePath string) (string, error) {
	jsonBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("can't read file at %s... %w", filePath, err)
	}
	return string(jsonBytes), nil
}

func ParseJson[T any](response *http.Response) (*T, error) {
	var result T
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &result, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return &result, err
	}
	return &result, nil
}
