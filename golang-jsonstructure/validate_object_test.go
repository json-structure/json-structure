package jsonstructure

import (
	"regexp"
	"testing"

	"github.com/shopspring/decimal"
)

func TestValidateSuccess(t *testing.T) {
	affirmative := true
	text := `{
		"a": 1,
		"b": 1.5,
		"c": "c",
		"d": true,
		"e": [1, 2, 3],
		"f": null
	}`
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "struct"
	structure.Definition.Main.Fields = make(map[string]*TypeDecl)
	structure.Definition.Main.Fields["a"] = &TypeDecl{Type: "integer"}
	structure.Definition.Main.Fields["b"] = &TypeDecl{Type: "number"}
	structure.Definition.Main.Fields["c"] = &TypeDecl{Type: "string"}
	structure.Definition.Main.Fields["d"] = &TypeDecl{Type: "boolean"}
	structure.Definition.Main.Fields["e"] = &TypeDecl{Type: "array", Items: &TypeDecl{Type: "integer"}}
	structure.Definition.Main.Fields["f"] = &TypeDecl{Type: "boolean", Nullable: &affirmative}
	err := structure.Validate([]byte(text))
	if err != nil {
		t.Error("JSON object validation error ", err)
	}
}

func TestValidateFailureBoolean(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "boolean"
	text := `"hello"`
	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}

func TestValidateFailureNumber(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "number"
	text := `"hello"`
	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Definition.Main.Minimum = nil
	structure.Definition.Main.Maximum = &decimal.Zero
	text = `1`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Definition.Main.Minimum = &decimal.Zero
	structure.Definition.Main.Maximum = nil
	text = `-1`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Definition.Main.Minimum = nil
	structure.Definition.Main.Maximum = nil
	structure.Definition.Main.ExclusiveMinimum = &decimal.Zero
	structure.Definition.Main.ExclusiveMaximum = &decimal.Zero
	text = `0`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Definition.Main.ExclusiveMinimum = nil
	structure.Definition.Main.ExclusiveMaximum = nil
	hundredths, _ := decimal.NewFromString("0.01")
	structure.Definition.Main.MultipleOf = &hundredths
	text = `0.001`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}

func TestValidateFailureInteger(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "integer"
	text := `"hello"`
	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	text = `1.5`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}

func TestValidateFailureString(t *testing.T) {
	one := 1
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "string"
	text := `1.5`
	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Definition.Main.MinLength = &one
	text = `""`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Definition.Main.MaxLength = &one
	text = `"foo"`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "string"
	structure.Definition.Main.Pattern = regexp.MustCompile("[0-9]+")
	text = `"foobar"`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}

func TestValidateFailureStruct(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "struct"
	structure.Definition.Main.Fields = make(map[string]*TypeDecl)
	structure.Definition.Main.Fields["a"] = &TypeDecl{Type: "integer"}
	structure.Definition.Main.Fields["b"] = &TypeDecl{Type: "string", DefaultRaw: []byte(`"foo"`)}
	text := `1.5`
	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	text = `{}`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	text = `{"a": "a", "b": 3}`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	text = `{"c": "c"}`
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
	text := `{}`
	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Definition.Main.Types["foo"] = &TypeDecl{}
	structure.Definition.Main.Types["foo"].Type = "integer"
	structure.Definition.Main.Types["bar"] = &TypeDecl{}
	structure.Definition.Main.Types["bar"].Type = "string"
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}

func TestValidateFailureArray(t *testing.T) {
	one := 1
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "array"
	structure.Definition.Main.Items = &TypeDecl{Type: "integer"}
	text := `1.5`
	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	text = `["a", true, null]`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Definition.Main.MinItems = &one
	text = `[]`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Definition.Main.MaxItems = &one
	text = `[1, 2, 3]`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)

}

func TestValidateFailureAlias(t *testing.T) {
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "foo"
	structure.Definition.Types = make(map[string]*TypeDecl)
	structure.Definition.Types["foo"] = &TypeDecl{Type: "string"}
	text := `1.5`
	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	text = `null`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Definition.Main.Type = "bar"
	text = `"foo"`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}

func TestValidateFailureFormat(t *testing.T) {
	text := `"hello@world@com"`
	format := "email"
	structure := EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "string"
	structure.Definition.Main.Format = &format
	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	format = "foobar"
	structure = EmptyJSONStructure()
	structure.Definition.Main = &TypeDecl{}
	structure.Definition.Main.Type = "string"
	structure.Definition.Main.Format = &format
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}
