package slice

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
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
		Map(test.Original, test.Function)
		assert.Equal(t, test.Expected, test.Original)
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
		Map(test.Original, test.Function)
		assert.Equal(t, test.Expected, test.Original)
	}
}

func TestWalk(t *testing.T) {
	// ints test
	for _, test := range []struct {
		Original []int
		Function func(a int) float32
		Expected []float32
	}{
		{[]int{1, 2, 3}, func(i int) float32 { return float32(0.1) + float32(i) }, []float32{1.1, 2.1, 3.1}},
		{[]int{2, 3, 5}, func(i int) float32 { return float32(0.1) * float32(i) }, []float32{0.2, 0.3, 0.5}},
		{[]int{7, 11, 13}, func(i int) float32 { return float32(0.1) }, []float32{0.1, 0.1, 0.1}},
	} {
		assert.Equal(t, test.Expected, Walk(test.Original, test.Function))
	}
}

func TestFilter(t *testing.T) {
	// ints test
	for _, test := range []struct {
		Original []int
		Function func(a int) bool
		Expected []int
	}{
		{[]int{1, 2, 3}, func(i int) bool { return i > 1 }, []int{2, 3}},
		{[]int{2, 3, 5}, func(i int) bool { return i%2 == 0 }, []int{2}},
	} {
		assert.Equal(t, test.Expected, Filter(test.Original, test.Function))
	}

	// strings test
	for _, test := range []struct {
		Original []string
		Function func(a string) bool
		Expected []string
	}{
		{[]string{"foo", "bar"}, func(str string) bool { return strings.Contains(str, "f") }, []string{"foo"}},
	} {
		assert.Equal(t, test.Expected, Filter(test.Original, test.Function))
	}
}

func TestIndexOf(t *testing.T) {
	// ints test
	for _, test := range []struct {
		Original []int
		Element  int
		Expected int
	}{
		{[]int{1, 2, 3}, 2, 1},
		{[]int{2, 3, 5}, 2, 0},
		{[]int{2, 3, 5}, 7, -1},
	} {
		assert.Equal(t, test.Expected, IndexOf(test.Original, test.Element))
	}

	// strings test
	for _, test := range []struct {
		Original []string
		Element  string
		Expected int
	}{
		{[]string{"foo", "bar", "baz"}, "baz", 2},
		{[]string{"foo", "bar", "baz"}, "faz", -1},
	} {
		assert.Equal(t, test.Expected, IndexOf(test.Original, test.Element))
	}
}
