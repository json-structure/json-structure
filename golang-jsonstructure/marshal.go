package jsonstructure

import (
	"bytes"
	"encoding/json"

	"github.com/shopspring/decimal"
)

type shadowStructure JSONStructureDefinition
type shadowTypeDecl TypeDecl

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
	err = Compose(shell)
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

func decToJSONNumber(dec *decimal.Decimal) *json.Number {
	val := dec.String()
	num := json.Number(val)
	return &num
}

func (td *TypeDecl) MarshalJSON() ([]byte, error) {
	var shadow *shadowTypeDecl
	if td.Pattern != nil {
		raw := td.Pattern.String()
		td.PatternRaw = &raw
	}
	if td.MultipleOf != nil {
		td.MultipleOfRaw = decToJSONNumber(td.MultipleOf)
	}
	if td.Minimum != nil {
		td.MinimumRaw = decToJSONNumber(td.Minimum)
	}
	if td.Maximum != nil {
		td.MaximumRaw = decToJSONNumber(td.Maximum)
	}
	if td.ExclusiveMinimum != nil {
		td.ExclusiveMinimumRaw = decToJSONNumber(td.ExclusiveMinimum)
	}
	if td.ExclusiveMaximum != nil {
		td.ExclusiveMaximumRaw = decToJSONNumber(td.ExclusiveMaximum)
	}
	shadow = (*shadowTypeDecl)(td)
	return json.Marshal(shadow)
}
