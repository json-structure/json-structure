package jsonstructure

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
)

type JSONStructure struct {
	Definition JSONStructureDefinition
	Options    JSONStructureOptions
	StructErr  error
}

var errNotInit = errors.New("")

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
		StructErr:  errNotInit,
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
		StructErr:  errNotInit,
	}
	err = result.ValidateStructure()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (structure *JSONStructure) JSONMarshalDefinition() ([]byte, error) {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	enc := json.NewEncoder(writer)
	enc.SetIndent("", "  ")
	err := enc.Encode(structure.Definition)
	return buf.Bytes(), err
}
