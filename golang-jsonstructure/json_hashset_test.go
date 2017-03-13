package jsonstructure

import (
	"encoding/json"
	"testing"
)

func TestHashSetString(t *testing.T) {
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
}

func TestHashSetNumber(t *testing.T) {
	s := createSet()
	dec := json.Number("50.0")
	res, err := s.PutIfAbsent(dec)
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

func TestHashSetMap(t *testing.T) {
	m1 := make(map[string]interface{})
	m1["foo"] = json.Number("50.0")
	s := createSet()
	res, err := s.PutIfAbsent(m1)
	if err != nil {
		t.Errorf("Insertion error %v", err)
	}
	if !res {
		t.Error("Insert failed")
	}
	m2 := make(map[string]interface{})
	m2["foo"] = json.Number("5e1")
	res, err = s.PutIfAbsent(m2)
	if err != nil {
		t.Errorf("Insertion error %v", err)
	}
	if res {
		t.Error("Duplicate insertion was successful")
	}
}

func TestHashSetSlice(t *testing.T) {
	m1 := []interface{}{json.Number("50.0")}
	s := createSet()
	res, err := s.PutIfAbsent(m1)
	if err != nil {
		t.Errorf("Insertion error %v", err)
	}
	if !res {
		t.Error("Insert failed")
	}
	m2 := []interface{}{json.Number("5e1")}
	res, err = s.PutIfAbsent(m2)
	if err != nil {
		t.Errorf("Insertion error %v", err)
	}
	if res {
		t.Error("Duplicate insertion was successful")
	}
}
