package gottings

import (
	"encoding/json"
	"errors"
)

func LoadConfiguration(data []byte, v any) error {
	var err error
	if len(data) > 0 {
		err = json.Unmarshal(data, v)
		if err != nil {
			return err
		}
	}
	err = LoadEnv(v)
	if err != nil {
		return err
	}
	return nil
}

func IsInteger(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, *int, *int8, *int16, *int32, *int64:
		return true
	default:
		return false
	}
}

func IsFloat(v any) bool {
	switch v.(type) {
	case float32, float64, *float32, *float64:
		return true
	default:
		return false
	}
}

func ToInt64(v any) (int64, error) {
	isInt := IsInteger(v)
	if !isInt {
		return 0, errors.New("expected v to satisfy IsInteger(v) == true")
	}

	switch n := v.(type) {
	case int:
		return int64(n), nil
	case int8:
		return int64(n), nil
	case int16:
		return int64(n), nil
	case int32:
		return int64(n), nil
	case int64:
		return n, nil
	case *int:
		return int64(*n), nil
	case *int8:
		return int64(*n), nil
	case *int16:
		return int64(*n), nil
	case *int32:
		return int64(*n), nil
	case *int64:
		return int64(*n), nil
	default:
		// This should never be reached due to the IsInteger check
		return 0, errors.New("unexpected type encountered")
	}
}

func ToFloat64(v any) (float64, error) {
	isF := IsFloat(v)
	if !isF {
		return 0, errors.New("expected v to satisfy isFloat(v) == true")
	}

	switch n := v.(type) {
	case float32:
		return float64(n), nil
	case float64:
		return float64(n), nil
	case *float32:
		return float64(*n), nil
	case *float64:
		return float64(*n), nil
	default:
		// This should never be reached
		return 0., errors.New("unexpected type encountered")
	}
}
