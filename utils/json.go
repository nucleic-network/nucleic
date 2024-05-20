package utils

import (
	"encoding/json"
	"os"
)

// ParseJsonFromFile parses a json file into a slice of type T
func ParseJsonFromFile[T any](path string) ([]T, error) {
	var result []T

	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
