package jsonstructure

import (
	"encoding/json"
	"sync"
)

type JSONStructure struct {
	Definition JSONStructureDefinition
	Options    JSONStructureOptions
	InitOnce   sync.Once
	InitError  error
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
		Definition: JSONStructureDefinition{},
		Options:    DefaultOptions(),
		InitOnce:   sync.Once{},
	}
}

func CreateJSONStructure(data []byte, options JSONStructureOptions) (*JSONStructure, error) {
	var definition JSONStructureDefinition
	err := json.Unmarshal(data, &definition)
	if err != nil {
		return nil, err
	}
	result := JSONStructure{
		Definition: definition,
		Options:    options,
		InitOnce:   sync.Once{},
	}
	result.InitOnce.Do(result.doValidation)
	err = result.InitError
	if err != nil {
		return nil, err
	}
	return &result, nil
}
