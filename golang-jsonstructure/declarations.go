package jsonstructure

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"

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
	"set":     true,
	"map":     true,
	"union":   true,
}

type TypeDecl struct {
	// required
	Type string `json:"type"`
	// common to primitive types
	Format   *string `json:"format,omitempty"`
	Nullable *bool   `json:"nullable,omitempty"`
	Optional *bool   `json:"optional,omitempty"`
	// store RawMessage to unmarhshal numbers correctly
	DefaultRaw json.RawMessage   `json:"default,omitempty"`
	EnumRaw    []json.RawMessage `json:"enum,omitempty"`
	Default    interface{}       `json:"-"`
	Enum       sethash           `json:"-"`
	// number
	// decimal library marshals JSON with quotes
	// (unless terrible hack flag is used)
	// preserving raw JSON numbers for marshaling
	MultipleOfRaw       *json.Number `json:"multipleOf,omitempty"`
	MinimumRaw          *json.Number `json:"minimum,omitempty"`
	MaximumRaw          *json.Number `json:"maximum,omitempty"`
	ExclusiveMinimumRaw *json.Number `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximumRaw *json.Number `json:"exclusiveMaximum,omitempty"`
	// number (parsed)
	MultipleOf       *decimal.Decimal `json:"-"`
	Minimum          *decimal.Decimal `json:"-"`
	Maximum          *decimal.Decimal `json:"-"`
	ExclusiveMinimum *decimal.Decimal `json:"-"`
	ExclusiveMaximum *decimal.Decimal `json:"-"`
	// string
	PatternRaw *string        `json:"pattern,omitempty"`
	Pattern    *regexp.Regexp `json:"-"`
	MinLength  *int           `json:"minLength,omitempty"`
	MaxLength  *int           `json:"maxLength,omitempty"`
	// struct
	Fields map[string]*TypeDecl `json:"fields,omitempty"`
	// array, set, and map types
	Items    *TypeDecl `json:"items,omitempty"`
	MinItems *int      `json:"minItems,omitempty"`
	MaxItems *int      `json:"maxItems,omitempty"`
	// union
	Types map[string]*TypeDecl `json:"types,omitempty"`
}

var commonFields = map[string]bool{
	"enum":     true,
	"default":  true,
	"format":   true,
	"nullable": true,
	"optional": true,
}

var PermissibleFields map[string]map[string]bool

func init() {
	PermissibleFields = map[string]map[string]bool{
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
		"json": map[string]bool{},
		"struct": map[string]bool{
			"fields": true,
		},
		"array": map[string]bool{
			"items":    true,
			"minItems": true,
			"maxItems": true,
		},
		"set": map[string]bool{
			"items":    true,
			"minItems": true,
			"maxItems": true,
		},
		"map": map[string]bool{
			"items":    true,
			"minItems": true,
			"maxItems": true,
		},
		"union": map[string]bool{
			"types": true,
		},
	}
	for _, m := range PermissibleFields {
		for k := range commonFields {
			m[k] = true
		}
	}
}

func (td *TypeDecl) ValidateDecl(structure *JSONStructure, scope []string) error {
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
	if pf != nil && decl != nil {
		err := fmt.Errorf("Unexpected error. Type declaration shadows primitive type '%s'", td.Type)
		return errorAt(err, scope)
	}
	if decl != nil {
		err := detectTypeAliasCycle(structure, decl, nil)
		err = errorAt(err, scope)
		errs = multierror.Append(errs, err)
	} else {
		err := validateFormatTypeDecl(td, structure, scope)
		errs = multierror.Append(errs, err)
	}
	e1 := permissible("enum", td.Type, pf, td.EnumRaw != nil || td.Enum != nil, scope)
	e2 := permissible("default", td.Type, pf, td.DefaultRaw != nil || td.Default != nil, scope)
	e3 := permissible("format", td.Type, pf, td.Format != nil, scope)
	e4 := permissible("nullable", td.Type, pf, td.Nullable != nil, scope)
	e5 := permissible("optional", td.Type, pf, td.Optional != nil, scope)
	e6 := permissible("multipleOf", td.Type, pf, td.MultipleOfRaw != nil || td.MultipleOf != nil, scope)
	e7 := permissible("minimum", td.Type, pf, td.MinimumRaw != nil || td.Minimum != nil, scope)
	e8 := permissible("maximum", td.Type, pf, td.MaximumRaw != nil || td.Maximum != nil, scope)
	e9 := permissible("exclusiveMinimum", td.Type, pf, td.ExclusiveMinimumRaw != nil || td.ExclusiveMinimum != nil, scope)
	e10 := permissible("exclusiveMaximum", td.Type, pf, td.ExclusiveMaximumRaw != nil || td.ExclusiveMaximum != nil, scope)
	e11 := permissible("pattern", td.Type, pf, td.PatternRaw != nil || td.Pattern != nil, scope)
	e12 := permissible("minLength", td.Type, pf, td.MinLength != nil, scope)
	e13 := permissible("maxLength", td.Type, pf, td.MaxLength != nil, scope)
	e14 := permissible("fields", td.Type, pf, td.Fields != nil, scope)
	e15 := permissible("items", td.Type, pf, td.Items != nil, scope)
	e16 := permissible("minItems", td.Type, pf, td.MinItems != nil, scope)
	e17 := permissible("maxItems", td.Type, pf, td.MaxItems != nil, scope)
	e18 := permissible("types", td.Type, pf, td.Types != nil, scope)
	errs = multierror.Append(errs, e1, e2, e3, e4, e5, e6, e7, e8, e9,
		e10, e11, e12, e13, e14, e15, e16, e17, e18)

	e1 = validateNumberTypeDecl(td, scope)
	e2 = validateStringTypeDecl(td, structure, scope)
	e3 = validateStructTypeDecl(td, structure, scope)
	e4 = validateCollectionTypeDecl(td, structure, scope)
	e5 = validateUnionTypeDecl(td, structure, scope)
	errs = multierror.Append(errs, e1, e2, e3, e4, e5)

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

func detectTypeAliasCycle(structure *JSONStructure, td *TypeDecl, prev map[string]bool) error {
	name := td.Type
	decl := structure.Definition.Types[td.Type]
	if prev[name] {
		keys := keysSet(prev)
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

func convertNumber(num *json.Number, scope []string, name string) (*decimal.Decimal, error) {
	dec, err := decimal.NewFromString(num.String())
	if err != nil {
		err = errorAt(err, append(scope, name))
		return nil, err
	}
	return &dec, nil
}

func convertNumbers(td *TypeDecl, scope []string) error {
	var e1, e2, e3, e4, e5 error
	if td.MinimumRaw != nil {
		td.Minimum, e1 = convertNumber(td.MinimumRaw, scope, "minimum")
	}
	if td.MaximumRaw != nil {
		td.Maximum, e2 = convertNumber(td.MaximumRaw, scope, "maximum")
	}
	if td.ExclusiveMinimumRaw != nil {
		td.ExclusiveMinimum, e3 = convertNumber(td.ExclusiveMinimumRaw, scope, "exclusiveMinimum")
	}
	if td.ExclusiveMaximumRaw != nil {
		td.ExclusiveMaximum, e4 = convertNumber(td.ExclusiveMaximumRaw, scope, "exclusiveMaximum")
	}
	if td.MultipleOfRaw != nil {
		td.MultipleOf, e5 = convertNumber(td.MultipleOfRaw, scope, "multipleOf")
	}
	return multierror.Append(nil, e1, e2, e3, e4, e5)
}

func validateNumberTypeDecl(td *TypeDecl, scope []string) error {
	var errs error
	if td.Type != "integer" && td.Type != "number" {
		return nil
	}
	errs = convertNumbers(td, scope)
	if errs != nil {
		return errs
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
		err := errors.New("'maximum' is less than 'minimum'")
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

func validateStringTypeDecl(td *TypeDecl, structure *JSONStructure, scope []string) error {
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
			err = fmt.Errorf("'pattern' expression. %s", err)
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
		} else if structure.Options.Regex == StrictRegex {
			err = simpleregexp.Accept(*td.PatternRaw)
			if err != nil {
				err = fmt.Errorf("'pattern' expression is not cross-platform. "+
					"Consider the NativeRegex option. %s",
					err.Error())
				err = errorAt(err, scope)
				errs = multierror.Append(errs, err)
			}
		}
	}

	return errs
}

func validateStructTypeDecl(td *TypeDecl, structure *JSONStructure, scope []string) error {
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
		if v == nil {
			continue
		}
		newscope := append(scope, "fields", k)
		err := v.ValidateDecl(structure, newscope)
		errs = multierror.Append(errs, err)
	}
	return errs
}

func validateCollectionTypeDecl(td *TypeDecl, structure *JSONStructure, scope []string) error {
	var errs error
	if td.Type != "array" && td.Type != "set" && td.Type != "map" {
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

func validateUnionTypeDecl(td *TypeDecl, structure *JSONStructure, scope []string) error {
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
		if v == nil {
			continue
		}
		newscope := append(scope, "types", k)
		err := v.ValidateDecl(structure, newscope)
		errs = multierror.Append(errs, err)
	}
	return errs
}

func validateFormatTypeDecl(td *TypeDecl, structure *JSONStructure, scope []string) error {
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

func (td *TypeDecl) ValidateEmbedded(structure *JSONStructure, scope []string) error {
	var err, errs error
	// enums must be populated before default is validated
	if td.EnumRaw != nil {
		td.Enum = createSet()
		for i, raw := range td.EnumRaw {
			iScope := append(scope, "enum", strconv.Itoa(i))
			var insert bool
			var value interface{}
			reader := bytes.NewReader(raw)
			decoder := json.NewDecoder(reader)
			decoder.UseNumber()
			err = decoder.Decode(&value)
			if err != nil {
				err = errorAt(err, iScope)
				errs = multierror.Append(errs, err)
				continue
			}
			if value == nil {
				err = errors.New("null enum value is not permitted")
				err = errorAt(err, iScope)
				errs = multierror.Append(errs, err)
				continue
			}
			insert, err = td.Enum.PutIfAbsent(value)
			if err != nil {
				err = errorAt(err, iScope)
				errs = multierror.Append(errs, err)
			} else if !insert {
				err = errors.New("duplicate enum value")
				err = errorAt(err, iScope)
				errs = multierror.Append(errs, err)
			} else {
				err = td.ValidateValue(value, structure, iScope)
				if err != nil {
					errs = multierror.Append(errs, err)
				}
			}
		}
	}
	newscope := append(scope, "default")
	defaultDecoded := false
	if td.DefaultRaw != nil {
		reader := bytes.NewReader(td.DefaultRaw)
		decoder := json.NewDecoder(reader)
		decoder.UseNumber()
		err = decoder.Decode(&td.Default)
		if err != nil {
			err = errorAt(err, newscope)
			errs = multierror.Append(errs, err)
		} else {
			defaultDecoded = true
		}
	}
	if defaultDecoded || (td.Default != nil) || (td.Optional != nil && *td.Optional) {
		err = td.ValidateValue(td.Default, structure, newscope)
		if err != nil {
			errs = multierror.Append(errs, err)
		}
	}
	for k, v := range td.Fields {
		if v == nil {
			continue
		}
		newscope = append(scope, "fields", k)
		err = v.ValidateEmbedded(structure, newscope)
		errs = multierror.Append(errs, err)
	}
	for k, v := range td.Types {
		if v == nil {
			continue
		}
		newscope = append(scope, "types", k)
		err = v.ValidateEmbedded(structure, newscope)
		errs = multierror.Append(errs, err)
	}
	if td.Items != nil {
		newscope = append(scope, "items")
		err = td.Items.ValidateEmbedded(structure, newscope)
		errs = multierror.Append(errs, err)
	}
	return errs
}
