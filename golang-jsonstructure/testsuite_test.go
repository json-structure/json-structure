package jsonstructure

import (
	"encoding/json"
	"testing"

	"github.com/json-structure/json-structure/golang-jsonstructure/testsuite"
)

type TestSuite struct {
	Description string                   `json:"description"`
	Structure   *JSONStructureDefinition `json:"structure"`
	Valid       bool                     `json:"valid"`
}

func TestTestSuite(t *testing.T) {
	for _, name := range testsuite.AssetNames() {
		var suites []TestSuite
		data, err := testsuite.Asset(name)
		if err != nil {
			t.Errorf("Unable to read %s: %s", name, err)
			continue
		}
		err = json.Unmarshal(data, &suites)
		if err != nil {
			t.Errorf("Unable to parse %s: %s", name, err)
			continue
		}
		for _, suite := range suites {
			if suite.Structure == nil {
				t.Errorf("Missing JSON structure %s, %s", name, suite.Description)
				continue
			}
			structure := EmptyJSONStructure()
			structure.Definition = *suite.Structure
			err = structure.ValidateStructure()
			if err != nil && suite.Valid {
				t.Errorf("Unexpected structure validation error %s, %s, %s", name, suite.Description, err)
				continue
			} else if err != nil {
				t.Logf("%s: %s", name, err)
			} else if err == nil && !suite.Valid {
				t.Errorf("Expected structure validation error %s, %s", name, suite.Description)
				continue
			}
		}
	}
}
