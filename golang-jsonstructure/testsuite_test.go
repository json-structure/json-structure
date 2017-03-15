package jsonstructure

import (
	"encoding/json"
	"testing"

	"github.com/json-structure/json-structure/golang-jsonstructure/testsuite"
)

type TestDeclaration struct {
	Description string          `json:"description"`
	Structure   json.RawMessage `json:"structure"`
	Valid       bool            `json:"valid"`
	Tests       []*TestCase     `json:"tests"`
}

type TestCase struct {
	Description string          `json:"description"`
	Data        json.RawMessage `json:"data"`
	Valid       bool            `json:"valid"`
}

func TestSuite(t *testing.T) {
	for _, name := range testsuite.AssetNames() {
		var suites []TestDeclaration
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
			err = json.Unmarshal(suite.Structure, &structure.Definition)
			if err == nil {
				err = structure.ValidateStructure()
			}
			if err != nil && suite.Valid {
				t.Errorf("Unexpected structure validation error %s, %s, %s", name, suite.Description, err)
				continue
			} else if err != nil {
				t.Logf("%s: %s", name, err)
				continue
			} else if err == nil && !suite.Valid {
				t.Errorf("Expected structure validation error %s, %s", name, suite.Description)
				continue
			}
			for _, test := range suite.Tests {
				err = structure.Validate(test.Data)
				if err != nil && test.Valid {
					t.Errorf("Unexpected object validation error %s, %s, %s, %s", name, suite.Description, test.Description, err)
				} else if err != nil {
					t.Logf("%s: %s", name, err)
				} else if err == nil && !test.Valid {
					t.Errorf("Expected object validation error %s, %s, %s", name, suite.Description, test.Description)
				}
			}
		}
	}
}
