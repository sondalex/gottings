package gottings

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type NullString struct {
	String string
	Valid  bool
}

type NullBool struct {
	Bool  bool
	Valid bool
}

type NullFloat32 struct {
	Float32 float32
	Valid   bool
}

type NullFloat64 struct {
	Float64 float64
	Valid   bool
}

type NullInt struct {
	Int   int
	Valid bool
}

type NullInt8 struct {
	Int8  int8
	Valid bool
}

type NullInt16 struct {
	Int16 int16
	Valid bool
}

type NullInt32 struct {
	Int32 int32
	Valid bool
}

type NullInt64 struct {
	Int64 int64
	Valid bool
}

func NewNullString(s string) NullString {
	return NullString{
		Valid:  true,
		String: s,
	}
}

func NewNullInt(i int) NullInt {
	return NullInt{
		Valid: true,
		Int:   i,
	}
}

func NewNullInt8(i int8) NullInt8 {
	return NullInt8{
		Valid: true,
		Int8:  i,
	}
}

func NewNullInt16(i int16) NullInt16 {
	return NullInt16{
		Valid: true,
		Int16: i,
	}
}

func NewNullInt32(i int32) NullInt32 {
	return NullInt32{
		Valid: true,
		Int32: i,
	}
}

func NewNullInt64(i int64) NullInt64 {
	return NullInt64{
		Valid: true,
		Int64: i,
	}
}

func NewNullFloat32(d float32) NullFloat32 {
	return NullFloat32{
		Valid:   true,
		Float32: d,
	}
}

func NewNullFloat64(d float64) NullFloat64 {
	return NullFloat64{
		Valid:   true,
		Float64: d,
	}
}

func NewNullBool(b bool) NullBool {
	return NullBool{
		Valid: true,
		Bool:  b,
	}
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s NullInt) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Int)
}

func (s NullInt8) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Int8)
}

func (s NullInt16) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Int16)
}

func (s NullInt32) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Int32)
}

func (s NullInt64) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Int64)
}

func (s NullFloat32) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Float32)
}

func (s NullFloat64) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Float64)
}

func (s NullBool) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Bool)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		s.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &s.String); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

func (s *NullString) UnmarshalEnvironmentValue(data []byte) error {
	s.Valid = true
	s.String = string(data)
	return nil
}

func (s *NullString) UnmarshalOption(data string) error {
	s.Valid = true
	s.String = data
	return nil
}

func (s *NullInt) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		s.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &s.Int); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

func (s *NullInt) UnmarshalEnvironmentValue(data []byte) error {
	value, err := strconv.ParseInt(string(data), 10, 0)
	if err != nil {
		return err
	}
	s.Valid = true
	s.Int = int(value)
	return nil
}

func (s *NullInt) UnmarshalOption(data int) error {
	s.Valid = true
	s.Int = data
	return nil
}

func (s *NullInt8) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		s.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &s.Int8); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

func (s *NullInt8) UnmarshalEnvironmentValue(data []byte) error {
	value, err := strconv.ParseInt(string(data), 10, 8)
	if err != nil {
		return err
	}
	s.Valid = true
	s.Int8 = int8(value)
	return nil
}

func (s *NullInt8) UnmarshalOption(data int8) error {
	s.Valid = true
	s.Int8 = data
	return nil
}

func (s *NullInt16) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		s.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &s.Int16); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

func (s *NullInt16) UnmarshalEnvironmentValue(data []byte) error {
	value, err := strconv.ParseInt(string(data), 10, 16)
	if err != nil {
		return err
	}
	s.Valid = true
	s.Int16 = int16(value)
	return nil
}
func (s *NullInt16) UnmarshalOption(data int16) error {
	s.Valid = true
	s.Int16 = data
	return nil
}

func (s *NullInt32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		s.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &s.Int32); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

func (s *NullInt32) UnmarshalEnvironmentValue(data []byte) error {
	value, err := strconv.ParseInt(string(data), 10, 32)
	if err != nil {
		return err
	}
	s.Valid = true
	s.Int32 = int32(value)
	return nil
}
func (s *NullInt32) UnmarshalOption(data int32) error {
	s.Valid = true
	s.Int32 = data
	return nil
}

func (s *NullInt64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		s.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &s.Int64); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

func (s *NullInt64) UnmarshalEnvironmentValue(data []byte) error {
	value, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	s.Valid = true
	s.Int64 = value
	return nil
}

func (s *NullInt64) UnmarshalOption(data int64) error {
	s.Valid = true
	s.Int64 = data
	return nil
}

func (s *NullFloat32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		s.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &s.Float32); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

func (s *NullFloat32) UnmarshalEnvironmentValue(data []byte) error {
	value, err := strconv.ParseFloat(string(data), 32)
	if err != nil {
		return err
	}
	s.Valid = true
	s.Float32 = float32(value)
	return nil
}

func (s *NullFloat32) UnmarshalOption(data float32) error {
	s.Valid = true
	s.Float32 = data
	return nil
}

func (s *NullFloat64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		s.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &s.Float64); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

func (s *NullFloat64) UnmarshalEnvironmentValue(data []byte) error {
	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}
	s.Valid = true
	s.Float64 = value
	return nil
}

func (s *NullFloat64) UnmarshalOption(data float64) error {
	s.Valid = true
	s.Float64 = data
	return nil
}

func (s *NullBool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		s.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &s.Bool); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

func (s *NullBool) UnmarshalEnvironmentValue(data []byte) error {
	value, err := strconv.ParseBool(string(data))
	if err != nil {
		return err
	}
	s.Valid = true
	s.Bool = value
	return nil
}

func (s *NullBool) UnmarshalOption(data bool) error {
	s.Valid = true
	s.Bool = data
	return nil
}
