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
	err := validateStructureDecls(structure)
	if err != nil {
		return err
	}
	err = validateStructureDefaults(structure)
	if err != nil {
		return err
	}
	return nil
}

func validateStructureDecls(structure *JSONStructure) error {
	var errs error
	for k, v := range structure.Types {
		scope := []string{"types", k}
		if v == nil {
			err := errors.New("type declaration must be non-nil")
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
		} else {
			err := v.ValidateDecl(structure, scope)
			errs = multierror.Append(errs, err)
		}
	}
	scope := []string{"main"}
	if structure.Main == nil {
		err := errors.New("type declaration must be non-nil")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	} else {
		err := structure.Main.ValidateDecl(structure, scope)
		errs = multierror.Append(errs, err)
	}
	return errs
}

func validateStructureDefaults(structure *JSONStructure) error {
	var errs error
	for k, v := range structure.Types {
		scope := []string{"types", k}
		err := v.ValidateDefault(structure, scope)
		errs = multierror.Append(errs, err)
	}
	scope := []string{"main"}
	err := structure.Main.ValidateDefault(structure, scope)
	errs = multierror.Append(errs, err)
	return errs
}
