package use


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

// func IsZero[T any](obj T) bool {
// 	// For pointers with nil value, or if val is empty time
// 	if val == nil || isZeroTime(val) {
// 		return true
// 	}
// 	v := reflect.ValueOf(val)
// 	switch v.Kind() {
// 	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
// 		return v.Len() == 0
// 	case reflect.Bool:
// 		return !v.Bool()
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		return v.Int() == 0
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
// 		return v.Uint() == 0
// 	case reflect.Float32, reflect.Float64:
// 		return v.Float() == 0
// 	// Are these required if val == nil is checked in the start of this function?
// 	case reflect.Ptr, reflect.Interface:
// 		return v.IsNil()
// 	}
// 	return false
// }

// func Patch[T any](base T, override T) T {
// 	v := reflect.ValueOf(base)
// 	t := reflect.TypeOf(base)
// 	for i := 0; i < v.NumField(); i++ {
// 		key := t.Field(i).Name
// 	}
// }
