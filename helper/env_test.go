package helper

import (
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLoadDotEnv(t *testing.T) {
	err := LoadDotEnv()
	assert.Equal(t, nil, err)
}

func TestGetEnvWithDefaultValue(t *testing.T) {
	os.Setenv("SOME_ENV_VAR_KEY", "SOME_ENV_VAR_VALUE")
	value := GetEnvWithDefaultValue("SOME_ENV_VAR_KEY", "DEFAULT_VALUE")
	assert.Equal(t, "SOME_ENV_VAR_VALUE", value)

	defaultValue := GetEnvWithDefaultValue("WHO_NAMES_ENV_VARIABLE_LIKE_THIS", "DEFAULT_VALUE")
	assert.Equal(t, "DEFAULT_VALUE", defaultValue)
}
