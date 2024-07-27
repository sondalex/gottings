package gottings

import (
	"errors"
	"fmt"
	"reflect"
)

type Options map[string]interface{}

type UnmarshalableOption interface {
	UnmarshalOption(v any) error
}

// Example
//
//	 type Config struct {
//		    Port    string
//	 }
//
//	 func NewConfig() (*config, error) {
//	     config := Config{}
//	     err := LoadOptions(map[string]interface{}{
//	     "Port": 1312,
//	     }, &config)
//
//	     if err != nil {
//	         return nil, err
//	     }
//	     return &config, nil
//	 }
func LoadOptions(options Options, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.Elem().Kind() != reflect.Struct {
		return errors.New("expected pointer to struct")
	}
	elem := rv.Elem()
	for i := 0; i < elem.NumField(); i++ {
		fieldValue := elem.Field(i)
		var targetValue reflect.Value
		if fieldValue.Kind() == reflect.Pointer {
			if fieldValue.IsNil() {
				fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
			}
			targetValue = fieldValue.Elem()
		} else {
			targetValue = fieldValue
		}
		fieldName := elem.Type().Field(i).Name
		value, ok := options[fieldName]
		if !ok {
			continue
		}
		switch targetValue.Kind() {
		case reflect.Bool:
			var v bool
			switch n := value.(type) {
			case bool:
				v = bool(n)
			case *bool:
				v = bool(*n)
			default:
				return fmt.Errorf("failed to set bool field %s due to type mismatch", fieldName)
			}
			targetValue.SetBool(v)
		case reflect.Int:
			v, err := ToInt64(value)
			if err != nil {
				return fmt.Errorf("failed to set field %s due to type mismatch err %s", fieldName, err)
			}
			targetValue.SetInt(v)

		case reflect.Int8:
			v, err := ToInt64(value)
			if err != nil {
				return fmt.Errorf("failed to set field %s due to type mismatch err %s", fieldName, err)
			}
			targetValue.SetInt(int64(v))

		case reflect.Int16:
			v, err := ToInt64(value)
			if err != nil {
				return fmt.Errorf("failed to set field %s due to type mismatch err %s", fieldName, err)
			}
			targetValue.SetInt(int64(v))

		case reflect.Int32:
			v, err := ToInt64(value)
			if err != nil {
				return fmt.Errorf("failed to set field %s due to type mismatch err %s", fieldName, err)
			}
			targetValue.SetInt(v)

		case reflect.Int64:
			v, err := ToInt64(value)
			if err != nil {
				return fmt.Errorf("failed to set field %s due to type mismatch err %s", fieldName, err)
			}
			targetValue.SetInt(v)

		case reflect.String:
			var v string
			switch n := value.(type) {
			case string:
				v = string(n)
			case *string:
				v = string(*n)
			default:
				return fmt.Errorf("failed to set string field %s due to type mismatch", fieldName)
			}
			targetValue.SetString(v)
		case reflect.Float32:
			v, err := ToFloat64(value)
			if err != nil {
				return fmt.Errorf("failed to set field %s due to type mismatch err %s", fieldName, err)
			}
			targetValue.SetFloat(float64(v))

		case reflect.Float64:
			v, err := ToFloat64(value)
			if err != nil {
				return fmt.Errorf("failed to set field %s due to type mismatch err %s", fieldName, err)
			}
			targetValue.SetFloat(v)
		case reflect.Struct:
			unMarshalInterface := reflect.TypeOf((*UnmarshalableOption)(nil)).Elem()
			if fieldValue.Addr().Type().Implements(unMarshalInterface) {
				unmarshaler := fieldValue.Addr().Interface().(UnmarshalableOption)
				if err := unmarshaler.UnmarshalOption(value); err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("unexpected field type: %s", elem.Type().Field(i).Type.Kind())
		}
	}
	return nil
}

// I.e. if the flag has the type not specified --> Can still set to int, int8, ... depending on the field.
// TODO: Pointer support
