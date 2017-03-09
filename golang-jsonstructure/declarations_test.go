package jsonstructure

import (
	"testing"

	multierror "github.com/mspiegel/go-multierror"
	"github.com/shopspring/decimal"
)

var negOneInt = -1
var negOneDec = decimal.NewFromFloat(-1.0)
var zero = 0
var posOneInt = 1

func TestValidateStructureSuccess(t *testing.T) {
	structure := JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "struct"
	structure.Main.Fields = make(map[string]*TypeDecl)
	structure.Main.Fields["foo"] = &TypeDecl{}
	structure.Main.Fields["foo"].Type = "boolean"
	err := structure.ValidateStructure()
	if err != nil {
		t.Error("Validation failure", err)
	}
}

func TestValidateStructureFailure(t *testing.T) {
	structure := JSONStructure{}
	structure.Types = make(map[string]*TypeDecl)
	structure.Types["foo"] = nil
	structure.Types["bar"] = nil
	err := structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = JSONStructure{}
	structure.Main = &TypeDecl{}
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure when 'type' is empty")
	}
	t.Log(err)
	structure = JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "foo"
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure when 'type' is foo")
	}
	t.Log(err)
	structure = JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "boolean"
	structure.Main.MultipleOf = &decimal.Zero
	structure.Main.Minimum = &decimal.Zero
	structure.Main.Maximum = &decimal.Zero
	structure.Main.ExclusiveMinimum = &decimal.Zero
	structure.Main.ExclusiveMaximum = &decimal.Zero
	structure.Main.Pattern = "foobar"
	structure.Main.MinLength = &zero
	structure.Main.MaxLength = &zero
	structure.Main.Fields = make(map[string]*TypeDecl)
	structure.Main.Items = &TypeDecl{}
	structure.Main.MinItems = &zero
	structure.Main.MaxItems = &zero
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	if len(err.(*multierror.Error).Errors) != 12 {
		t.Error("Expected 12 errors")
	}
	t.Log(err)
	structure = JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "number"
	structure.Main.MultipleOf = &negOneDec
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "number"
	structure.Main.Minimum = &decimal.Zero
	structure.Main.Maximum = &decimal.Zero
	structure.Main.ExclusiveMinimum = &decimal.Zero
	structure.Main.ExclusiveMaximum = &decimal.Zero
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "number"
	structure.Main.Minimum = &decimal.Zero
	structure.Main.Maximum = &negOneDec
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "string"
	structure.Main.MinLength = &negOneInt
	structure.Main.MaxLength = &negOneInt
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "string"
	structure.Main.MinLength = &posOneInt
	structure.Main.MaxLength = &zero
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "struct"
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "array"
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "array"
	structure.Main.Items = &TypeDecl{}
	structure.Main.Items.Type = "boolean"
	structure.Main.MinItems = &negOneInt
	structure.Main.MaxItems = &negOneInt
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
	structure = JSONStructure{}
	structure.Main = &TypeDecl{}
	structure.Main.Type = "array"
	structure.Main.Items = &TypeDecl{}
	structure.Main.Items.Type = "boolean"
	structure.Main.MinItems = &posOneInt
	structure.Main.MaxItems = &zero
	err = structure.ValidateStructure()
	if err == nil {
		t.Error("Expected validation failure")
	}
	t.Log(err)
}
