package use

import (
	"fmt"
	"reflect"
	"encoding/json"
	"github.com/RaveNoX/go-jsonmerge"
)


// SliceWalk takes a slice (of type []T) and applies a function to all its elements
func SliceWalk[T any](slice []T, apply func(T) T) []T {
	var updated []T
	for _, element := range slice {
		updated = append(updated, apply(element))
	}
	return updated
}

func SliceFilter[T comparable](slice []T, filter func(T) bool) []T {
    var updated []T
	for _, element := range slice {
		if filter(element) {
			updated = append(updated, element)
		}
	}
	return updated
}

func IndexOf[T comparable](slice []T, item T) int {
    for i, element := range slice {
        if element == item {
            return i
        }
    }
    return -1
}

func IsZero[T comparable](obj T) bool {
	var zeroVal T
	return obj == zeroVal
}

func Patch[T any](base T, patch T) (T, error) {
	baseJSON, err := json.Marshal(base)
	if err != nil {
		return base, fmt.Errorf("could not marshal base: %s", err.Error())
	}
	patchJSON, err := json.Marshal(patch)
	if err != nil {
		return base, fmt.Errorf("could not marshal patch: %s", err.Error())
	}

	// Delete all empty fields
	patchWithoutEmpty := make(map[string]interface{})
	err = json.Unmarshal(patchJSON, &patchWithoutEmpty)
	if err != nil {
		return base, fmt.Errorf("could not marshal to patch to map to clean empty: %s", err.Error())
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
		return base, fmt.Errorf("could not marshal patch with new values to JSON: %s", err.Error())
	}

	// Combine with jsonmerge
	combined, _, err := jsonmerge.MergeBytes(baseJSON, patchJSON)
	if err != nil {
		return base, fmt.Errorf("could not marshal to combined JSON: %s", err.Error())
	}

	var patched T
	err = json.Unmarshal(combined, &patched)
	if err != nil {
		return base, fmt.Errorf("could not unmarshal to model: %s", err.Error())
	}

	return patched, nil
}

func isEmptyValue(val interface{}) bool {
	isZeroTime := func(val interface{}) bool {
		return val == "0001-01-01T00:00:00Z"
	}
	// For pointers with nil value, or if val is empty time
	if val == nil || isZeroTime(val) {
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
