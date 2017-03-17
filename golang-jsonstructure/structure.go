package jsonstructure

import "encoding/json"

type JSONStructure struct {
	definition JSONStructureDefinition
	options    JSONStructureOptions
	init       bool
	initErr    error
}

type JSONStructureDefinition struct {
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	Imports     map[string]string          `json:"imports"`
	Fragments   map[string]json.RawMessage `json:"fragments"`
	Types       map[string]*TypeDecl       `json:"types"`
	Main        *TypeDecl                  `json:"main"`
}

func EmptyJSONStructure() *JSONStructure {
	return &JSONStructure{
		definition: JSONStructureDefinition{},
		options:    DefaultOptions(),
		init:       false,
		initErr:    nil,
	}
}

func CreateJSONStructure(data []byte, options JSONStructureOptions) (*JSONStructure, error) {
	var definition JSONStructureDefinition
	err := json.Unmarshal(data, &definition)
	if err != nil {
		return nil, err
	}
	result := JSONStructure{
		definition: definition,
		options:    options,
		init:       false,
		initErr:    nil,
	}
	err = result.ValidateStructure()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (structure *JSONStructure) JSONMarshalDefinition() ([]byte, error) {
	return json.Marshal(structure.definition)
}
