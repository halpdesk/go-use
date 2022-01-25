package number

type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type Number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func Sum[T Number](slice []T) T {
	var sum T
	for _, num := range slice {
		sum = sum + num
	}
	return sum
}

// Reduce updates a slice by applying a function to all members of a slice. The apply function must use and return the same type as the elements of the slice.
func Reduce[T Number](slice []T, reduce func(T, T) T) T {
	var sum T
	for _, num := range slice {
		sum = reduce(sum, num)
	}
	return sum
}
