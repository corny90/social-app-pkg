package pkg_environment

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

var EnvName string
var startEnvironment string = os.Getenv("APP_ENV")
var currentDir, _ = os.Getwd()

func GetEnvPath(env string) (string, error) {
	var setEnv string

	switch env {
	case "dev":
		setEnv = filepath.Join(currentDir, "../.dev.env")
		EnvName = "LOCAL"
	case "docker:staging":
		setEnv = filepath.Join(currentDir, ".env")
		EnvName = "STAGING"
	case "docker:local":
		setEnv = filepath.Join(currentDir, ".env")
		EnvName = "LOCAL·DOCKER"
	default:
		if _, err := os.Stat("/.dockerenv"); err != nil {
			return "", errors.New("no environment variable set. If you not run in Docker environment, add 'APP_ENV=dev' in the front of run command")
		}
		EnvName = "STAGING"
		return env, nil

	}

	return setEnv, nil
}

func SetEnv() (string, error) {
	envPath, err := GetEnvPath(startEnvironment)
	if err != nil {
		return "", err
	}

	if envPath != "" {
		err := godotenv.Load(envPath)
		if err != nil {
			return "", errors.New(fmt.Sprintf("Failed to load env file %v: %v", envPath, err))
		}
		return EnvName, nil
	}

	return EnvName, nil
}
