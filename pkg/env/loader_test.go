package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_LoadEnv_WhenFieldsAreValidAndVariablesExist(t *testing.T) {
	env := struct {
		DbHost string `env:"TEST_DB_HOST"`
		DbPort int    `env:"TEST_DB_PORT"`
		DbName string `env:"TEST_DB_NAME"`
		DbUser string `env:"TEST_DB_USER"`
		DbPass string `env:"TEST_DB_PASS"`
	}{}

	os.Setenv("TEST_DB_HOST", "testhost")
	os.Setenv("TEST_DB_PORT", "3306")
	os.Setenv("TEST_DB_NAME", "testdb")
	os.Setenv("TEST_DB_USER", "testuser")
	os.Setenv("TEST_DB_PASS", "testpass")

	err := LoadEnv(&env)
	require.Nil(t, err)

	assert.Equal(t, "testhost", env.DbHost)
	assert.Equal(t, 3306, env.DbPort)
	assert.Equal(t, "testdb", env.DbName)
	assert.Equal(t, "testuser", env.DbUser)
	assert.Equal(t, "testpass", env.DbPass)
}

func Test_LoadEnv_WhenFieldsAreInvalid(t *testing.T) {
	env := struct {
		DbHost string  `env:"TEST_DB_HOST"`
		DbPort int     `env:"TEST_DB_PORT"`
		DbName string  `env:"TEST_DB_NAME"`
		DbUser string  `env:"TEST_DB_USER"`
		DbPass float32 `env:"TEST_DB_PASS"`
	}{}

	os.Setenv("TEST_DB_HOST", "testhost")
	os.Setenv("TEST_DB_PORT", "3306")
	os.Setenv("TEST_DB_NAME", "testdb")
	os.Setenv("TEST_DB_USER", "testuser")
	os.Setenv("TEST_DB_PASS", "testpass")

	err := LoadEnv(&env)
	assert.EqualError(t, err, "float32 is not a valid type")
}

func Test_LoadEnv_WhenVariablesNotExist(t *testing.T) {
	env := struct {
		DbHost string `env:"TEST_DB_HOST"`
		DbPort int    `env:"TEST_DB_PORT"`
		DbName string `env:"TEST_DB_NAME"`
		DbUser string `env:"TEST_DB_USER"`
		DbPass string `env:"TEST_DB_PASS"`
	}{}

	os.Setenv("TEST_DB_HOST", "testhost")
	os.Setenv("TEST_DB_PORT", "3306")
	os.Setenv("TEST_DB_NAME", "testdb")
	os.Unsetenv("TEST_DB_USER")
	os.Setenv("TEST_DB_PASS", "testpass")

	err := LoadEnv(&env)
	assert.EqualError(t, err, "TEST_DB_USER was not found")
}
