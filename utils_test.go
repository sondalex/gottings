package gottings

import (
	"os"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	type Config struct {
		Host string `json:"host" env:"TEST_HOST"`
		Port int    `json:"port" env:"TEST_PORT"`
	}
	t.Run("load from file - all keys in json", func(t *testing.T) {
		t.Setenv("TEST_PORT", "8080")
		config := Config{}
		filePath := "testdata/info.json"
		data, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		err = LoadConfiguration(data, &config)
		if err != nil {
			t.Errorf("expected error to be nil, got %s", err)
		}
		if config.Port != 8080 {
			t.Errorf("Expected 8080 got config.Port=%d\n", config.Port)
		}
		if config.Host != "127.0.0.1" {
			t.Errorf("Expected 127.0.0.1 got config.Host=%s\n", config.Host)
		}
	})
	t.Run("not all keys in json", func(t *testing.T) {
		t.Setenv("TEST_PORT", "8080")
		data := []byte(`{"host": "127.0.0.1"}`)
		config := Config{}
		err := LoadConfiguration(data, &config)
		if err != nil {
			t.Fatalf("expected error to be nil, got %s", err)
		}
		if config.Port != 8080 {
			t.Errorf("Expected 8080 got config.Port=%d\n", config.Port)
		}
		if config.Host != "127.0.0.1" {
			t.Errorf("Expected 127.0.0.1 got config.Host=%s\n", config.Host)
		}
	})
	t.Run("keys in json", func(t *testing.T) {
		t.Setenv("TEST_PORT", "8080")
		t.Setenv("TEST_HOST", "127.0.0.1")
		data := []byte(``)
		config := Config{}
		err := LoadConfiguration(data, &config)
		if err != nil {
			t.Fatalf("expected error to be nil, got %s", err)
		}
		if config.Port != 8080 {
			t.Errorf("Expected 8080 got config.Port=%d\n", config.Port)
		}
		if config.Host != "127.0.0.1" {
			t.Errorf("Expected 127.0.0.1 got config.Host=%s\n", config.Host)
		}
	})
}
func TestIsInteger(t *testing.T) {
	var i32 int32 = int32(5)
	var f64 float64 = float64(3.14)
	testCases := []struct {
		input    any
		expected bool
	}{
		{input: 5, expected: true},              // int
		{input: int8(5), expected: true},        // int8
		{input: int16(5), expected: true},       // int16
		{input: i32, expected: true},            // int32
		{input: &i32, expected: true},           // int32
		{input: int64(5), expected: true},       // int64
		{input: "string", expected: false},      // string
		{input: f64, expected: false},           // float64
		{input: &f64, expected: false},          // float64
		{input: false, expected: false},         // bool
		{input: nil, expected: false},           // nil
		{input: uint(5), expected: false},       // uint
		{input: complex(1, 1), expected: false}, // complex
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := IsInteger(tc.input)
			if result != tc.expected {
				t.Errorf("IsInteger(%v) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestIsFloat(t *testing.T) {
	var f32 float32 = float32(3.14)
	var i8 int8 = int8(5)
	testCases := []struct {
		input    any
		expected bool
	}{
		{input: 3.14, expected: true},      // float64
		{input: f32, expected: true},       // float32
		{input: &f32, expected: true},      // int8
		{input: 5, expected: false},        // int
		{input: i8, expected: false},       // int8
		{input: &i8, expected: false},      // *int8
		{input: "string", expected: false}, // string
		{input: nil, expected: false},      // nil
		{input: false, expected: false},    // bool
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := IsFloat(tc.input)
			if result != tc.expected {
				t.Errorf("IsFloat(%v) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	var i32 int32 = int32(5)
	var f64 float64 = float64(3.14)
	testCases := []struct {
		input    any
		expected int64
		hasError bool
	}{
		{input: 5, expected: 5, hasError: false},
		{input: int8(5), expected: 5, hasError: false},
		{input: int16(5), expected: 5, hasError: false},
		{input: i32, expected: 5, hasError: false},
		{input: &i32, expected: 5, hasError: false},
		{input: int64(5), expected: 5, hasError: false},
		{input: f64, expected: 0, hasError: true},
		{input: &f64, expected: 0, hasError: true},
		{input: "string", expected: 0, hasError: true},
		{input: nil, expected: 0, hasError: true},
		{input: uint(5), expected: 0, hasError: true},
		{input: complex(1, 1), expected: 0, hasError: true},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result, err := ToInt64(tc.input)
			if (err != nil) != tc.hasError {
				t.Errorf("ToInt64(%v) error = %v, expected error = %v", tc.input, err, tc.hasError)
			}
			if result != tc.expected {
				t.Errorf("ToInt64(%v) = %d; expected %d", tc.input, result, tc.expected)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	var i int = 5
	var f64 float64 = float64(3.14)
	testCases := []struct {
		input    any
		expected float64
		hasError bool
	}{
		{input: float32(3.14), expected: float64(float32(3.14)), hasError: false},
		{input: f64, expected: f64, hasError: false},
		{input: &f64, expected: f64, hasError: false},
		{input: i, expected: 0, hasError: true},
		{input: &i, expected: 0, hasError: true},
		{input: "string", expected: 0, hasError: true},
		{input: nil, expected: 0, hasError: true},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result, err := ToFloat64(tc.input)
			if (err != nil) != tc.hasError {
				t.Errorf("ToFloat64(%v) error = %v, expected error = %v", tc.input, err, tc.hasError)
			}
			if result != tc.expected {
				t.Errorf("ToFloat64(%v) = %f; expected %f", tc.input, result, tc.expected)
			}
		})
	}
}
