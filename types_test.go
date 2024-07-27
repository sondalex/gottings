package gottings

import (
	"encoding/json"
	"testing"
)

func TestNullString(t *testing.T) {
	type Config struct {
		Key NullString `json:"key"`
	}
	t.Run("unmarshaling", func(t *testing.T) {
		config := &Config{}
		data := []byte(`{"key": "value"}`)
		err := json.Unmarshal(data, config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !config.Key.Valid || config.Key.String != "value" {
			t.Fatalf("Expected valid string with value 'value', got valid=%v, value=%v", config.Key.Valid, config.Key.String)
		}

		data = []byte(`{"key":null}`)
		err = json.Unmarshal(data, config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if config.Key.Valid {
			t.Fatalf("Expected invalid string, got valid=%v", config.Key.Valid)
		}
	})
	t.Run("marshaling", func(t *testing.T) {
		config := &Config{}
		expected := []byte(`{"key":null}`)

		raw, err := json.Marshal(config)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}

		expected = []byte(`{"key":"value"}`)
		config.Key = NewNullString("value")
		raw, err = json.Marshal(config)
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}
	})
}

func TestNullInt(t *testing.T) {
	type Config struct {
		Key NullInt `json:"key"`
	}

	t.Run("unmarshaling", func(t *testing.T) {
		config := &Config{}
		data := []byte(`{"key": 123}`)
		err := json.Unmarshal(data, config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !config.Key.Valid || config.Key.Int != 123 {
			t.Fatalf("Expected valid int with value 123, got valid=%v, value=%v", config.Key.Valid, config.Key.Int)
		}

		data = []byte(`{"key": null}`)
		err = json.Unmarshal(data, config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if config.Key.Valid {
			t.Fatalf("Expected invalid int, got valid=%v", config.Key.Valid)
		}
	})
	t.Run("marshaling", func(t *testing.T) {
		config := &Config{}

		expected := []byte(`{"key":null}`)

		raw, err := json.Marshal(config)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}

		expected = []byte(`{"key":123}`)
		config.Key = NewNullInt(123)
		raw, err = json.Marshal(config)
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}
	})
}

func TestNullInt8(t *testing.T) {
	type Config struct {
		Key NullInt8 `json:"key"`
	}

	t.Run("marshaling", func(t *testing.T) {
		config := &Config{}
		data := []byte(`{"key": 123}`)
		err := json.Unmarshal(data, config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !config.Key.Valid || config.Key.Int8 != 123 {
			t.Fatalf("Expected valid int8 with value 123, got valid=%v, value=%v", config.Key.Valid, config.Key.Int8)
		}

		data = []byte(`{"key": null}`)
		err = json.Unmarshal(data, &config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if config.Key.Valid {
			t.Fatalf("Expected invalid int8, got valid=%v", config.Key.Valid)
		}

	})
	t.Run("unmarshaling", func(t *testing.T) {
		config := &Config{}
		expected := []byte(`{"key":null}`)

		raw, err := json.Marshal(config)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}

		expected = []byte(`{"key":123}`)
		config.Key = NewNullInt8(123)
		raw, err = json.Marshal(config)
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}
	})
}

func TestNullInt16(t *testing.T) {
	type Config struct {
		Key NullInt16 `json:"key"`
	}

	t.Run("marshaling", func(t *testing.T) {
		config := &Config{}
		data := []byte(`{"key": 123}`)
		err := json.Unmarshal(data, config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !config.Key.Valid || config.Key.Int16 != 123 {
			t.Fatalf("Expected valid int16 with value 123, got valid=%v, value=%v", config.Key.Valid, config.Key.Int16)
		}

		data = []byte(`{"key": null}`)
		err = json.Unmarshal(data, &config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if config.Key.Valid {
			t.Fatalf("Expected invalid int16, got valid=%v", config.Key.Valid)
		}

	})
	t.Run("unmarshaling", func(t *testing.T) {
		config := &Config{}
		expected := []byte(`{"key":null}`)

		raw, err := json.Marshal(config)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}

		expected = []byte(`{"key":123}`)
		config.Key = NewNullInt16(123)
		raw, err = json.Marshal(config)
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}
	})
}

