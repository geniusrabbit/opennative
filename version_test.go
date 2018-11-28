package opennative

import (
	"encoding/json"
	"testing"
)

func TestNativeVersion_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		data   string
		result string
	}{
		// TODO: Add test cases.
		{"base", `{"ver": 1}`, "1"},
		{"base string", `{"ver": "v1.2"}`, "v1.2"},
		{"raw number", `{"ver": 1.2}`, "1.2"},
	}

	for _, tt := range tests {
		var n struct {
			Ver Version
		}

		if err := json.Unmarshal([]byte(tt.data), &n); err != nil {
			t.Errorf("%q. NativeVersion.UnmarshalJSON() error = %v", tt.name, err)
		}

		if string(n.Ver) != tt.result {
			t.Errorf("Wrong native version: %q", n.Ver)
		}
	}
}
