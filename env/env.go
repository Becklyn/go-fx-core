package env

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"os"
	"strconv"
)

const APP_ENV = "APP_ENV"

var Module = fx.Invoke(
	useEnvironment,
)

func useEnvironment(logger *logrus.Logger) {
	production := false
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		if _, err := os.Stat("../.env"); os.IsNotExist(err) {
			production = true
			_ = os.Setenv(APP_ENV, "production")
		}
	}

	if !production {
		if err := godotenv.Load(".env"); err != nil {
			if err := godotenv.Load("../.env"); err != nil {
				logger.Warn("Development environment file could not be located")
			}
		}
	}

	if os.Getenv(APP_ENV) == "" {
		_ = os.Setenv(APP_ENV, "development")
	}
}

func GetEnvironment() string {
	return os.Getenv(APP_ENV)
}

func IsDevelopment() bool {
	return os.Getenv(APP_ENV) == "development"
}

func String(key string) string {
	return os.Getenv(key)
}

func Int(key string) int {
	value := os.Getenv(key)

	if len(value) == 0 {
		return 0
	}

	intValue, err := strconv.Atoi(value)

	if err != nil {
		return 0
	}

	return intValue
}

func Bool(key string) bool {
	value := os.Getenv(key)
	return value == "true" || value == "1"
}
