package opennative

import (
	"encoding/json"
	"errors"
)

// ErrExtensionUnmarshal in case of Extension is not initialised
var ErrExtensionUnmarshal = errors.New("opennative.Extension: UnmarshalJSON on nil pointer")

// Extension is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler, defined in encoding/json package,
// works similarly to Extension,
// but does not need to be used as pointer for proper JSON marshaling.
type Extension []byte

// MarshalJSON returns e as the JSON encoding of e.
func (e Extension) MarshalJSON() ([]byte, error) {
	return e, nil
}

// UnmarshalJSON sets *e to a copy of data.
func (e *Extension) UnmarshalJSON(data []byte) error {
	if e == nil {
		return ErrExtensionUnmarshal
	}
	*e = append((*e)[0:0], data...)
	return nil
}

var (
	_ json.Marshaler   = (Extension)(nil)
	_ json.Unmarshaler = (*Extension)(nil)
)
