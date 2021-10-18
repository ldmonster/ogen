package json

import (
	"bytes"

	json "github.com/json-iterator/go"
)

// Marshal value to json.
func Marshal(val interface{}) ([]byte, error) {
	return json.Marshal(val)
}

// Unmarshal value from json.
func Unmarshal(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}

// Unmarshaler implements json reading.
type Unmarshaler interface {
	ReadJSON(i *json.Iterator) error
}

// Marshaler implements json writing.
type Marshaler interface {
	WriteJSON(s *json.Stream)
}

// Value represents a json value.
type Value interface {
	Marshaler
	Unmarshaler
}

// Settable value can be set (present) or unset
// (i.e. not provided or undefined).
type Settable interface {
	IsSet() bool
}

// Resettable value can be unset.
type Resettable interface {
	Reset()
}

// Nullable can be nil (but defined) or not.
type Nullable interface {
	IsNil() bool
}

// Encode Marshaler to byte slice.
func Encode(m Marshaler) []byte {
	buf := new(bytes.Buffer)
	s := NewStream(buf)
	m.WriteJSON(s)
	_ = s.Flush()
	return buf.Bytes()
}