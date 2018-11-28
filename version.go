package opennative

import (
	"encoding/json"
)

// Version type of field
type Version string

// List of versions
var (
	Version1_0 Version = "1.0"
	Version1_1 Version = "1.1"
	Version1_2 Version = "1.2"
)

func (n Version) String() string {
	return string(n)
}

// IsVer1_0 of protocol
func (n Version) IsVer1_0() bool {
	return n == Version1_0 || n == Version("1")
}

// IsVer1_1 of protocol
func (n Version) IsVer1_1() bool {
	return n == Version1_1
}

// IsVer1_2 of protocol
func (n Version) IsVer1_2() bool {
	return n == Version1_2
}

// UnmarshalJSON implements json.Unmarshaler
func (n *Version) UnmarshalJSON(data []byte) (err error) {
	var v string
	switch {
	case len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"':
		err = json.Unmarshal(data, &v)
	default:
		v = string(data)
	}
	if err == nil {
		*n = Version(v)
	}
	return
}

var _ json.Unmarshaler = (*Version)(nil)
