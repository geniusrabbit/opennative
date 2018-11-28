package opennative

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// NumberOrString attempts to fix OpenRTB incompatibilities
// of exchanges. On decoding, it can handle numbers and strings.
// On encoding, it will generate a number, as intended by the
// standard.
type NumberOrString int

// UnmarshalJSON implements json.Unmarshaler
func (n *NumberOrString) UnmarshalJSON(data []byte) (err error) {
	var v int

	if data = bytes.Trim(data, "\""); len(data) > 0 {
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
	}

	*n = NumberOrString(v)
	return nil
}

// StringOrNumber attempts to fix OpenRTB incompatibilities
// of exchanges. On decoding, it can handle numbers and strings.
// On encoding, it will generate a string, as intended by the
// standard.
type StringOrNumber string

// UnmarshalJSON implements json.Unmarshaler
func (n *StringOrNumber) UnmarshalJSON(data []byte) (err error) {
	var v string
	switch {
	case len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"':
		err = json.Unmarshal(data, &v)
	case bytes.ContainsRune(data, '.'):
		if _, err = strconv.ParseFloat(string(data), 64); err == nil {
			v = string(data)
		}
	default:
		if _, err = strconv.ParseInt(string(data), 10, 64); err == nil {
			v = string(data)
		}
	}
	if err == nil {
		*n = StringOrNumber(v)
	}
	return
}
