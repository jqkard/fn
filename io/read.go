// The io package contains functions related to reading files
package io

import (
	"encoding/json"
	"os"
	"strings"
)

func ReadFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ReadLines(path string) ([]string, error) {
	text, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(text, "\n")
	return lines, nil
}

func ReadJSON[T any](path string) (T, error) {
	var item T
	bytes, err := os.ReadFile(path)
	if err != nil {
		return item, err
	}
	err = json.Unmarshal(bytes, &item)
	if err != nil {
		return item, err
	}
	return item, nil
}
