package env

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

const APP_ENV = "APP_ENV"

var Module = fx.Invoke(
	useEnvironment,
)

func useEnvironment() {
	if err := godotenv.Load(".env.local"); err != nil {
		_ = godotenv.Load("../.env.local")
	}

	if err := godotenv.Load(".env"); err != nil {
		_ = godotenv.Load("../.env")
	}

	if os.Getenv(APP_ENV) == "" {
		_ = os.Setenv(APP_ENV, "production")
	}
}

func GetEnvironment() string {
	return os.Getenv(APP_ENV)
}

func IsDevelopment() bool {
	return strings.Contains(strings.ToLower(os.Getenv(APP_ENV)), "dev")
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
	value := strings.ToLower(os.Getenv(key))
	return value == "true" || value == "1"
}
