package bootstrap

import (
	"encoding/json"
	"testing"

	jsonstructure "github.com/json-structure/json-structure/golang-jsonstructure"
)

func TestBootstrap(t *testing.T) {
	var err error
	var structure *jsonstructure.JSONStructure
	var data, transform []byte
	data, err = Asset("json-structure.json")
	if err != nil {
		t.Fatal("Error reading bindata ", err)
	}
	structure, err = jsonstructure.CreateJSONStructure(data, jsonstructure.DefaultOptions())
	if err != nil {
		t.Fatal("Error creating JSON Structure ", err)
	}
	transform, err = json.Marshal(structure.Definition)
	if err != nil {
		t.Fatal("Error marshalling JSON Structure ", err)
	}
	err = structure.Validate(transform)
	if err != nil {
		t.Fatal("Error validating JSON Structure using JSON Structure ", err)
	}
}
