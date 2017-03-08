package jsonstructure

import (
	"fmt"
	"strings"
)

type SchemaError struct {
	Scope []string
	Err   error
}

type ValidationError struct {
	SchemaScope []string
	ObjectScope []string
	Err         error
}

func (e *SchemaError) Error() string {
	return fmt.Sprintf("At /%s: %s", strings.Join(e.Scope, "/"), e.Err.Error())
}

func errorAt(err error, scope []string) error {
	if err == nil {
		return nil
	}
	return &SchemaError{
		Scope: scope,
		Err:   err,
	}
}
