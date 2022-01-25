package slice

func ForEach[T any](slice []T, call func(int, T)) {
	for i, element := range slice {
		call(i, element)
	}
}

func ForEachWithErr[T any](slice []T, call func(int, T) error) error {
	var err error
	for i, element := range slice {
		err = call(i, element)
		if err != nil {
			return err
		}
	}
	return nil
}

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

func RemoveIndex[T comparable](slice []T, i int) []T {
	var empty T
	copy(slice[i:], slice[i+1:]) // Shift a[i+1:] left one index.
	slice[len(slice)-1] = empty  // Erase last element (write zero value).
	return slice[:len(slice)-1] // Truncate slice.
}

func RemoveMatching[T comparable](slice []T, removal func(T) bool) []T {
	var updated []T
	for _, element := range slice {
		if !removal(element) {
			updated = append(updated, element)
		}
	}
	return updated
}

func Contains[T comparable](slice []T, item T) bool {
    for _, element := range slice {
        if element == item {
            return true
        }
    }
    return false
}

func IndexOf[T comparable](slice []T, item T) int {
    for i, element := range slice {
        if element == item {
            return i
        }
    }
    return -1
}

func Chunk[T any](slice []T, size int) (chunks [][]T) {
	for size < len(slice) {
		slice, chunks = slice[size:], append(chunks, slice[0:size:size])
	}
	return append(chunks, slice)
}

func Flatten[T any](slice [][]T) []T {
	var flattened []T
	for i := range slice {
		for j := range slice[i] {
			flattened = append(flattened, slice[i][j])
		}
	}
	return flattened
}
