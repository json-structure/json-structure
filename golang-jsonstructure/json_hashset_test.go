package jsonstructure

import (
	"encoding/json"
	"testing"
)

func TestHashSet(t *testing.T) {
	var dec json.Number
	s := createSet()
	res, err := s.Has("foo")
	if err != nil {
		t.Errorf("Get error %v", err)
	}
	if res {
		t.Error("Get failed")
	}
	res, err = s.PutIfAbsent("foo")
	if err != nil {
		t.Errorf("Insertion error %v", err)
	}
	if !res {
		t.Error("Insert failed")
	}
	res, err = s.Has("foo")
	if err != nil {
		t.Errorf("Get error %v", err)
	}
	if !res {
		t.Error("Get failed")
	}
	res, err = s.PutIfAbsent("foo")
	if err != nil {
		t.Errorf("Insertion error %v", err)
	}
	if res {
		t.Error("Insert did not fail")
	}
	dec = json.Number("50.0")
	if err != nil {
		t.Errorf("Construction error %v", err)
	}
	res, err = s.PutIfAbsent(dec)
	if err != nil {
		t.Errorf("Insertion error %v", err)
	}
	if !res {
		t.Error("Insert failed")
	}
	dec = json.Number("5e1")
	if err != nil {
		t.Errorf("Construction error %v", err)
	}
	res, err = s.PutIfAbsent(dec)
	if err != nil {
		t.Errorf("Insertion error %v", err)
	}
	if res {
		t.Error("Duplicate insertion was successful")
	}
}
