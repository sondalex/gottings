package gottings

import (
	"reflect"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	type Config struct {
		Int         int         `env:"TEST_INT"`
		NullInt     NullInt     `env:"TEST_NULLINT"`
		Int8        int8        `env:"TEST_INT8"`
		NullInt8    NullInt8    `env:"TEST_NULLINT8"`
		Int16       int16       `env:"TEST_INT16"`
		NullInt16   NullInt16   `env:"TEST_NULLINT16"`
		Int32       int32       `env:"TEST_INT32"`
		NullInt32   NullInt32   `env:"TEST_NULLINT32"`
		Int64       int64       `env:"TEST_INT64"`
		NullInt64   NullInt64   `env:"TEST_NULLINT64"`
		Float32     float32     `env:"TEST_FLOAT32"`
		NullFloat32 NullFloat32 `env:"TEST_NULLFLOAT32"`
		Float64     float64     `env:"TEST_FLOAT64"`
		NullFloat64 NullFloat64 `env:"TEST_NULLFLOAT64"`
		Bool        bool        `env:"TEST_BOOL"`
		NullBool    NullBool    `env:"TEST_NULLBOOL"`
		String      string      `env:"TEST_STRING"`
		NullString  NullString  `env:"TEST_NULLSTRING"`
	}
	t.Setenv("TEST_INT", "1")
	t.Setenv("TEST_NULLINT", "1")
	t.Setenv("TEST_INT8", "8")
	t.Setenv("TEST_NULLINT8", "8")
	t.Setenv("TEST_INT16", "16")
	t.Setenv("TEST_NULLINT16", "16")
	t.Setenv("TEST_INT32", "32")
	t.Setenv("TEST_NULLINT32", "32")
	t.Setenv("TEST_INT64", "64")
	t.Setenv("TEST_NULLINT64", "64")
	t.Setenv("TEST_BOOL", "true")
	t.Setenv("TEST_NULLBOOL", "true")
	t.Setenv("TEST_STRING", "string")
	t.Setenv("TEST_NULLSTRING", "string")
	t.Setenv("TEST_FLOAT32", "3.2")
	t.Setenv("TEST_NULLFLOAT32", "3.2")
	t.Setenv("TEST_NULLFLOAT64", "6.4")
	t.Setenv("TEST_FLOAT64", "6.4")
	config := Config{}
	expected := Config{
		Int:         1,
		NullInt:     NullInt{Int: 1, Valid: true},
		Int8:        8,
		NullInt8:    NullInt8{Int8: 8, Valid: true},
		Int16:       16,
		NullInt16:   NullInt16{Int16: 16, Valid: true},
		Int32:       32,
		NullInt32:   NullInt32{Int32: 32, Valid: true},
		Int64:       64,
		NullInt64:   NullInt64{Int64: 64, Valid: true},
		Float32:     3.2,
		NullFloat32: NullFloat32{Float32: 3.2, Valid: true},
		Float64:     6.4,
		NullFloat64: NullFloat64{Float64: 6.4, Valid: true},
		Bool:        true,
		NullBool:    NullBool{Bool: true, Valid: true},
		String:      "string",
		NullString:  NullString{String: "string", Valid: true},
	}

	t.Run("validation - struct", func(t *testing.T) {
		err := LoadEnv(config)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

	})
	t.Run("validation - pointer to struct", func(t *testing.T) {
		err := LoadEnv(&config)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if config != expected {
			t.Errorf("result %v does not match expected %v\n", config, expected)
		}
	})
	t.Run("validation - pointer to struct - pointer fields", func(t *testing.T) {
		type Config2 struct {
			Int         *int         `env:"TEST_INT"`
			NullInt     *NullInt     `env:"TEST_NULLINT"`
			Int8        *int8        `env:"TEST_INT8"`
			NullInt8    *NullInt8    `env:"TEST_NULLINT8"`
			Int16       *int16       `env:"TEST_INT16"`
			NullInt16   *NullInt16   `env:"TEST_NULLINT16"`
			Int32       *int32       `env:"TEST_INT32"`
			NullInt32   *NullInt32   `env:"TEST_NULLINT32"`
			Int64       *int64       `env:"TEST_INT64"`
			NullInt64   *NullInt64   `env:"TEST_NULLINT64"`
			Float32     *float32     `env:"TEST_FLOAT32"`
			NullFloat32 *NullFloat32 `env:"TEST_NULLFLOAT32"`
			Float64     *float64     `env:"TEST_FLOAT64"`
			NullFloat64 *NullFloat64 `env:"TEST_NULLFLOAT64"`
			Bool        *bool        `env:"TEST_BOOL"`
			NullBool    *NullBool    `env:"TEST_NULLBOOL"`
			String      *string      `env:"TEST_STRING"`
			NullString  *NullString  `env:"TEST_NULLSTRING"`
		}
		config2 := Config2{}
		var i int = 1
		var i8 int8 = 8
		var i16 int16 = 16
		var i32 int32 = 32
		var i64 int64 = 64
		var Bool bool = true
		var f64 float64 = 6.4
		var f32 float32 = 3.2
		var String string = "string"
		expected2 := Config2{
			Int:         &i,
			NullInt:     &NullInt{Int: 1, Valid: true},
			Int8:        &i8,
			NullInt8:    &NullInt8{Int8: 8, Valid: true},
			Int16:       &i16,
			NullInt16:   &NullInt16{Int16: 16, Valid: true},
			Int32:       &i32,
			NullInt32:   &NullInt32{Int32: 32, Valid: true},
			Int64:       &i64,
			NullInt64:   &NullInt64{Int64: 64, Valid: true},
			Float32:     &f32,
			NullFloat32: &NullFloat32{Float32: 3.2, Valid: true},
			Float64:     &f64,
			NullFloat64: &NullFloat64{Float64: 6.4, Valid: true},
			Bool:        &Bool,
			NullBool:    &NullBool{Bool: true, Valid: true},
			String:      &String,
			NullString:  &NullString{String: "string", Valid: true},
		}
		err := LoadEnv(&config2)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if reflect.DeepEqual(config2, expected2) {
			t.Errorf("result %v does not match expected %v\n", config2, expected2)
		}
	})
}
