// Package env implements utilities for reading environment variables.
package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// Load environment variable values into fields declared in env using the tag `env:"VAR_NAME"`.
// The supported field types are string and int.
func LoadEnv(env any) error {
	elements := reflect.ValueOf(env).Elem()
	types := elements.Type()

	for i := 0; i < elements.NumField(); i++ {
		field := elements.Field(i)

		varName := types.Field(i).Tag.Get("env")
		varValue, ok := os.LookupEnv(varName)
		if !ok {
			return fmt.Errorf("%s was not found", varName)
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(varValue)
		case reflect.Int:
			value, err := strconv.Atoi(varValue)
			if err != nil {
				return err
			}
			field.SetInt(int64(value))
		default:
			return fmt.Errorf("%s is not a valid type", field.Kind().String())
		}
	}

	return nil
}