func TestNullInt32(t *testing.T) {
	type Config struct {
		Key NullInt32 `json:"key"`
	}

	t.Run("marshaling", func(t *testing.T) {
		config := &Config{}
		data := []byte(`{"key": 123}`)
		err := json.Unmarshal(data, config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !config.Key.Valid || config.Key.Int32 != 123 {
			t.Fatalf("Expected valid int32 with value 123, got valid=%v, value=%v", config.Key.Valid, config.Key.Int32)
		}

		data = []byte(`{"key": null}`)
		err = json.Unmarshal(data, &config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if config.Key.Valid {
			t.Fatalf("Expected invalid int32, got valid=%v", config.Key.Valid)
		}

	})
	t.Run("unmarshaling", func(t *testing.T) {
		config := &Config{}
		expected := []byte(`{"key":null}`)

		raw, err := json.Marshal(config)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}

		expected = []byte(`{"key":123}`)
		config.Key = NewNullInt32(123)
		raw, err = json.Marshal(config)
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}
	})
}

func TestNullInt64(t *testing.T) {
	type Config struct {
		Key NullInt64 `json:"key"`
	}

	t.Run("marshaling", func(t *testing.T) {
		config := &Config{}
		data := []byte(`{"key": 123}`)
		err := json.Unmarshal(data, config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !config.Key.Valid || config.Key.Int64 != 123 {
			t.Fatalf("Expected valid int64 with value 123, got valid=%v, value=%v", config.Key.Valid, config.Key.Int64)
		}

		data = []byte(`{"key": null}`)
		err = json.Unmarshal(data, &config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if config.Key.Valid {
			t.Fatalf("Expected invalid int64, got valid=%v", config.Key.Valid)
		}

	})
	t.Run("unmarshaling", func(t *testing.T) {
		config := &Config{}
		expected := []byte(`{"key":null}`)

		raw, err := json.Marshal(config)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}

		expected = []byte(`{"key":123}`)
		config.Key = NewNullInt64(123)
		raw, err = json.Marshal(config)
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}
	})
}

func TestNullFloat32(t *testing.T) {
	type Config struct {
		Key NullFloat32 `json:"key"`
	}
	t.Run("unmarshaling", func(t *testing.T) {
		config := &Config{}
		data := []byte(`{"key": 12.34}`)
		err := json.Unmarshal(data, config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !config.Key.Valid || config.Key.Float32 != 12.34 {
			t.Fatalf("Expected valid float32 with value 12.34, got valid=%v, value=%v", config.Key.Valid, config.Key.Float32)
		}

		data = []byte(`{"key": null}`)
		err = json.Unmarshal(data, &config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if config.Key.Valid {
			t.Fatalf("Expected invalid float32, got valid=%v", config.Key.Valid)
		}

	})
	t.Run("marshaling", func(t *testing.T) {
		config := &Config{}
		expected := []byte(`{"key":null}`)

		raw, err := json.Marshal(config)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}

		expected = []byte(`{"key":12.34}`)
		config.Key = NewNullFloat32(12.34)
		raw, err = json.Marshal(config)
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}
	})
}

func TestNullFloat64(t *testing.T) {
	type Config struct {
		Key NullFloat64 `json:"key"`
	}
	t.Run("unmarshaling", func(t *testing.T) {
		config := &Config{}
		data := []byte(`{"key": 12.34}`)
		err := json.Unmarshal(data, config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !config.Key.Valid || config.Key.Float64 != 12.34 {
			t.Fatalf("Expected valid float64 with value 12.34, got valid=%v, value=%v", config.Key.Valid, config.Key.Float64)
		}

		data = []byte(`{"key": null}`)
		err = json.Unmarshal(data, &config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if config.Key.Valid {
			t.Fatalf("Expected invalid float64, got valid=%v", config.Key.Valid)
		}

	})
	t.Run("marshaling", func(t *testing.T) {
		config := &Config{}
		expected := []byte(`{"key":null}`)

		raw, err := json.Marshal(config)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}

		expected = []byte(`{"key":12.34}`)
		config.Key = NewNullFloat64(12.34)
		raw, err = json.Marshal(config)
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}
	})
}

