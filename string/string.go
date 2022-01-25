package string

import (
	"strings"
)

type AnyString interface {
    ~string
    // ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string | ~rune
}

const alphanum = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsAlphanumeric[T AnyString](a T) bool {
	for _, r := range a {
		if !strings.ContainsRune(alphanum, r) {
			return false
		}
	}
	return true
}
