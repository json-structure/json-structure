package jsonstructure

import (
	"errors"
	"fmt"

	multierror "github.com/mspiegel/go-multierror"
)

func unionFilterErrors(errs map[string]error, structure *JSONStructure, scope []string) error {
	if structure.Options.UnionError == AllUnionReport {
		return unionAllErrors(errs)
	}
	return unionPriorityErrors(errs, structure, scope)
}

func filterPriority(errs []error, scope []string) error {
	var result error
	for _, err := range errs {
		switch e := err.(type) {
		case *EnumError:
			if len(e.Scope)-len(scope) <= 1 {
				result = multierror.Append(result, err)
			}
		}
	}
	return result
}

func unionPriorityErrors(errs map[string]error, structure *JSONStructure, scope []string) error {
	result := errors.New("union validation failure. Reporting a subset of errors for each type")
	lowPriority := make(map[string]error)
	hasPriority := false
	for k, err := range errs {
		var filter error
		switch v := err.(type) {
		case *multierror.Error:
			filter = filterPriority(v.Errors, scope)
		default:
			filter = filterPriority([]error{v}, scope)
		}
		lowPriority[k] = filter
		if filter == nil {
			hasPriority = true
		}
	}
	for k, err := range errs {
		p := lowPriority[k]
		if hasPriority && p != nil {
			continue
		}
		err = fmt.Errorf("Attempted to validate against union type '%s' - %s", k, err.Error())
		result = multierror.Append(result, err)
	}
	return result
}

func unionAllErrors(errs map[string]error) error {
	result := errors.New("union validation failure. Reporting all errors for each type")
	for k, err := range errs {
		err = fmt.Errorf("Attempted to validate against union type '%s' - %s", k, err.Error())
		result = multierror.Append(result, err)
	}
	return result
}
