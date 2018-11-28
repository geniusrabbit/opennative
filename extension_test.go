package opennative

import (
	"bytes"
	"encoding/json"
	"testing"
)

func Test_Extension(t *testing.T) {
	t.Run("marshal_JSON", func(t *testing.T) {
		data := Extension(`{"key":"value"}`)
		b, err := json.Marshal(data)

		if err != nil {
			t.Errorf("Marshaling error: %s", err.Error())
		}

		if !bytes.Equal(data, b) {
			t.Errorf("Marshaling result must be equal to: %s", string(data))
		}
	})

	t.Run("unmarshal_JSON", func(t *testing.T) {
		source := []byte(`{"key":"value"}`)
		data := Extension(``)
		err := json.Unmarshal(source, &data)

		if err != nil {
			t.Errorf("Unmarshaling error: %s", err.Error())
		}

		if !bytes.Equal(data, source) {
			t.Errorf("Unmarshaling result must be equal to: %s", string(source))
		}
	})
}
