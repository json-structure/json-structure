package jsonstructure

import (
	"errors"
	"fmt"

	multierror "github.com/mspiegel/go-multierror"
	"github.com/shopspring/decimal"
)

var PrimitiveTypes = map[string]bool{
	"boolean": true,
	"integer": true,
	"number":  true,
	"string":  true,
	"struct":  true,
	"array":   true,
}

type TypeDecl struct {
	// common
	Type     string `json:"type"`
	Format   string `json:"format"`
	Required bool   `json:"required"`
	Nullable bool   `json:"nullable"`
	// number
	MultipleOf       *decimal.Decimal `json:"multipleOf"`
	Minimum          *decimal.Decimal `json:"minimum"`
	Maximum          *decimal.Decimal `json:"maximum"`
	ExclusiveMinimum *decimal.Decimal `json:"exclusiveMinimum"`
	ExclusiveMaximum *decimal.Decimal `json:"exclusiveMaximum"`
	// string
	Pattern   string `json:"pattern"`
	MinLength *int   `json:"minLength"`
	MaxLength *int   `json:"maxLength"`
	// struct
	Fields map[string]*TypeDecl `json:"fields"`
	// array
	Items    *TypeDecl `json:"items"`
	MinItems *int      `json:"minItems"`
	MaxItems *int      `json:"maxItems"`
}

var PermissibleFields = map[string]map[string]bool{
	"boolean": map[string]bool{},
	"integer": map[string]bool{
		"multipleOf":       true,
		"minimum":          true,
		"maximum":          true,
		"exclusiveMinimum": true,
		"exclusiveMaximum": true,
	},
	"number": map[string]bool{
		"multipleOf":       true,
		"minimum":          true,
		"maximum":          true,
		"exclusiveMinimum": true,
		"exclusiveMaximum": true,
	},
	"string": map[string]bool{
		"pattern":   true,
		"minLength": true,
		"maxLength": true,
	},
	"struct": map[string]bool{
		"fields": true,
	},
	"array": map[string]bool{
		"items":    true,
		"minItems": true,
		"maxItems": true,
	},
}

