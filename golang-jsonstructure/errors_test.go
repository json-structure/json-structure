package jsonstructure

import "testing"

func TestEmptyErrors(t *testing.T) {
	if errorAt(nil, "") != nil {
		t.Error("Expected nil result")
	}
	if enumError(nil, "") != nil {
		t.Error("Expected nil result")
	}
}
