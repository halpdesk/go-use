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
		Original    []int
		MapFunction func(a int) int
		Expected    []int
	}{
		{[]int{1, 2, 3}, func(i int) int { return i + 1 }, []int{2, 3, 4}},
		{[]int{2, 3, 5}, func(i int) int { return i * i }, []int{4, 9, 25}},
		{[]int{7, 11, 13}, func(i int) int { return 0 }, []int{0, 0, 0}},
	} {
		Map(test.Original, test.MapFunction)
		assert.Equal(t, test.Expected, test.Original)
	}

	// strings test
	for _, test := range []struct {
		Original    []string
		MapFunction func(a string) string
		Expected    []string
	}{
		{[]string{"foo", "bar"}, func(str string) string { return fmt.Sprintf("a%s", str) }, []string{"afoo", "abar"}},
		{[]string{"faz", "baz"}, func(str string) string { return strings.ToUpper(str) }, []string{"FAZ", "BAZ"}},
	} {
		Map(test.Original, test.MapFunction)
		assert.Equal(t, test.Expected, test.Original)
	}
}

func TestWalk(t *testing.T) {
	// ints test
	for _, test := range []struct {
		Original     []int
		WalkFunction func(a int) float32
		Expected     []float32
	}{
		{[]int{1, 2, 3}, func(i int) float32 { return float32(0.1) + float32(i) }, []float32{1.1, 2.1, 3.1}},
		{[]int{2, 3, 5}, func(i int) float32 { return float32(0.1) * float32(i) }, []float32{0.2, 0.3, 0.5}},
		{[]int{7, 11, 13}, func(i int) float32 { return float32(0.1) }, []float32{0.1, 0.1, 0.1}},
	} {
		assert.Equal(t, test.Expected, Walk(test.Original, test.WalkFunction))
	}
}

func TestFilter(t *testing.T) {
	// ints test
	for _, test := range []struct {
		Original       []int
		FilterFunction func(a int) bool
		Expected       []int
	}{
		{[]int{1, 2, 3}, func(i int) bool { return i > 1 }, []int{2, 3}},
		{[]int{2, 3, 5}, func(i int) bool { return i%2 == 0 }, []int{2}},
	} {
		assert.Equal(t, test.Expected, Filter(test.Original, test.FilterFunction))
	}

	// strings test
	for _, test := range []struct {
		Original       []string
		FilterFunction func(a string) bool
		Expected       []string
	}{
		{[]string{"foo", "bar"}, func(str string) bool { return strings.Contains(str, "f") }, []string{"foo"}},
	} {
		assert.Equal(t, test.Expected, Filter(test.Original, test.FilterFunction))
	}
}

func TestRemoveIndex(t *testing.T) {
	// ints test
	for _, test := range []struct {
		Original []int
		Index    int
		Expected []int
	}{
		{[]int{1, 2, 3}, 2, []int{1, 2}},
		{[]int{2, 3, 5}, 0, []int{3, 5}},
	} {
		assert.Equal(t, test.Expected, RemoveIndex(test.Original, test.Index))
	}

	// strings test
	for _, test := range []struct {
		Original []string
		Index    int
		Expected []string
	}{
		{[]string{"foo", "bar"}, 0, []string{"bar"}},
	} {
		assert.Equal(t, test.Expected, RemoveIndex(test.Original, test.Index))
	}
}

func TestRemoveMatching(t *testing.T) {
	// ints test
	for _, test := range []struct {
		Original        []int
		RemovalFunction func(a int) bool
		Expected        []int
	}{
		{[]int{1, 2, 3}, func(i int) bool { return i > 1 }, []int{1}},
		{[]int{2, 3, 5}, func(i int) bool { return i%2 == 0 }, []int{3, 5}},
	} {
		assert.Equal(t, test.Expected, RemoveMatching(test.Original, test.RemovalFunction))
	}

	// strings test
	for _, test := range []struct {
		Original        []string
		RemovalFunction func(a string) bool
		Expected        []string
	}{
		{[]string{"foo", "bar"}, func(str string) bool { return strings.Contains(str, "f") }, []string{"bar"}},
	} {
		assert.Equal(t, test.Expected, RemoveMatching(test.Original, test.RemovalFunction))
	}
}

func TestContains(t *testing.T) {
	// ints test
	for _, test := range []struct {
		Original []int
		Element  int
		Expected bool
	}{
		{[]int{1, 2, 3}, 2, true},
		{[]int{2, 3, 5}, 2, true},
		{[]int{2, 3, 5}, 7, false},
	} {
		assert.Equal(t, test.Expected, Contains(test.Original, test.Element))
	}

	// strings test
	for _, test := range []struct {
		Original []string
		Element  string
		Expected bool
	}{
		{[]string{"foo", "bar", "baz"}, "baz", true},
		{[]string{"foo", "bar", "baz"}, "faz", false},
	} {
		assert.Equal(t, test.Expected, Contains(test.Original, test.Element))
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

func TestChunk(t *testing.T) {
	// ints test
	for _, test := range []struct {
		Original []int
		Size     int
		Expected [][]int
	}{
		{[]int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
		{[]int{1, 2, 3}, 2, [][]int{{1, 2}, {3}}},
		{[]int{1, 2, 3}, 3, [][]int{{1, 2, 3}}},
	} {
		assert.Equal(t, test.Expected, Chunk(test.Original, test.Size))
	}

	// strings test
	for _, test := range []struct {
		Original []string
		Size     int
		Expected [][]string
	}{
		{[]string{"foo", "bar", "baz"}, 1, [][]string{{"foo"}, {"bar"}, {"baz"}}},
		{[]string{"foo", "bar", "baz"}, 2, [][]string{{"foo", "bar"}, {"baz"}}},
		{[]string{"foo", "bar", "baz"}, 3, [][]string{{"foo", "bar", "baz"}}},
	} {
		assert.Equal(t, test.Expected, Chunk(test.Original, test.Size))
	}
}

func TestFlatten(t *testing.T) {
	// ints test
	for _, test := range []struct {
		Original [][]int
		Expected []int
	}{
		{[][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
		{[][]int{{1, 2}, {3}}, []int{1, 2, 3}},
		{[][]int{{1, 2, 3}}, []int{1, 2, 3}},
	} {
		assert.Equal(t, test.Expected, Flatten(test.Original))
	}

	// strings test
	for _, test := range []struct {
		Original [][]string
		Expected []string
	}{
		{[][]string{{"foo"}, {"bar"}, {"baz"}}, []string{"foo", "bar", "baz"}},
		{[][]string{{"foo", "bar"}, {"baz"}}, []string{"foo", "bar", "baz"}},
		{[][]string{{"foo", "bar", "baz"}}, []string{"foo", "bar", "baz"}},
	} {
		assert.Equal(t, test.Expected, Flatten(test.Original))
	}
}
