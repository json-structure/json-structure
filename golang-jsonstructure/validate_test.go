package jsonstructure

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestValidateSuccess(t *testing.T) {
	text := `{
		"a": 1,
		"b": 1.5,
		"c": "c",
		"d": true,
		"e": [1, 2, 3],
		"f": null
	}`
	structure := JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "struct"
	structure.Main.Fields = make(map[string]*TypeDecl)
	structure.Main.Fields["a"] = &TypeDecl{Type: "integer"}
	structure.Main.Fields["b"] = &TypeDecl{Type: "number"}
	structure.Main.Fields["c"] = &TypeDecl{Type: "string"}
	structure.Main.Fields["d"] = &TypeDecl{Type: "boolean"}
	structure.Main.Fields["e"] = &TypeDecl{Type: "array", Items: &TypeDecl{Type: "integer"}}
	structure.Main.Fields["f"] = &TypeDecl{Type: "boolean", Nullable: true}
	err := structure.Validate([]byte(text))
	if err != nil {
		t.Error("JSON object validation error ", err)
	}
}

func TestValidateFailureBoolean(t *testing.T) {
	structure := JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "boolean"
	text := `"hello"`
	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}

func TestValidateFailureNumber(t *testing.T) {
	structure := JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "number"
	text := `"hello"`
	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Main.Minimum = nil
	structure.Main.Maximum = &decimal.Zero
	text = `1`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Main.Minimum = &decimal.Zero
	structure.Main.Maximum = nil
	text = `-1`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Main.Minimum = nil
	structure.Main.Maximum = nil
	structure.Main.ExclusiveMinimum = &decimal.Zero
	structure.Main.ExclusiveMaximum = &decimal.Zero
	text = `0`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Main.ExclusiveMinimum = nil
	structure.Main.ExclusiveMaximum = nil
	hundredths, _ := decimal.NewFromString("0.01")
	structure.Main.MultipleOf = &hundredths
	text = `0.001`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}

func TestValidateFailureInteger(t *testing.T) {
	structure := JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "integer"
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
	structure := JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "string"
	text := `1.5`
	err := structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Main.MinLength = &one
	text = `""`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Main.MaxLength = &one
	text = `"foo"`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}

func TestValidateFailureStruct(t *testing.T) {
	structure := JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "struct"
	structure.Main.Fields = make(map[string]*TypeDecl)
	structure.Main.Fields["a"] = &TypeDecl{Type: "integer"}
	structure.Main.Fields["b"] = &TypeDecl{Type: "string", DefaultRaw: []byte(`"foo"`)}
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

func TestValidateFailureArray(t *testing.T) {
	one := 1
	structure := JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "array"
	structure.Main.Items = &TypeDecl{Type: "integer"}
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
	structure.Main.MinItems = &one
	text = `[]`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
	structure.Main.MaxItems = &one
	text = `[1, 2, 3]`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)

}

func TestValidateFailureAlias(t *testing.T) {
	structure := JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "foo"
	structure.Types = make(map[string]*TypeDecl)
	structure.Types["foo"] = &TypeDecl{Type: "string"}
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
	structure.Main.Type = "bar"
	text = `"foo"`
	err = structure.Validate([]byte(text))
	if err == nil {
		t.Error("JSON object validation did not fail")
	}
	t.Log(err)
}
