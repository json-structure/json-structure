package jsonstructure

import "encoding/json"

type JSONStructure struct {
	Definition *JSONStructureDefinition
	Options    *JSONStructureOptions
}

type JSONStructureDefinition struct {
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	Imports     map[string]string          `json:"imports"`
	Fragments   map[string]json.RawMessage `json:"fragments"`
	Types       map[string]*TypeDecl       `json:"types"`
	Main        *TypeDecl                  `json:"main"`
}

func EmptyJSONStructure() JSONStructure {
	return JSONStructure{
		Definition: &JSONStructureDefinition{},
		Options:    DefaultOptions(),
	}
}

func CreateJSONStructure(data []byte, options *JSONStructureOptions) (JSONStructure, error) {
	var definition JSONStructureDefinition
	err := json.Unmarshal(data, &definition)
	if err != nil {
		return JSONStructure{}, err
	}
	if options == nil {
		options = DefaultOptions()
	}
	result := JSONStructure{
		Definition: &definition,
		Options:    options,
	}
	err = result.ValidateStructure()
	if err != nil {
		return JSONStructure{}, err
	}
	return result, nil
}
