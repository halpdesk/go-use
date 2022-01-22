package use

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceWalk(t *testing.T) {
	// ints test
	for _, test := range []struct {
		Original []int
		Function func(a int) int
		Expected []int
	}{
		{[]int{1, 2, 3}, func(i int) int { return i + 1 }, []int{2, 3, 4}},
		{[]int{2, 3, 5}, func(i int) int { return i * i }, []int{4, 9, 25}},
		{[]int{7, 11, 13}, func(i int) int { return 0 }, []int{0, 0, 0}},
	} {
		assert.Equal(t, test.Expected, SliceWalk(test.Original, test.Function))
	}

	// strings test
	for _, test := range []struct {
		Original []string
		Function func(a string) string
		Expected []string
	}{
		{[]string{"foo", "bar"}, func(str string) string { return fmt.Sprintf("a%s", str) }, []string{"afoo", "abar"}},
		{[]string{"faz", "baz"}, func(str string) string { return strings.ToUpper(str) }, []string{"FAZ", "BAZ"}},
	} {
		assert.Equal(t, test.Expected, SliceWalk(test.Original, test.Function))
	}
}

func TestSliceFilter(t *testing.T) {
	// ints test
	for _, test := range []struct {
		Original []int
		Function func(a int) bool
		Expected []int
	}{
		{[]int{1, 2, 3}, func(i int) bool { return i > 1 }, []int{2, 3}},
		{[]int{2, 3, 5}, func(i int) bool { return i%2 == 0 }, []int{2}},
	} {
		assert.Equal(t, test.Expected, SliceFilter(test.Original, test.Function))
	}

	// strings test
	for _, test := range []struct {
		Original []string
		Function func(a string) bool
		Expected []string
	}{
		{[]string{"foo", "bar"}, func(str string) bool { return strings.Contains(str, "f") }, []string{"foo"}},
	} {
		assert.Equal(t, test.Expected, SliceFilter(test.Original, test.Function))
	}
}
