package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func LoadEnvVariableOrFatal(envName string) string {
	env, found := os.LookupEnv(envName)
	if !found {
		logrus.Fatalf("failed to load env variable, %v", envName)
	}
	return env
}

func LoadUintEnvVariableOrFatal(envName string) uint64 {
	env, found := os.LookupEnv(envName)
	if !found {
		logrus.Fatalf("failed to load env variable, %v", envName)
	}

	parsed, err := strconv.ParseUint(env, 10, 64)
	if err != nil {
		logrus.Fatalf("failed to convert env variable, %v", envName)
	}

	return parsed
}
