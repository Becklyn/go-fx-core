package env_test

import (
	"os"
	"testing"

	"github.com/Becklyn/go-fx-core/env"
	"github.com/stretchr/testify/assert"
)

func TestIsDevelopment(t *testing.T) {
	assert.False(t, env.IsDevelopment())

	os.Setenv(env.APP_ENV, "dev")

	assert.True(t, env.IsDevelopment())
}

func TestString(t *testing.T) {
	assert.Empty(t, env.String("test"))

	os.Setenv("test", "value")

	assert.Equal(t, "value", env.String("test"))
}

func TestInt(t *testing.T) {
	assert.Equal(t, -1, env.Int("test"))

	os.Setenv("test", "not an int")
	assert.Equal(t, -1, env.Int("test"))

	os.Setenv("test", "123")
	assert.Equal(t, 123, env.Int("test"))
}

func TestIntWithDefault(t *testing.T) {
	assert.Equal(t, 321, env.IntWithDefault("test", 321))

	os.Setenv("test", "not an int")
	assert.Equal(t, 321, env.IntWithDefault("test", 321))

	os.Setenv("test", "123")
	assert.Equal(t, 123, env.IntWithDefault("test", 321))

	os.Setenv("test", "-1")
	assert.Equal(t, -1, env.IntWithDefault("test", 321))
}

func TestBool(t *testing.T) {
	assert.False(t, env.Bool("test"))

	os.Setenv("test", "not a bool")
	assert.False(t, env.Bool("test"))

	os.Setenv("test", "1")
	assert.True(t, env.Bool("test"))

	os.Setenv("test", "true")
	assert.True(t, env.Bool("test"))

	os.Setenv("test", "0")
	assert.False(t, env.Bool("test"))

	os.Setenv("test", "false")
	assert.False(t, env.Bool("test"))
}

func TestBoolWithDefault(t *testing.T) {
	assert.False(t, env.BoolWithDefault("test", false))
	assert.True(t, env.BoolWithDefault("test", true))

	os.Setenv("test", "true")
	assert.True(t, env.BoolWithDefault("test", false))
	assert.True(t, env.BoolWithDefault("test", true))

	os.Setenv("test", "false")
	assert.False(t, env.BoolWithDefault("test", false))
	assert.False(t, env.BoolWithDefault("test", true))
}
