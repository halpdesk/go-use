package slice


// Map updates a slice by applying a function to all members of a slice. The apply function must use and return the same type as the elements of the slice.
func Map[T any](slice []T, apply func(T) T) {
	for i, element := range slice {
		slice[i] = apply(element)
	}
}

// Walk takes a slice of any type and returns another slice. The type of the elements matches the types used in the apply function
func Walk[T any, E any](slice []T, apply func(T) E) []E {
	var updated []E
	for _, element := range slice {
		updated = append(updated, apply(element))
	}
	return updated
}

// Filter reduces the elements of a slice by a condition. The condition is given as a boolean function
func Filter[T comparable](slice []T, filter func(T) bool) []T {
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
