package jsonstructure

import (
	"encoding/json"
	"reflect"

	"github.com/mitchellh/hashstructure"
	"github.com/shopspring/decimal"
)

type sethash map[uint64][]interface{}

func createSet() sethash {
	return make(map[uint64][]interface{})
}

func cleanJSON(val interface{}) interface{} {
	switch tval := val.(type) {
	case json.Number:
		dec, _ := decimal.NewFromString(tval.String())
		return dec.String()
	case map[string]interface{}:
		res := make(map[string]interface{})
		for k, v := range tval {
			res[k] = cleanJSON(v)
		}
		return res
	case []interface{}:
		res := make([]interface{}, 0, len(tval))
		for i, v := range tval {
			res[i] = cleanJSON(v)
		}
		return res
	default:
		return val
	}
}

// PutIfAbsent inserts an element into the set if it doesn't exist
func (s sethash) PutIfAbsent(val interface{}) (bool, error) {
	val = cleanJSON(val)
	hash, err := hashstructure.Hash(val, nil)
	if err != nil {
		return false, err
	}
	slice, ok := s[hash]
	if !ok {
		slice = []interface{}{val}
		s[hash] = slice
		return true, nil
	}
	for _, v := range slice {
		if reflect.DeepEqual(v, val) {
			return false, nil
		}
	}
	slice = append(slice, val)
	s[hash] = slice
	return true, nil
}

// Has tests whether an element is a member of the set
func (s sethash) Has(val interface{}) (bool, error) {
	val = cleanJSON(val)
	hash, err := hashstructure.Hash(val, nil)
	if err != nil {
		return false, err
	}
	slice, ok := s[hash]
	if !ok {
		return false, nil
	}
	for _, v := range slice {
		if reflect.DeepEqual(v, val) {
			return true, nil
		}
	}
	return false, nil
}
