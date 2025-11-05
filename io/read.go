package io

import (
	"os"

	"github.com/jqkard/fn/str"
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
	lines := str.Lines(text)
	return lines, nil
}
