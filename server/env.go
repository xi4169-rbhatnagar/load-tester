package main

import (
	"os"
	"strconv"
)

// getEnv value from the env for the given key.
// if the ENV doesn't have the given key, it will return the defaultValue
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvInt gets env's int variables
func getEnvInt(key string, defaultValue int) (int, error) {
	if value, exists := os.LookupEnv(key); exists {
		return strconv.Atoi(value)
	}
	return defaultValue, nil
}
