package jsonstructure

import (
	"bytes"
	"encoding/json"
)

type shadowStructure JSONStructureDefinition

func (structure *JSONStructureDefinition) UnmarshalJSON(data []byte) error {
	var shell map[string]interface{}
	var shadow shadowStructure

	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	err := decoder.Decode(&shell)
	if err != nil {
		return err
	}
	err = transformCompose(shell)
	if err != nil {
		return err
	}
	data, err = json.Marshal(shell)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &shadow)
	if err != nil {
		return err
	}
	*structure = JSONStructureDefinition(shadow)
	return nil
}
