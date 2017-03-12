package jsonstructure

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/json-structure/json-structure/golang-jsonstructure/simpleregexp"
	multierror "github.com/mspiegel/go-multierror"
	"github.com/shopspring/decimal"
)

var PrimitiveTypes = map[string]bool{
	"boolean": true,
	"integer": true,
	"number":  true,
	"string":  true,
	"json":    true,
	"struct":  true,
	"array":   true,
	"union":   true,
}

type TypeDecl struct {
	// required
	Type string `json:"type"`
	// store RawMessage to distinguish between "nil" and "null"
	// property 'default' is available to all type declarations
	DefaultRaw json.RawMessage `json:"default,omitempty"`
	Default    interface{}     `json:"-"`
	// common to primitive types
	Format   *string `json:"format,omitempty"`
	Nullable *bool   `json:"nullable,omitempty"`
	// number
	MultipleOf       *decimal.Decimal `json:"multipleOf,omitempty"`
	Minimum          *decimal.Decimal `json:"minimum,omitempty"`
	Maximum          *decimal.Decimal `json:"maximum,omitempty"`
	ExclusiveMinimum *decimal.Decimal `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum *decimal.Decimal `json:"exclusiveMaximum,omitempty"`
	// string
	PatternRaw *string        `json:"pattern,omitempty"`
	Pattern    *regexp.Regexp `json:"-"`
	MinLength  *int           `json:"minLength,omitempty"`
	MaxLength  *int           `json:"maxLength,omitempty"`
	// struct
	Fields map[string]*TypeDecl `json:"fields,omitempty"`
	// array
	Items    *TypeDecl `json:"items,omitempty"`
	MinItems *int      `json:"minItems,omitempty"`
	MaxItems *int      `json:"maxItems,omitempty"`
	// union
	Types map[string]*TypeDecl `json:"types,omitempty"`
}

var PermissibleFields = map[string]map[string]bool{
	"boolean": map[string]bool{
		"format":   true,
		"nullable": true,
	},
	"integer": map[string]bool{
		"format":           true,
		"nullable":         true,
		"multipleOf":       true,
		"minimum":          true,
		"maximum":          true,
		"exclusiveMinimum": true,
		"exclusiveMaximum": true,
	},
	"number": map[string]bool{
		"format":           true,
		"nullable":         true,
		"multipleOf":       true,
		"minimum":          true,
		"maximum":          true,
		"exclusiveMinimum": true,
		"exclusiveMaximum": true,
	},
	"string": map[string]bool{
		"format":    true,
		"nullable":  true,
		"pattern":   true,
		"minLength": true,
		"maxLength": true,
	},
	"json": map[string]bool{
		"format":   true,
		"nullable": true,
	},
	"struct": map[string]bool{
		"format":   true,
		"nullable": true,
		"fields":   true,
	},
	"array": map[string]bool{
		"format":   true,
		"nullable": true,
		"items":    true,
		"minItems": true,
		"maxItems": true,
	},
	"union": map[string]bool{
		"format":   true,
		"nullable": true,
	},
}

type shadowDecl TypeDecl

func (td *TypeDecl) UnmarshalJSON(data []byte) error {
	var shadow shadowDecl

	err := json.Unmarshal(data, &shadow)
	if err != nil {
		return err
	}
	if len(shadow.DefaultRaw) > 0 {
		reader := bytes.NewReader(shadow.DefaultRaw)
		decoder := json.NewDecoder(reader)
		decoder.UseNumber()
		err := decoder.Decode(&shadow.Default)
		if err != nil {
			return err
		}
	}
	*td = TypeDecl(shadow)
	return nil
}

func (td *TypeDecl) ValidateDecl(structure JSONStructure, scope []string) error {
	var errs error
	if len(td.Type) == 0 {
		err := errors.New("missing required property 'type'")
		return errorAt(err, scope)
	}
	pf := PermissibleFields[td.Type]
	decl := structure.Definition.Types[td.Type]
	if pf == nil && decl == nil {
		err := fmt.Errorf("Unknown type '%s'", td.Type)
		return errorAt(err, scope)
	}
	if pf == nil {
		err := detectTypeAliasCycle(structure, decl, nil)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	e1 := permissible("format", td.Type, pf, td.Format != nil, scope)
	e2 := permissible("nullable", td.Type, pf, td.Nullable != nil, scope)
	e3 := permissible("multipleOf", td.Type, pf, td.MultipleOf != nil, scope)
	e4 := permissible("minimum", td.Type, pf, td.Minimum != nil, scope)
	e5 := permissible("maximum", td.Type, pf, td.Maximum != nil, scope)
	e6 := permissible("exclusiveMinimum", td.Type, pf, td.ExclusiveMinimum != nil, scope)
	e7 := permissible("exclusiveMaximum", td.Type, pf, td.ExclusiveMaximum != nil, scope)
	e8 := permissible("pattern", td.Type, pf, td.PatternRaw != nil, scope)
	e9 := permissible("minLength", td.Type, pf, td.MinLength != nil, scope)
	e10 := permissible("maxLength", td.Type, pf, td.MaxLength != nil, scope)
	e11 := permissible("fields", td.Type, pf, td.Fields != nil, scope)
	e12 := permissible("items", td.Type, pf, td.Items != nil, scope)
	e13 := permissible("minItems", td.Type, pf, td.MinItems != nil, scope)
	e14 := permissible("maxItems", td.Type, pf, td.MaxItems != nil, scope)
	errs = multierror.Append(errs, e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14)
	e1 = validateNumberTypeDecl(td, scope)
	e2 = validateStringTypeDecl(td, structure, scope)
	e3 = validateStructTypeDecl(td, structure, scope)
	e4 = validateArrayTypeDecl(td, structure, scope)
	e5 = validateUnionTypeDecl(td, structure, scope)
	e6 = validateFormatTypeDecl(td, structure, scope)
	errs = multierror.Append(errs, e1, e2, e3, e4, e5, e6)
	return errs
}

func permissible(name string, typ string, fields map[string]bool, observed bool, scope []string) error {
	allowed := fields[name]
	if observed && !allowed {
		err := fmt.Errorf("Property %s is not allowed with type %s", name, typ)
		return errorAt(err, scope)
	}
	return nil
}

func detectTypeAliasCycle(structure JSONStructure, td *TypeDecl, prev map[string]bool) error {
	name := td.Type
	decl := structure.Definition.Types[td.Type]
	if prev[name] {
		keys := make([]string, 0, len(prev))
		for k := range prev {
			keys = append(keys, k)
		}
		return fmt.Errorf("Type alias cycle detected %v", keys)
	}
	if decl == nil {
		return nil
	}
	if prev == nil {
		prev = make(map[string]bool)
	}
	prev[name] = true
	return detectTypeAliasCycle(structure, decl, prev)
}

func validateNumberTypeDecl(td *TypeDecl, scope []string) error {
	var errs error
	if td.Type != "integer" && td.Type != "number" {
		return nil
	}
	if td.Minimum != nil && td.ExclusiveMinimum != nil {
		err := errors.New("'minimum' and 'exclusiveMinimum' are both defined")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.Maximum != nil && td.ExclusiveMaximum != nil {
		err := errors.New("'maximum' and 'exclusiveMaximum' are both defined")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	var min, max *decimal.Decimal
	if td.Minimum != nil {
		min = td.Minimum
	} else if td.ExclusiveMinimum != nil {
		min = td.ExclusiveMinimum
	}
	if td.Maximum != nil {
		max = td.Maximum
	} else if td.ExclusiveMaximum != nil {
		max = td.ExclusiveMaximum
	}
	if min != nil && max != nil && min.Cmp(*max) > 0 {
		err := errors.New("maximum is less than minimum")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.MultipleOf != nil && td.MultipleOf.Cmp(decimal.Zero) < 0 {
		err := errors.New("'multipleOf' must be a non-negative number")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	return errs
}

func validateStringTypeDecl(td *TypeDecl, structure JSONStructure, scope []string) error {
	var errs error
	if td.Type != "string" {
		return nil
	}
	if td.MinLength != nil && *td.MinLength < 0 {
		err := errors.New("'minLength' must be a non-negative integer")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.MaxLength != nil && *td.MaxLength < 0 {
		err := errors.New("'maxLength' must be a non-negative integer")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.MinLength != nil && td.MaxLength != nil && *td.MinLength > *td.MaxLength {
		err := errors.New("'maxLength' is less than 'minLength'")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.PatternRaw != nil {
		var err error
		td.Pattern, err = regexp.Compile(*td.PatternRaw)
		if err != nil {
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
		} else if structure.Options.Regex == StrictRegex {
			err = simpleregexp.Accept(*td.PatternRaw)
			if err != nil {
				err = fmt.Errorf("regular expression is not cross-platform. "+
					"Consider the NativeRegex option. %s",
					err.Error())
				err = errorAt(err, scope)
				errs = multierror.Append(errs, err)
			}
		}
	}

	return errs
}

func validateStructTypeDecl(td *TypeDecl, structure JSONStructure, scope []string) error {
	var errs error
	if td.Type != "struct" {
		return nil
	}
	if td.Fields == nil {
		err := errors.New("missing required property 'fields'")
		err = errorAt(err, scope)
		return err
	}
	for k, v := range td.Fields {
		newscope := append(scope, "fields", k)
		err := v.ValidateDecl(structure, newscope)
		errs = multierror.Append(errs, err)
	}
	return errs
}

func validateArrayTypeDecl(td *TypeDecl, structure JSONStructure, scope []string) error {
	var errs error
	if td.Type != "array" {
		return nil
	}
	if td.Items == nil {
		err := errors.New("missing required property 'items'")
		err = errorAt(err, scope)
		return err
	}
	newscope := append(scope, "items")
	err := td.Items.ValidateDecl(structure, newscope)
	errs = multierror.Append(errs, err)
	if td.MinItems != nil && *td.MinItems < 0 {
		err := errors.New("'minItems' must be a non-negative integer")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.MaxItems != nil && *td.MaxItems < 0 {
		err := errors.New("'maxItems' must be a non-negative integer")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	if td.MinItems != nil && td.MaxItems != nil && *td.MinItems > *td.MaxItems {
		err := errors.New("'maxItems' is less than 'minItems'")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	return errs
}

func validateUnionTypeDecl(td *TypeDecl, structure JSONStructure, scope []string) error {
	var errs error
	if td.Type != "union" {
		return nil
	}
	if td.Types == nil {
		err := errors.New("missing required property 'types'")
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	}
	for k, v := range td.Types {
		newscope := append(scope, "types", k)
		err := v.ValidateDecl(structure, newscope)
		errs = multierror.Append(errs, err)
	}
	return errs
}

func validateFormatTypeDecl(td *TypeDecl, structure JSONStructure, scope []string) error {
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
	if !format.Accept(td.Type) {
		err := fmt.Errorf(`format "%s" does not accept type "%s"`, name, td.Type)
		err = errorAt(err, scope)
		return err
	}
	return nil
}

func (td *TypeDecl) ValidateDefault(structure JSONStructure, scope []string) error {
	var errs error
	if len(td.DefaultRaw) > 0 {
		err := td.Validate(td.Default, structure, scope)
		if err != nil {
			err = fmt.Errorf("Unable to validate default value. %s", err.Error())
			errs = multierror.Append(errs, err)
		}
	}
	for k, v := range td.Fields {
		newscope := append(scope, k)
		err := v.ValidateDefault(structure, newscope)
		errs = multierror.Append(errs, err)
	}
	if td.Items != nil {
		newscope := append(scope, "items")
		err := td.Items.ValidateDefault(structure, newscope)
		errs = multierror.Append(errs, err)
	}
	return errs
}