func TestNullBool(t *testing.T) {
	type Config struct {
		Key NullBool `json:"key"`
	}
	t.Run("unmarshaling", func(t *testing.T) {
		config := &Config{}
		data := []byte(`{"key": true}`)
		err := json.Unmarshal(data, config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !config.Key.Valid || config.Key.Bool != true {
			t.Fatalf("Expected valid bool with value true, got valid=%v, value=%v", config.Key.Valid, config.Key.Bool)
		}

		data = []byte(`{"key": null}`)
		err = json.Unmarshal(data, &config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if config.Key.Valid {
			t.Fatalf("Expected invalid bool, got valid=%v", config.Key.Valid)
		}

	})
	t.Run("marshaling", func(t *testing.T) {
		config := &Config{}
		expected := []byte(`{"key":null}`)

		raw, err := json.Marshal(config)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}

		expected = []byte(`{"key":true}`)
		config.Key = NewNullBool(true)
		raw, err = json.Marshal(config)
		if string(raw) != string(expected) {
			t.Fatalf("Expected data %v got %v", string(expected), string(raw))
		}
	})
}

func TestNullBoolValue(t *testing.T) {
	tests := []struct {
		name string
		nb   NullBool
		want bool
	}{
		{"true", NullBool{Bool: true}, true},
		{"false", NullBool{Bool: false}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.nb.Value(); got != tt.want {
				t.Errorf("NullBool.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullStringValue(t *testing.T) {
	tests := []struct {
		name string
		ns   NullString
		want string
	}{
		{"hello", NullString{String: "hello"}, "hello"},
		{"empty", NullString{String: ""}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ns.Value(); got != tt.want {
				t.Errorf("NullString.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullIntValue(t *testing.T) {
	tests := []struct {
		name string
		ni   NullInt
		want int
	}{
		{"42", NullInt{Int: 42}, 42},
		{"0", NullInt{Int: 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ni.Value(); got != tt.want {
				t.Errorf("NullInt.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullInt8Value(t *testing.T) {
	tests := []struct {
		name string
		ni8  NullInt8
		want int8
	}{
		{"42", NullInt8{Int8: 42}, 42},
		{"0", NullInt8{Int8: 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ni8.Value(); got != tt.want {
				t.Errorf("NullInt8.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullInt16Value(t *testing.T) {
	tests := []struct {
		name string
		ni16 NullInt16
		want int16
	}{
		{"42", NullInt16{Int16: 42}, 42},
		{"0", NullInt16{Int16: 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ni16.Value(); got != tt.want {
				t.Errorf("NullInt16.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullInt32Value(t *testing.T) {
	tests := []struct {
		name string
		ni32 NullInt32
		want int32
	}{
		{"42", NullInt32{Int32: 42}, 42},
		{"0", NullInt32{Int32: 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ni32.Value(); got != tt.want {
				t.Errorf("NullInt32.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullInt64Value(t *testing.T) {
	tests := []struct {
		name string
		ni64 NullInt64
		want int64
	}{
		{"42", NullInt64{Int64: 42}, 42},
		{"0", NullInt64{Int64: 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ni64.Value(); got != tt.want {
				t.Errorf("NullInt64.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullFloat32Value(t *testing.T) {
	tests := []struct {
		name string
		nf32 NullFloat32
		want float32
	}{
		{"3.14", NullFloat32{Float32: 3.14}, 3.14},
		{"0", NullFloat32{Float32: 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.nf32.Value(); got != tt.want {
				t.Errorf("NullFloat32.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullFloat64Value(t *testing.T) {
	tests := []struct {
		name string
		nf64 NullFloat64
		want float64
	}{
		{"3.14", NullFloat64{Float64: 3.14}, 3.14},
		{"0", NullFloat64{Float64: 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.nf64.Value(); got != tt.want {
				t.Errorf("NullFloat64.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
