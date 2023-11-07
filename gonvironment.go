package gonvironment

import (
	"fmt"
	"os"
	"reflect"
)

type Options struct {
	AdditionalVariables []string
	EnvStruct           interface{}
	RequiredVariables   []string
}

func EnsureAndInitEnvironment(options Options) (interface{}, error) {
	additional := options.AdditionalVariables
	required := options.RequiredVariables

	envStructPtr := reflect.ValueOf(options.EnvStruct)

	if envStructPtr.Kind() != reflect.Ptr || envStructPtr.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("EnvStruct must be a pointer to a struct")
	}

	envStruct := envStructPtr.Elem()

	for _, r := range required {
		value := os.Getenv(r)
		if value == "" {
			return nil, fmt.Errorf("required environment variable %s is not set", r)
		}

		// Find the field with the same name as the environment variable
		field := envStruct.FieldByName(r)
		if field.IsValid() && field.CanSet() {
			field.SetString(value)
		}
	}

	for _, a := range additional {
		value := os.Getenv(a)
		if value != "" {
			// Find the field with the same name as the environment variable
			field := envStruct.FieldByName(a)
			if field.IsValid() && field.CanSet() {
				field.SetString(value)
			}
		}
	}

	return envStruct.Interface(), nil
}
