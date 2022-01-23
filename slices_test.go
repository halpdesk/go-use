package use

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestIsZero(t *testing.T) {
	// DOES NOT WORK AT THE MOMENT
	// "cannot use generic type Thing[T any] without instantiation" in test case struct declaration

	// type Thing[T any] *T

	// for _, test := range []struct {
	// 	Test 	 Thing
	// 	Expected bool
	// }{
	// 	{int(0), true},
	// 	{string(""), true},
	// 	{string("foo"), false},
	// }{
	// 	assert.Equal(t, test.Expected, IsZero(test.Test))
	// }

	assert.Equal(t, true, IsZero(int(0)))
	assert.Equal(t, false, IsZero(int(1)))
	assert.Equal(t, true, IsZero(string("")))
	assert.Equal(t, false, IsZero(string("foo")))
	assert.Equal(t, true, IsZero(float64(0)))
	assert.Equal(t, false, IsZero(float64(1)))
	assert.Equal(t, true, IsZero(0))
	assert.Equal(t, false, IsZero(1))
	assert.Equal(t, true, IsZero(""))
	assert.Equal(t, false, IsZero("foo"))
}

func TestPatch(t *testing.T) {
	type TestStruct struct {
		ID        uint64
		Num       int
		Name      string
		Date      time.Time
		DeletedAt *time.Time
	}
	now := time.Now()
	for _, test := range []struct {
		Test1    TestStruct
		Test2    TestStruct
		Expected TestStruct
	}{
		{TestStruct{ID: 1, Num: 1, Name: "A", Date: time.Now()}, TestStruct{ID: 2}, TestStruct{ID: 2, Num: 1, Name: "A", Date: time.Now()}},
		{TestStruct{DeletedAt: &now}, TestStruct{ID: 3}, TestStruct{ID: 3, DeletedAt: &now}},
	} {
		result, err := Patch(test.Test1, test.Test2)
		require.NoError(t, err)
		assert.Equal(t, test.Expected.ID, result.ID)
		assert.Equal(t, test.Expected.Num, result.Num)
		assert.Equal(t, test.Expected.Name, result.Name)

		// Time can be a bit different after marshalling: https://github.com/golang/go/issues/22957
		assert.Equal(t, test.Expected.Date.Unix(), result.Date.Unix())
		if test.Expected.DeletedAt != nil {
			assert.Equal(t, test.Expected.DeletedAt.Unix(), result.DeletedAt.Unix())
		}
	}

}
