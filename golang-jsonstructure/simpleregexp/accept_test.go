package simpleregexp

import "testing"

func TestAcceptSuccess(t *testing.T) {
	input := "^(foo)|(bar)[^a-z][A-Za-z].+.*.?.+?.*?.??.{1}.{1,3}.{1,}$"
	err := Accept(input)
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

func TestAcceptFailure(t *testing.T) {
	input := "\\d"
	err := Accept(input)
	if err == nil {
		t.Error("Expected Error")
	}
	t.Log(err)
}
