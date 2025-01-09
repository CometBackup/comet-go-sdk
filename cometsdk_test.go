package cometsdk

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestUnknownFields tests that unknown fields are preserved in JSON roundtrips
func TestUnknownFields(t *testing.T) {

	// Setup

	example := `{"NewUnknownField": {"foo": "NestedContentValue"}, "RandomDelaySecs": 128}`

	var g GlobalOverrideOptions
	err := json.Unmarshal([]byte(example), &g)
	if err != nil {
		t.Fatal(err)
	}

	// Tests

	t.Run("Known fields", func(t *testing.T) {
		if got, want := g.RandomDelaySecs, uint64(128); got != want {
			t.Errorf("RandomDelaySecs: got %d, want %d", got, want)
		}
	})

	t.Run("Unknown fields", func(t *testing.T) {
		dec, err := json.Marshal(g)
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(dec), "NestedContentValue") {
			t.Errorf("Missing expected value 'NestedContentValue' in result")
		}
	})

}
