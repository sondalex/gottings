package gottings

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadOptions(t *testing.T) {

	type Config struct {
		Int         int
		NullInt     NullInt
		Int8        int8
		NullInt8    NullInt8
		Int16       int16
		NullInt16   NullInt16
		Int32       int32
		NullInt32   NullInt32
		Int64       int64
		NullInt64   NullInt64
		Float32     float32
		NullFloat32 NullFloat32
		Float64     float64
		NullFloat64 NullFloat64
		Bool        bool
		NullBool    NullBool
		String      string
		NullString  NullString
	}
	type Config2 struct {
		Int         *int
		NullInt     *NullInt
		Int8        *int8
		NullInt8    *NullInt8
		Int16       *int16
		NullInt16   *NullInt16
		Int32       *int32
		NullInt32   *NullInt32
		Int64       *int64
		NullInt64   *NullInt64
		Float32     *float32
		NullFloat32 *NullFloat32
		Float64     *float64
		NullFloat64 *NullFloat64
		Bool        *bool
		NullBool    *NullBool
		String      *string
		NullString  *NullString
	}
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

	config := Config{}
	t.Run("explicit type in map", func(t *testing.T) {
		err := LoadOptions(map[string]interface{}{
			"Int":         1,
			"NullInt":     1,
			"Int8":        int8(8),
			"NullInt8":    int8(8),
			"Int16":       int16(16),
			"NullInt16":   int16(16),
			"Int32":       int32(32),
			"NullInt32":   int32(32),
			"Int64":       int64(64),
			"NullInt64":   int64(64),
			"Float32":     float32(3.2),
			"NullFloat32": float32(3.2),
			"Float64":     float64(6.4),
			"NullFloat64": float64(6.4),
			"Bool":        true,
			"NullBool":    true,
			"String":      "string",
			"NullString":  "string",
		}, &config)
		if err != nil {
			t.Fatalf("error not expected received: %s", err)
		}
		if reflect.DeepEqual(config, expected) {
			t.Fatalf("result %v and expected %v not equal", config, expected)
		}
	},
	)
	t.Run("non explicit type in map", func(t *testing.T) {
		err := LoadOptions(map[string]interface{}{
			"Int":         1,
			"NullInt":     1,
			"Int8":        8,
			"NullInt8":    8,
			"Int16":       16,
			"NullInt16":   16,
			"Int32":       32,
			"NullInt32":   32,
			"Int64":       64,
			"NullInt64":   64,
			"Float32":     3.2,
			"NullFloat32": 3.2,
			"Float64":     6.4,
			"NullFloat64": 6.4,
			"Bool":        true,
			"NullBool":    true,
			"String":      "string",
			"NullString":  "string",
		}, &config)
		if err != nil {
			t.Fatalf("error not expected received: %s", err)
		}
		if reflect.DeepEqual(config, expected) {
			t.Fatalf("result %v and expected %v not equal", config, expected)
		}
	})
	config2 := Config2{}
	t.Run("explicit type in map - pointers", func(t *testing.T) {
		err := LoadOptions(map[string]interface{}{
			"Int":         &i,
			"NullInt":     &i,
			"Int8":        &i8,
			"NullInt8":    &i8,
			"Int16":       &i16,
			"NullInt16":   &i16,
			"Int32":       &i32,
			"NullInt32":   &i32,
			"Int64":       &i64,
			"NullInt64":   &i64,
			"Float32":     &f32,
			"NullFloat32": &f32,
			"Float64":     &f64,
			"NullFloat64": &f64,
			"Bool":        &Bool,
			"NullBool":    &Bool,
			"String":      &String,
			"NullString":  &String,
		}, &config)
		if err != nil {
			t.Fatalf("error not expected received: %s", err)
		}
		if reflect.DeepEqual(config2, expected2) {
			t.Fatalf("result %v and expected %v not equal", config2, expected2)
		}
	})
}

func TestType(t *testing.T) {
	var a int = 2
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
}
