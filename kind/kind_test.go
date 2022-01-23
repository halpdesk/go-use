package kind

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