func (td *TypeDecl) Validate(structure *JSONStructure, scope []string) error {
	var errs error
	if len(td.Type) == 0 {
		err := errors.New("missing required property 'type'")
		return errorAt(err, scope)
	}
	pf := PermissibleFields[td.Type]
	decl := structure.Types[td.Type]
	if pf == nil && decl == nil {
		err := fmt.Errorf("Unknown type '%s'", td.Type)
		return errorAt(err, scope)
	}
	e1 := permissible("multipleOf", td.Type, pf, td.MultipleOf != nil, scope)
	e2 := permissible("minimum", td.Type, pf, td.Minimum != nil, scope)
	e3 := permissible("maximum", td.Type, pf, td.Maximum != nil, scope)
	e4 := permissible("exclusiveMinimum", td.Type, pf, td.ExclusiveMinimum != nil, scope)
	e5 := permissible("exclusiveMaximum", td.Type, pf, td.ExclusiveMaximum != nil, scope)
	e6 := permissible("pattern", td.Type, pf, len(td.Pattern) > 0, scope)
	e7 := permissible("minLength", td.Type, pf, td.MinLength != nil, scope)
	e8 := permissible("maxLength", td.Type, pf, td.MaxLength != nil, scope)
	e9 := permissible("fields", td.Type, pf, td.Fields != nil, scope)
	e10 := permissible("items", td.Type, pf, td.Items != nil, scope)
	e11 := permissible("minItems", td.Type, pf, td.MinItems != nil, scope)
	e12 := permissible("maxItems", td.Type, pf, td.MaxItems != nil, scope)
	errs = multierror.AppendNonNil(errs, e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12)
	e1 = validateNumberTypeDecl(td, scope)
	e2 = validateStringTypeDecl(td, scope)
	e3 = validateStructTypeDecl(td, structure, scope)
	e4 = validateArrayTypeDecl(td, structure, scope)
	errs = multierror.AppendNonNil(errs, e1, e2, e3, e4)
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

func validateNumberTypeDecl(td *TypeDecl, scope []string) error {
	var errs error
	if td.Type != "integer" && td.Type != "number" {
		return nil
	}
	if td.Minimum != nil && td.ExclusiveMinimum != nil {
		err := errors.New("'minimum' and 'exclusiveMinimum' are both defined")
		err = errorAt(err, scope)
		errs = multierror.AppendNonNil(errs, err)
	}
	if td.Maximum != nil && td.ExclusiveMaximum != nil {
		err := errors.New("'maximum' and 'exclusiveMaximum' are both defined")
		err = errorAt(err, scope)
		errs = multierror.AppendNonNil(errs, err)
	}
	min := decimal.Zero
	max := decimal.Zero
	if td.Minimum != nil {
		min = *td.Minimum
	} else if td.ExclusiveMinimum != nil {
		min = *td.ExclusiveMinimum
	}
	if td.Maximum != nil {
		max = *td.Maximum
	} else if td.ExclusiveMaximum != nil {
		max = *td.ExclusiveMaximum
	}
	if min.Cmp(max) > 0 {
		err := errors.New("maximum is less than minimum")
		err = errorAt(err, scope)
		errs = multierror.AppendNonNil(errs, err)
	}
	if td.MultipleOf != nil && td.MultipleOf.Cmp(decimal.Zero) < 0 {
		err := errors.New("'multipleOf' must be a non-negative number")
		err = errorAt(err, scope)
		errs = multierror.AppendNonNil(errs, err)
	}
	return errs
}

func validateStringTypeDecl(td *TypeDecl, scope []string) error {
	var errs error
	if td.Type != "string" {
		return nil
	}
	if td.MinLength != nil && *td.MinLength < 0 {
		err := errors.New("'minLength' must be a non-negative integer")
		err = errorAt(err, scope)
		errs = multierror.AppendNonNil(errs, err)
	}
	if td.MaxLength != nil && *td.MaxLength < 0 {
		err := errors.New("'maxLength' must be a non-negative integer")
		err = errorAt(err, scope)
		errs = multierror.AppendNonNil(errs, err)
	}
	if td.MinLength != nil && td.MaxLength != nil && *td.MinLength > *td.MaxLength {
		err := errors.New("'maxLength' is less than 'minLength'")
		err = errorAt(err, scope)
		errs = multierror.AppendNonNil(errs, err)
	}
	return errs
}

func validateStructTypeDecl(td *TypeDecl, structure *JSONStructure, scope []string) error {
	var errs error
	if td.Type != "struct" {
		return nil
	}
	if td.Fields != nil {
		err := errors.New("missing required property 'fields'")
		err = errorAt(err, scope)
		return err
	}
	for k, v := range td.Fields {
		newscope := append(scope, "fields", k)
		err := v.Validate(structure, newscope)
		errs = multierror.AppendNonNil(errs, err)
	}
	return errs
}

func validateArrayTypeDecl(td *TypeDecl, structure *JSONStructure, scope []string) error {
	var errs error
	if td.Type != "array" {
		return nil
	}
	if td.Items != nil {
		err := errors.New("missing required property 'items'")
		err = errorAt(err, scope)
		return err
	}
	newscope := append(scope, "items")
	err := td.Items.Validate(structure, newscope)
	errs = multierror.AppendNonNil(errs, err)
	if td.MinItems != nil && *td.MinItems < 0 {
		err := errors.New("'minItems' must be a non-negative integer")
		err = errorAt(err, scope)
		errs = multierror.AppendNonNil(errs, err)
	}
	if td.MaxItems != nil && *td.MaxItems < 0 {
		err := errors.New("'maxItems' must be a non-negative integer")
		err = errorAt(err, scope)
		errs = multierror.AppendNonNil(errs, err)
	}
	if td.MinItems != nil && td.MaxItems != nil && *td.MinItems > *td.MaxItems {
		err := errors.New("'maxItems' is less than 'minItems'")
		err = errorAt(err, scope)
		errs = multierror.AppendNonNil(errs, err)
	}
	return errs
}
