package gottings

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type UnmarshalableField interface {
	UnmarshalEnvironmentValue(data []byte) error
}

//	type Config struct {
//	     Host: NullString `env:ENV_HOST`
//	     Port: NullInt    `env:ENV_PORT`
//	}
//
// config := &Config{}
// LoadEnv(config)
func LoadEnv(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.Elem().Kind() != reflect.Struct {
		return errors.New("expected pointer to struct")
	}
	elem := rv.Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Type().Field(i)
		envKey := field.Tag.Get("env")
		if envKey == "" {
			continue
		}
		envValue := os.Getenv(envKey)
		if envValue == "" {
			continue
		}

		fieldValue := elem.Field(i)
		if !fieldValue.CanSet() {
			return fmt.Errorf("cannot set field %s", field.Name)
		}

		var targetValue reflect.Value
		if fieldValue.Kind() == reflect.Pointer {
			if fieldValue.IsNil() {
				fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
			}
			targetValue = fieldValue.Elem()
		} else {
			targetValue = fieldValue
		}

		switch targetValue.Kind() {
		case reflect.Bool:
			value, err := strconv.ParseBool(envValue)
			if err != nil {
				return err
			}
			targetValue.SetBool(value)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value, err := strconv.ParseInt(envValue, 10, targetValue.Type().Bits())
			if err != nil {
				return err
			}
			targetValue.SetInt(value)
		case reflect.String:
			targetValue.SetString(envValue)
		case reflect.Float64, reflect.Float32:
			value, err := strconv.ParseFloat(envValue, targetValue.Type().Bits())
			if err != nil {
				return err
			}
			targetValue.SetFloat(value)
		case reflect.Struct:
			unMarshalInterface := reflect.TypeOf((*UnmarshalableField)(nil)).Elem()
			if fieldValue.Addr().Type().Implements(unMarshalInterface) {
				unmarshaler := fieldValue.Addr().Interface().(UnmarshalableField)
				if err := unmarshaler.UnmarshalEnvironmentValue([]byte(envValue)); err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("unexpected field type: %s", field.Type.Kind())
		}
	}
	return nil
}
