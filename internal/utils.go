package internal

import (
	"log"
	"os"
	"strconv"
)

func LoadEnvVariableOrFatal(envName string) string {
	env, found := os.LookupEnv(envName)
	if !found {
		log.Fatalf("failed to load env variable, %v", envName)
	}
	return env
}

func LoadUintEnvVariableOrFatal(envName string) uint64 {
	env, found := os.LookupEnv(envName)
	if !found {
		log.Fatalf("failed to load env variable, %v", envName)
	}

	parsed, err := strconv.ParseUint(env, 10, 64)
	if err != nil {
		log.Fatalf("failed to convert env variable, %v", envName)
	}

	return parsed
}
