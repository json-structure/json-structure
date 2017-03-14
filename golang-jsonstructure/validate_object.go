package jsonstructure

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	multierror "github.com/mspiegel/go-multierror"
	"github.com/shopspring/decimal"
)

func (structure *JSONStructure) Validate(text []byte) error {
	var value interface{}
	structure.InitOnce.Do(structure.doValidation)
	err := structure.InitError
	if err != nil {
		return err
	}
	reader := bytes.NewReader(text)
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	err = decoder.Decode(&value)
	if err != nil {
		return err
	}
	return structure.Definition.Main.Validate(value, structure, nil)
}

func (td *TypeDecl) Validate(value interface{}, structure *JSONStructure, scope []string) error {
	var errs error
	name := td.Type

	if !PrimitiveTypes[name] {
		def := structure.Definition.Types[name]
		if def == nil {
			err := fmt.Errorf("Unknown type '%s'", name)
			return errorAt(err, scope)
		}
		return def.Validate(value, structure, scope)
	}

	if value == nil {
		if td.Nullable != nil && *td.Nullable {
			return nil
		}
		err := errors.New("JSON null value when nullable property is false")
		return errorAt(err, scope)
	}

	err := validateFormat(td, value, structure, scope)
	errs = multierror.Append(errs, err)
	err = validateEnum(td, value, structure, scope)
	errs = multierror.Append(errs, err)

	switch name {
	case "boolean":
		err := validateBoolean(td, value, structure, scope)
		errs = multierror.Append(errs, err)
	case "integer":
		err := validateInteger(td, value, structure, scope)
		errs = multierror.Append(errs, err)
	case "number":
		err := validateNumber(td, value, structure, scope)
		errs = multierror.Append(errs, err)
	case "string":
		err := validateString(td, value, structure, scope)
		errs = multierror.Append(errs, err)
	case "json":
		// no validation
	case "struct":
		err := validateStruct(td, value, structure, scope)
		errs = multierror.Append(errs, err)
	case "array":
		err := validateArray(td, value, structure, scope)
		errs = multierror.Append(errs, err)
	case "set":
		err := validateSet(td, value, structure, scope)
		errs = multierror.Append(errs, err)
	case "map":
		err := validateMap(td, value, structure, scope)
		errs = multierror.Append(errs, err)
	case "union":
		err := validateUnion(td, value, structure, scope)
		errs = multierror.Append(errs, err)
	default:
		err := fmt.Errorf("Internal error. Unhandled primitive type '%s'", name)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	return errs
}

func validateBoolean(td *TypeDecl, value interface{}, structure *JSONStructure, scope []string) error {
	_, ok := value.(bool)
	if !ok {
		err := errors.New("JSON value is not a boolean")
		return errorAt(err, scope)
	}
	return nil
}

func validateNumber(td *TypeDecl, value interface{}, structure *JSONStructure, scope []string) error {
	var errs error
	num, ok := value.(json.Number)
	if !ok {
		err := errors.New("JSON value is not a number")
		return errorAt(err, scope)
	}
	dec, err := decimal.NewFromString(num.String())
	if err != nil {
		return errorAt(err, scope)
	}
	if td.Minimum != nil && td.Minimum.Cmp(dec) > 0 {
		err := fmt.Errorf("%s is less than minimum %s", dec, td.Minimum)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.ExclusiveMinimum != nil && td.ExclusiveMinimum.Cmp(dec) >= 0 {
		err := fmt.Errorf("%s is less than or equal to exclusive minimum %s", dec, td.ExclusiveMinimum)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.Maximum != nil && td.Maximum.Cmp(dec) < 0 {
		err := fmt.Errorf("%s is greater than maximum %s", dec, td.Maximum)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.ExclusiveMaximum != nil && td.ExclusiveMaximum.Cmp(dec) <= 0 {
		err := fmt.Errorf("%s is greater than or equal to exclusive maximum %s", dec, td.ExclusiveMaximum)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.MultipleOf != nil && dec.Mod(*td.MultipleOf).Cmp(decimal.Zero) != 0 {
		err := fmt.Errorf("%s is not a multiple of %s", dec, td.MultipleOf)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	return errs
}

func validateInteger(td *TypeDecl, value interface{}, structure *JSONStructure, scope []string) error {
	err := validateNumber(td, value, structure, scope)
	if err != nil {
		return err
	}
	num := value.(json.Number)
	dec, _ := decimal.NewFromString(num.String())
	if dec.Truncate(0).Cmp(dec) != 0 {
		err := fmt.Errorf("%s is not an integer", dec)
		return errorAt(err, scope)
	}
	return nil
}

func validateString(td *TypeDecl, value interface{}, structure *JSONStructure, scope []string) error {
	var errs error
	str, ok := value.(string)
	if !ok {
		err := errors.New("JSON value is not a string")
		return errorAt(err, scope)
	}
	if td.MinLength != nil && len(str) < *td.MinLength {
		err := fmt.Errorf(`length of string "%s" is less than minimum length %d`, str, *td.MinLength)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.MaxLength != nil && len(str) > *td.MaxLength {
		err := fmt.Errorf(`length of string "%s" is greater than maximum length %d`, str, *td.MaxLength)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.Pattern != nil {
		idx := td.Pattern.FindStringIndex(str)
		if idx == nil {
			err := errors.New("'pattern' regular expression is not a match")
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
		}
	}
	return errs
}

func validateUnion(td *TypeDecl, value interface{}, structure *JSONStructure, scope []string) error {
	var err error
	errs := make(map[string]error)
	for k, decl := range td.Types {
		err = decl.Validate(value, structure, scope)
		if err == nil {
			return nil
		}
		errs[k] = err
	}
	return unionFilterErrors(errs, structure, scope)
}

func validateStruct(td *TypeDecl, value interface{}, structure *JSONStructure, scope []string) error {
	var errs error
	obj, ok := value.(map[string]interface{})
	if !ok {
		err := errors.New("JSON value is not an object")
		return errorAt(err, scope)
	}
	for name, decl := range td.Fields {
		child, ok2 := obj[name]
		if ok2 {
			newscope := append(scope, name)
			err := decl.Validate(child, structure, newscope)
			errs = multierror.Append(errs, err)
		} else if len(decl.DefaultRaw) == 0 {
			err := fmt.Errorf(`missing required field "%s"`, name)
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
		}
	}
	for name := range obj {
		if td.Fields[name] == nil {
			err := fmt.Errorf(`unrecognized field "%s"`, name)
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
		}
	}
	return errs
}

func validateArray(td *TypeDecl, value interface{}, structure *JSONStructure, scope []string) error {
	var errs error
	arr, ok := value.([]interface{})
	if !ok {
		err := errors.New("JSON value is not an array")
		return errorAt(err, scope)
	}
	if td.MinItems != nil && len(arr) < *td.MinItems {
		err := fmt.Errorf("length of array %d is less than minimum items %d", len(arr), *td.MinItems)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.MaxItems != nil && len(arr) > *td.MaxItems {
		err := fmt.Errorf("length of array %d is greater than maximum items %d", len(arr), *td.MaxItems)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	for i, val := range arr {
		newscope := append(scope, strconv.Itoa(i))
		err := td.Items.Validate(val, structure, newscope)
		errs = multierror.Append(errs, err)
	}
	return errs
}

func validateSet(td *TypeDecl, value interface{}, structure *JSONStructure, scope []string) error {
	var errs error
	arr, ok := value.([]interface{})
	if !ok {
		err := errors.New("JSON value is not a set")
		return errorAt(err, scope)
	}
	if td.MinItems != nil && len(arr) < *td.MinItems {
		err := fmt.Errorf("size of set %d is less than minimum items %d", len(arr), *td.MinItems)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.MaxItems != nil && len(arr) > *td.MaxItems {
		err := fmt.Errorf("size of set %d is greater than maximum items %d", len(arr), *td.MaxItems)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	uniq := createSet()
	for i, val := range arr {
		newscope := append(scope, strconv.Itoa(i))
		insert, err := uniq.PutIfAbsent(val)
		if err != nil {
			err = errorAt(err, newscope)
			errs = multierror.Append(errs, err)
		} else if !insert {
			err = errors.New("set has duplicate value")
			err = errorAt(err, newscope)
			errs = multierror.Append(errs, err)
		} else {
			err = td.Items.Validate(val, structure, newscope)
			errs = multierror.Append(errs, err)
		}
	}
	return errs
}

func validateMap(td *TypeDecl, value interface{}, structure *JSONStructure, scope []string) error {
	var errs error
	obj, ok := value.(map[string]interface{})
	if !ok {
		err := errors.New("JSON value is not a map")
		return errorAt(err, scope)
	}
	if td.MinItems != nil && len(obj) < *td.MinItems {
		err := fmt.Errorf("size of map %d is less than minimum items %d", len(obj), *td.MinItems)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.MaxItems != nil && len(obj) > *td.MaxItems {
		err := fmt.Errorf("size of map %d is greater than maximum items %d", len(obj), *td.MaxItems)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	for k, v := range obj {
		newscope := append(scope, k)
		err := td.Items.Validate(v, structure, newscope)
		errs = multierror.Append(errs, err)
	}
	return errs
}

func validateFormat(td *TypeDecl, value interface{}, structure *JSONStructure, scope []string) error {
	if td.Format == nil {
		return nil
	}
	name := *td.Format
	format := structure.Options.Formats[name]
	if format == nil {
		err := fmt.Errorf(`unknown format "%s"`, name)
		err = errorAt(err, scope)
		return err
	}
	err := format.Apply(value, td.Type)
	err = fmt.Errorf(`Validation error on format "%s": %s`, name, err.Error())
	err = errorAt(err, scope)
	return err
}

func validateEnum(td *TypeDecl, value interface{}, structure *JSONStructure, scope []string) error {
	if td.Enum == nil {
		return nil
	}
	found, err := td.Enum.Has(value)
	if err != nil {
		err = errorAt(err, scope)
		return err
	}
	if !found {
		err = errors.New("value not found in 'enum' set")
		err = enumError(err, scope)
		return err
	}
	return nil
}
