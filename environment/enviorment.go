package pkg_environment

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

var EnvName string
var StartEnvironment string = os.Getenv("APP_ENV")
var CurrentDir, _ = os.Getwd()

func GetEnvPath(env string) (string, error) {
	var setEnv = filepath.Join(CurrentDir, ".env")

	if env == "dev" {
		setEnv = filepath.Join(CurrentDir, "../.dev.env")
	}

	switch env {
	case "dev":
		EnvName = "DEV"
	case "local":
		EnvName = "LOCAL"
	case "staging":
		EnvName = "STAGING"
	case "production":
		EnvName = "PRODUCTION"
	default:
		if _, err := os.Stat("/.dockerenv"); err != nil {
			return "", errors.New("no environment variable set. If you not run in Docker environment, add 'APP_ENV=dev' in the front of run command")
		}
	}

	return setEnv, nil
}

func SetEnv() (string, error) {
	var envPath = filepath.Join(CurrentDir, ".env")

	if StartEnvironment == "dev" {
		envPath = filepath.Join(CurrentDir, "../.dev.env")

		err := godotenv.Load(envPath)
		if err != nil {
			return "", errors.New(fmt.Sprintf("Failed to load env file %v: %v", envPath, err))
		}
		//return EnvName, nil
	}

	switch StartEnvironment {
	case "dev":
		EnvName = "DEV"
	case "local":
		EnvName = "LOCAL"
	case "staging":
		EnvName = "STAGING"
	case "production":
		EnvName = "PRODUCTION"
	default:
		if _, err := os.Stat("/.dockerenv"); err != nil {
			return "", errors.New("no environment variable set. If you not run in Docker environment, add 'APP_ENV=dev' in the front of run command")
		}
	}

	return EnvName, nil
}

func GetDomainDependingEnv() (string, error) {

	var domain string
	if StartEnvironment == "" {
		return "", errors.New("no environment variable set")
	} else if StartEnvironment == "dev" {
		domain = "localhost"
	} else if StartEnvironment == "local" {
		domain = ".dating.local"
	} else if StartEnvironment == "staging" {
		domain = "staging.senti.live"
	} else if StartEnvironment == "production" {
		domain = "senti.live"
	} else {
		return "", errors.New("a problem occurred while setting the domain based on environment variable")
	}

	return domain, nil
}
