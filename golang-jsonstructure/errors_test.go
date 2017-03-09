package jsonstructure

import "testing"

func TestErrorAt(t *testing.T) {
	if errorAt(nil, nil) != nil {
		t.Error("Expected nil result")
	}
}
