package jsonstructure

import "fmt"

type SchemaError struct {
	Scope string
	Err   error
}

type EnumError struct {
	SchemaError
}

func (e *SchemaError) Error() string {
	msg := e.Scope
	if len(msg) == 0 {
		msg = "/"
	}
	return fmt.Sprintf("At %s: %s", msg, e.Err.Error())
}

func (e *EnumError) Error() string {
	msg := e.Scope
	if len(msg) == 0 {
		msg = "/"
	}
	return fmt.Sprintf("At %s: %s", msg, e.Err.Error())
}

func errorAt(err error, scope string) error {
	if err == nil {
		return nil
	}
	return &SchemaError{
		Scope: scope,
		Err:   err,
	}
}

func enumError(err error, scope string) error {
	if err == nil {
		return nil
	}
	return &EnumError{
		SchemaError: SchemaError{
			Scope: scope,
			Err:   err,
		},
	}
}
