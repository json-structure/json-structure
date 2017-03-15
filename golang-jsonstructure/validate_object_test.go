package jsonstructure

import (
	"encoding/json"
	"testing"
)

func TestValidateFailureUnknownType(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "string"
	structure.InitOnce.Do(structure.doValidation)
	err := structure.InitError
	if err != nil {
		t.Fatal("Unexpected validation error", err)
	}
	structure.Definition.Main.Type = "foobar"
	text := `"foobar"`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}

func TestValidateFailureUnion(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "union"
	structure.Definition.Main.Types = make(map[string]*TypeDecl)
	structure.Definition.Main.Types["foo"] = &TypeDecl{}
	structure.Definition.Main.Types["foo"].Type = "struct"
	structure.Definition.Main.Types["foo"].Fields = make(map[string]*TypeDecl)
	structure.Definition.Main.Types["foo"].Fields["type"] = &TypeDecl{}
	structure.Definition.Main.Types["foo"].Fields["type"].Type = "string"
	structure.Definition.Main.Types["foo"].Fields["type"].EnumRaw = []json.RawMessage{json.RawMessage(`"foo"`)}
	structure.Definition.Main.Types["foo"].Fields["foo"] = &TypeDecl{}
	structure.Definition.Main.Types["foo"].Fields["foo"].Type = "integer"
	structure.Definition.Main.Types["foo"].Fields["bar"] = &TypeDecl{}
	structure.Definition.Main.Types["foo"].Fields["bar"].Type = "integer"
	structure.Definition.Main.Types["bar"] = &TypeDecl{}
	structure.Definition.Main.Types["bar"].Type = "struct"
	structure.Definition.Main.Types["bar"].Fields = make(map[string]*TypeDecl)
	structure.Definition.Main.Types["bar"].Fields["type"] = &TypeDecl{}
	structure.Definition.Main.Types["bar"].Fields["type"].Type = "string"
	structure.Definition.Main.Types["bar"].Fields["type"].EnumRaw = []json.RawMessage{json.RawMessage(`"bar"`)}

	text := `{"type": "foo"}`

	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Options.UnionError = AllUnionReport
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}
