package env

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

const APP_ENV = "APP_ENV"

type Env struct{}

var Module = fx.Provide(
	newEnvironment,
)

func newEnvironment() *Env {
	if err := godotenv.Load(".env.local"); err != nil {
		_ = godotenv.Load("../.env.local")
	}

	if err := godotenv.Load(".env"); err != nil {
		_ = godotenv.Load("../.env")
	}

	if os.Getenv(APP_ENV) == "" {
		_ = os.Setenv(APP_ENV, "production")
	}

	return &Env{}
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

func StringWithDefault(key string, defaultValue string) string {
	value := String(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func Int(key string) int {
	value := os.Getenv(key)

	if len(value) == 0 {
		return -1
	}

	intValue, err := strconv.Atoi(value)

	if err != nil {
		return -1
	}

	return intValue
}

func IntWithDefault(key string, defaultValue int) int {
	value := os.Getenv(key)

	if len(value) == 0 {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)

	if err != nil {
		return defaultValue
	}

	return intValue
}

func Bool(key string) bool {
	value := strings.ToLower(os.Getenv(key))
	return value == "true" || value == "1"
}

func BoolWithDefault(key string, defaultValue bool) bool {
	if _, isSet := os.LookupEnv(key); !isSet {
		return defaultValue
	}
	return Bool(key)
}
