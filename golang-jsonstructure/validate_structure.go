package jsonstructure

import (
	"errors"

	multierror "github.com/mspiegel/go-multierror"
)

func (structure *JSONStructure) ValidateStructure() error {
	err := validateStructureTopLevel(structure)
	if err != nil {
		return err
	}
	err = validateStructureDecls(structure)
	if err != nil {
		return err
	}
	err = validateEmbeddedObjects(structure)
	if err != nil {
		return err
	}
	return nil
}

func (structure *JSONStructure) doValidation() {
	structure.InitError = structure.ValidateStructure()
}

func validateStructureTopLevel(structure *JSONStructure) error {
	var errs error
	definition := &structure.Definition
	for k, v := range definition.Types {
		scope := []string{"types", k}
		if v == nil {
			err := errors.New("type declaration must be non-nil")
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
		}
	}
	scope := []string{"main"}
	if definition.Main == nil {
		err := errors.New("main type declaration must be declared")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	return errs

}

func validateStructureDecls(structure *JSONStructure) error {
	var errs error
	definition := &structure.Definition
	for k, v := range definition.Types {
		scope := []string{"types", k}
		err := v.ValidateDecl(structure, scope)
		errs = multierror.Append(errs, err)
	}
	scope := []string{"main"}
	err := definition.Main.ValidateDecl(structure, scope)
	errs = multierror.Append(errs, err)
	return errs
}

func validateEmbeddedObjects(structure *JSONStructure) error {
	var errs error
	definition := &structure.Definition
	for k, v := range definition.Types {
		scope := []string{"types", k}
		err := v.ValidateEmbedded(structure, scope)
		errs = multierror.Append(errs, err)
	}
	scope := []string{"main"}
	err := definition.Main.ValidateEmbedded(structure, scope)
	errs = multierror.Append(errs, err)
	return errs
}
