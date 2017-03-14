package jsonstructure

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestElementOf(t *testing.T) {
	if !elementOf([]string{"a", "b", "c"}, "c") {
		t.Error("Could not find element")
	}
	if elementOf([]string{"a", "b", "c"}, "d") {
		t.Error("Found unexpected element")
	}
}

func TestMapMerge(t *testing.T) {
	src := map[string]interface{}{
		"foo": "hello",
		"bar": "alice",
	}
	dst := map[string]interface{}{
		"bar": "bob",
		"baz": "world",
	}
	expected := map[string]interface{}{
		"foo": "hello",
		"bar": "alice",
		"baz": "world",
	}
	err := mapMerge(dst, src, nil)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if !reflect.DeepEqual(dst, expected) {
		t.Error("Merge failure")
	}
}

func TestMapMergeRecursive(t *testing.T) {
	src := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": "baz",
		},
	}
	dst := map[string]interface{}{
		"foo": 3,
	}
	err := mapMerge(dst, src, nil)
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	dst = map[string]interface{}{}
	err = mapMerge(dst, src, nil)
	if err != nil {
		t.Error("Unexpected error")
	}
	if !reflect.DeepEqual(src, dst) {
		t.Error("Merge failure")
	}
	dst = map[string]interface{}{
		"foo": map[string]interface{}{
			"hello": "world",
		},
	}
	err = mapMerge(dst, src, nil)
	if err != nil {
		t.Error("Unexpected error")
	}
	expected := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar":   "baz",
			"hello": "world",
		},
	}
	if !reflect.DeepEqual(expected, dst) {
		t.Error("Merge failure")
	}
}

func TestIntersection(t *testing.T) {
	m1 := map[string]interface{}{
		"a": true,
		"b": true,
		"c": true,
		"d": true,
	}
	m2 := map[string]interface{}{
		"c": true,
		"d": true,
		"x": true,
	}
	iSect := intersection(m1, m2)
	if len(iSect) != 2 {
		t.Error("Incorrect intersection ", iSect)
	}
	if !elementOf(iSect, "c") {
		t.Error("Missing element c")
	}
	if !elementOf(iSect, "d") {
		t.Error("Missing element d")
	}
}

func TestUnmarshalStructureSuccess(t *testing.T) {
	text := `{
		"main": {"type": "number"}
	}`
	structure, err := CreateJSONStructure([]byte(text), DefaultOptions())
	if err != nil {
		t.Error("Unmarshal error", err)
	}
	text = `{
		"fragments": {
			"req": {
				"multipleOf": 2
			}
		},
		"types": {
			"bar": {
				"type": "integer",
				"multipleOf": 4
			},
			"foo": {
				"\u2384": [ "req", "bar" ],
				"type": "integer"
			}
		},
		"main": {
			"type": "foo"
		}
	}`
	structure, err = CreateJSONStructure([]byte(text), DefaultOptions())
	if err != nil {
		t.Error("Unmarshal error", err)
	}
	if structure.Definition.Types["foo"].MultipleOf == nil {
		t.Error("Composition failure")
	}
	if !structure.Definition.Types["foo"].MultipleOf.Equal(decimal.NewFromFloat(4.0)) {
		t.Error("Composition failure")
	}
	if structure.Definition.Types["foo"].Type != "integer" {
		t.Error("Composition failure")
	}
}

func TestUnmarshalStructureFailure(t *testing.T) {
	text := `{}`
	_, err := CreateJSONStructure([]byte(text), DefaultOptions())
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	text = `{"fragments": "foo", "types": "bar", "main": {}}`
	_, err = CreateJSONStructure([]byte(text), DefaultOptions())
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	text = `{"fragments": "foo", "types": "bar", "main": {}}`
	_, err = CreateJSONStructure([]byte(text), DefaultOptions())
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	text = `{
		"fragments": {"a": true, "b":true},
	 	"types": {"a": true, "b":true}, 
		"main": {}}`
	_, err = CreateJSONStructure([]byte(text), DefaultOptions())
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	text = `{
	 	"types": {"a": true, "b":true}, 
		"main": {}}`
	_, err = CreateJSONStructure([]byte(text), DefaultOptions())
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	text = `{
		"fragments": {
			"number": {"type": "boolean"}
		},
		"types": {
			"integer": {"type": "boolean"}
		},		
		"main": {
			"type": "boolean"
		}
	}`
	_, err = CreateJSONStructure([]byte(text), DefaultOptions())
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	text = `{
		"fragments": {
			"a": {
				"\u2384": [ "a" ]
			}
		},
		"main": {
			"type": "foo"
		}
	}`
	_, err = CreateJSONStructure([]byte(text), DefaultOptions())
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	text = `{
		"fragments": {
			"a": {
				"\u2384": [ "b" ]
			}
		},
		"main": {
			"type": "foo"
		}
	}`
	_, err = CreateJSONStructure([]byte(text), DefaultOptions())
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
	text = `{
		"fragments": {
			"a": {
				"\u2384": [ true ]
			},
			"b": {
				"\u2384": 5
			}
		},
		"main": {
			"type": "foo"
		}
	}`
	_, err = CreateJSONStructure([]byte(text), DefaultOptions())
	if err == nil {
		t.Error("Expected error")
	}
	t.Log(err)
}

func TestMarshalStructure(t *testing.T) {
	var result JSONStructureDefinition
	var data []byte
	dec, _ := decimal.NewFromString("2")
	s1 := EmptyJSONStructure()
	s1.Definition.Main = &TypeDecl{}
	s1.Definition.Main.Type = "integer"
	s1.Definition.Main.MultipleOf = &dec
	s1.Definition.Main.Minimum = &dec
	s1.Definition.Main.Maximum = &dec
	err := s1.ValidateStructure()
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	data, err = json.MarshalIndent(s1.Definition, "", "  ")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	if result.Main.MultipleOfRaw == nil {
		t.Error("Marshal failure")
	}
	if result.Main.MinimumRaw == nil {
		t.Error("Marshal failure")
	}
	if result.Main.MaximumRaw == nil {
		t.Error("Marshal failure")
	}
	s2 := EmptyJSONStructure()
	s2.Definition.Main = &TypeDecl{}
	s2.Definition.Main.Type = "integer"
	s2.Definition.Main.ExclusiveMinimum = &dec
	s2.Definition.Main.ExclusiveMaximum = &dec
	err = s2.ValidateStructure()
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	data, err = json.MarshalIndent(s2.Definition, "", "  ")
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Error("Unexpected error ", err)
	}
	if result.Main.ExclusiveMinimumRaw == nil {
		t.Error("Marshal failure")
	}
	if result.Main.ExclusiveMaximumRaw == nil {
		t.Error("Marshal failure")
	}
}
