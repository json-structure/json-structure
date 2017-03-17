package jsonstructure

import (
	"encoding/json"
	"testing"
)

func TestValidateFailureUnknownType(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.definition.Main = &TypeDecl{}
	structure.definition.Main.Type = "string"
	err := structure.ValidateStructure()
	if err != nil {
		t.Fatal("Unexpected validation error", err)
	}
	structure.definition.Main.Type = "foobar"
	text := `"foobar"`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}

func TestValidateFailureUnion(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.definition.Main = &TypeDecl{}
	structure.definition.Main.Type = "union"
	structure.definition.Main.Types = make(map[string]*TypeDecl)
	structure.definition.Main.Types["foo"] = &TypeDecl{}
	structure.definition.Main.Types["foo"].Type = "struct"
	structure.definition.Main.Types["foo"].Fields = make(map[string]*TypeDecl)
	structure.definition.Main.Types["foo"].Fields["type"] = &TypeDecl{}
	structure.definition.Main.Types["foo"].Fields["type"].Type = "string"
	structure.definition.Main.Types["foo"].Fields["type"].EnumRaw = []json.RawMessage{json.RawMessage(`"foo"`)}
	structure.definition.Main.Types["foo"].Fields["foo"] = &TypeDecl{}
	structure.definition.Main.Types["foo"].Fields["foo"].Type = "integer"
	structure.definition.Main.Types["foo"].Fields["bar"] = &TypeDecl{}
	structure.definition.Main.Types["foo"].Fields["bar"].Type = "integer"
	structure.definition.Main.Types["bar"] = &TypeDecl{}
	structure.definition.Main.Types["bar"].Type = "struct"
	structure.definition.Main.Types["bar"].Fields = make(map[string]*TypeDecl)
	structure.definition.Main.Types["bar"].Fields["type"] = &TypeDecl{}
	structure.definition.Main.Types["bar"].Fields["type"].Type = "string"
	structure.definition.Main.Types["bar"].Fields["type"].EnumRaw = []json.RawMessage{json.RawMessage(`"bar"`)}

	text := `{"type": "foo"}`

	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.options.UnionError = AllUnionReport
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}
