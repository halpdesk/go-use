package string

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestIsAlphaNumeric(t *testing.T) {
	type AlphaNumTest[T AnyString] struct {
		String   T
		Expected bool
	}
	for _, test := range []AlphaNumTest[string]{
		{"abcauiria7841", true},
		{"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", true},
		{"źasd0", false},
		{"12f€12d", false},
		{"aaaä", false},
	} {
		assert.Equal(t, test.Expected, IsAlphanumeric(test.String))
	}
}
