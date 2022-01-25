package kind

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/RaveNoX/go-jsonmerge"
)

func IsZero[T comparable](obj T) bool {
	var zeroVal T
	return obj == zeroVal
}

func Patch[T any](base T, patch T) (T, error) {
	var patched T
	baseJSON, err := json.Marshal(base)
	if err != nil {
		return patched, fmt.Errorf("could not marshal base: %s", err.Error())
	}
	patchJSON, err := json.Marshal(patch)
	if err != nil {
		return patched, fmt.Errorf("could not marshal patch: %s", err.Error())
	}
	// Delete all empty fields
	patchWithoutEmpty := make(map[string]any)
	err = json.Unmarshal(patchJSON, &patchWithoutEmpty)
	if err != nil {
		return patched, fmt.Errorf("could not marshal to patch to map to clean empty: %s", err.Error())
	}
	for key, value := range patchWithoutEmpty {
		// Cannot use IsZero because of types are unknown (interface{}) in map
		if isEmptyValue(value) {
			delete(patchWithoutEmpty, key)
		}
	}
	// Marshal to new JSON
	patchJSON, err = json.Marshal(patchWithoutEmpty)
	if err != nil {
		return patched, fmt.Errorf("could not marshal patch with new values to JSON: %s", err.Error())
	}
	// Combine with jsonmerge
	combined, _, err := jsonmerge.MergeBytes(baseJSON, patchJSON)
	if err != nil {
		return patched, fmt.Errorf("could not marshal to combined JSON: %s", err.Error())
	}
	err = json.Unmarshal(combined, &patched)
	if err != nil {
		return patched, fmt.Errorf("could not unmarshal to model: %s", err.Error())
	}
	return patched, nil
}

func isEmptyValue(val any) bool {
	isZeroTime := func(val any) bool {
		// If val is marshalled to json, this is what zero time looks like
		if val == "0001-01-01T00:00:00Z" {
			return true
		}
		if fmt.Sprintf("%T", val) == "time.Time" {
			t := val.(time.Time)
			return t == time.Time{}
		}
		return false
	}
	if isZeroTime(val) {
		return true
	}
	// For pointers with nil value, or if val is empty time
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	// Are these required if val == nil is checked in the start of this function?
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	}
	return false
}
