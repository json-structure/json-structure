package jsonstructure

import (
	"errors"
	"fmt"

	"github.com/mohae/deepcopy"
	multierror "github.com/mspiegel/go-multierror"
)

var ComposeSymbol = "\u0ADD"

func mapMerge(dst map[string]interface{}, src map[string]interface{}, scope []string) error {
	var errs error
	for k, v := range src {
		newscope := append(scope, k)
		srcMap, srcOK := v.(map[string]interface{})
		dstValue := dst[k]
		if dstValue != nil {
			dstMap, dstOK := dstValue.(map[string]interface{})
			if (srcOK && !dstOK) || (!srcOK && dstOK) {
				err := errors.New("Attempted merge between map and non-map types")
				err = errorAt(err, newscope)
				errs = multierror.Append(errs, err)
			} else if srcOK && dstOK {
				err := mapMerge(dstMap, srcMap, newscope)
				errs = multierror.Append(errs, err)
			} else {
				dst[k] = deepcopy.Copy(v)
			}
		} else {
			dst[k] = deepcopy.Copy(v)
		}
	}
	return errs
}

func doCompose(target interface{},
	types map[string]interface{},
	fragments map[string]interface{},
	prev map[string]bool,
	scope []string) error {
	var err, errs error
	object, ok := target.(map[string]interface{})
	if !ok {
		err = errors.New("definition is not a JSON object")
		err = errorAt(err, scope)
		return err
	}
	for k, v := range object {
		if obj, ok2 := v.(map[string]interface{}); ok2 {
			newscope := append(scope, k)
			err = doCompose(obj, types, fragments, prev, newscope)
			errs = multierror.Append(errs, err)
		}
	}
	if errs != nil {
		return errs
	}
	c := object[ComposeSymbol]
	if c == nil {
		return nil
	}
	defs, ok3 := c.([]interface{})
	if !ok3 {
		err = errors.New("'\\u0ADD' property must be an string array")
		return errorAt(err, scope)
	}
	dest := make(map[string]interface{})
	for _, iDef := range defs {
		defName, ok := iDef.(string)
		if !ok {
			err = errors.New("'\\u0ADD' property must be an string array")
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
			continue
		}
		if prev[defName] {
			keys := keysSet(prev)
			err = fmt.Errorf("cycle detected with definitions %v", keys)
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
			continue
		}
		tDef := types[defName]
		fDef := fragments[defName]
		if (tDef == nil) && (fDef == nil) {
			err = fmt.Errorf("Unknown definition %s", defName)
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
			continue
		}
		if (tDef != nil) && (fDef != nil) {
			// this should have been previously detected
			err = fmt.Errorf("Internal error duplicate definition %s", defName)
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
			continue
		}
		var def map[string]interface{}
		if tDef != nil {
			next := cloneSet(prev)
			next[defName] = true
			newscope := []string{"types", defName}
			err = doCompose(tDef, types, fragments, next, newscope)
			errs = multierror.Append(errs, err)
			if err != nil {
				continue
			}
			def = tDef.(map[string]interface{})
		} else {
			next := cloneSet(prev)
			next[defName] = true
			newscope := []string{"fragments", defName}
			err = doCompose(fDef, types, fragments, next, newscope)
			errs = multierror.Append(errs, err)
			if err != nil {
				continue
			}
			def = fDef.(map[string]interface{})
		}
		err = mapMerge(dest, def, scope)
		errs = multierror.Append(errs, err)
	}
	err = mapMerge(dest, object, scope)
	errs = multierror.Append(errs, err)
	if errs != nil {
		return errs
	}
	for k := range object {
		delete(object, k)
	}
	for k, v := range dest {
		object[k] = v
	}
	delete(object, ComposeSymbol)
	return nil
}

func removeNilValues(m map[string]interface{}) {
	for k, v := range m {
		if v == nil {
			delete(m, k)
		}
	}
}

func Compose(shell map[string]interface{}) error {
	var typ, frag map[string]interface{}
	var ok bool
	var err, errs error
	t := shell["types"]
	f := shell["fragments"]
	m := shell["main"]
	if t == nil {
		t = make(map[string]interface{})
	}
	if f == nil {
		f = make(map[string]interface{})
	}
	typ, ok = t.(map[string]interface{})
	if !ok {
		err = errors.New("'types' property must be a JSON object")
		err = errorAt(err, nil)
		errs = multierror.Append(errs, err)
	}
	frag, ok = f.(map[string]interface{})
	if !ok {
		err = errors.New("'fragments' property must be a JSON object")
		err = errorAt(err, nil)
		errs = multierror.Append(errs, err)
	}
	if errs != nil {
		return errs
	}
	removeNilValues(typ)
	removeNilValues(frag)
	dupls := intersection(typ, frag)
	if len(dupls) > 0 {
		err = fmt.Errorf("Duplicate keys across 'types' and 'fragments': %v", dupls)
		err = errorAt(err, nil)
		return err
	}
	for k, v := range frag {
		prev := map[string]bool{k: true}
		scope := []string{"fragments", k}
		if PrimitiveTypes[k] {
			err = errors.New("Cannot declare fragment with primitive type name")
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
		} else if v != nil {
			err = doCompose(v, typ, frag, prev, scope)
			errs = multierror.Append(errs, err)
		}
	}
	for k, v := range typ {
		prev := map[string]bool{k: true}
		scope := []string{"types", k}
		if PrimitiveTypes[k] {
			err = errors.New("Cannot declare type with primitive type name")
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
		} else if v != nil {
			err = doCompose(v, typ, frag, prev, scope)
			errs = multierror.Append(errs, err)
		}
	}
	if m != nil {
		prev := map[string]bool{}
		scope := []string{"main"}
		err = doCompose(m, typ, frag, prev, scope)
		errs = multierror.Append(errs, err)
	}
	return errs
}

func intersection(m1 map[string]interface{}, m2 map[string]interface{}) []string {
	var iSect []string
	if len(m1) > len(m2) {
		tmp := m1
		m1 = m2
		m2 = tmp
	}
	for k := range m1 {
		if _, ok := m2[k]; ok {
			iSect = append(iSect, k)
		}
	}
	return iSect
}
