package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadDotEnv() error {
	return godotenv.Load(ProjectRootPath() + "/.env")
}

func GetEnvWithDefaultValue(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
