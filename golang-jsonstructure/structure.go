package jsonstructure

import (
	"encoding/json"
	"errors"

	multierror "github.com/mspiegel/go-multierror"
)

type JSONStructure struct {
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	Imports     map[string]string          `json:"imports"`
	Fragments   map[string]json.RawMessage `json:"fragments"`
	Types       map[string]*TypeDecl       `json:"types"`
	Main        *TypeDecl                  `json:"main"`
}

func (structure *JSONStructure) ValidateStructure() error {
	var errs error
	for k, v := range structure.Types {
		scope := []string{"types", k}
		if v == nil {
			err := errors.New("type declaration must be non-nil")
			err = errorAt(err, scope)
			errs = multierror.AppendNonNil(errs, err)
		} else {
			err := v.Validate(structure, scope)
			errs = multierror.AppendNonNil(errs, err)
		}
	}
	scope := []string{"main"}
	if structure.Main == nil {
		err := errors.New("type declaration must be non-nil")
		err = errorAt(err, scope)
		errs = multierror.AppendNonNil(errs, err)
	} else {
		err := structure.Main.Validate(structure, scope)
		errs = multierror.AppendNonNil(errs, err)
	}
	return errs
}
