package kind

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsEmptyValue(t *testing.T) {
	assert.Equal(t, true, isEmptyValue(int(0)))
	assert.Equal(t, false, isEmptyValue(int(1)))
	assert.Equal(t, true, isEmptyValue(string("")))
	assert.Equal(t, false, isEmptyValue(string("foo")))
	assert.Equal(t, true, isEmptyValue(float64(0)))
	assert.Equal(t, false, isEmptyValue(float64(1)))
	assert.Equal(t, true, isEmptyValue(0))
	assert.Equal(t, false, isEmptyValue(1))
	assert.Equal(t, true, isEmptyValue(""))
	assert.Equal(t, false, isEmptyValue("foo"))
	assert.Equal(t, true, isEmptyValue(false))
	assert.Equal(t, true, isEmptyValue(time.Time{}))
	assert.Equal(t, false, isEmptyValue(time.Now()))
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
	assert.Equal(t, true, IsZero(false))
	assert.Equal(t, false, IsZero(true))
	assert.Equal(t, true, IsZero(time.Time{}))
	assert.Equal(t, false, IsZero(time.Now()))
}

func TestPatch(t *testing.T) {
	type TestStruct struct {
		ID        uint64
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
		{TestStruct{ID: 1, Name: "A", Date: time.Now(), DeletedAt: &now}, TestStruct{ID: 2}, TestStruct{ID: 2, Name: "A", Date: time.Now(), DeletedAt: &now}},
		{TestStruct{DeletedAt: &now}, TestStruct{Name: "B"}, TestStruct{Name: "B", DeletedAt: &now}},
		{TestStruct{}, TestStruct{ID: 4, Name: "C", Date: time.Now(), DeletedAt: &now}, TestStruct{ID: 4, Name: "C", Date: time.Now(), DeletedAt: &now}},
	} {
		result, err := Patch(test.Test1, test.Test2)
		require.NoError(t, err)
		assert.Equal(t, test.Expected.ID, result.ID)
		assert.Equal(t, test.Expected.Name, result.Name)
		// Time can be a bit different after marshalling: https://github.com/golang/go/issues/22957
		assert.Equal(t, test.Expected.Date.Unix(), result.Date.Unix())
		if test.Expected.DeletedAt != nil {
			assert.Equal(t, test.Expected.DeletedAt.Unix(), result.DeletedAt.Unix())
		} else {
			assert.Equal(t, nil, result.DeletedAt)
		}
	}
}

func TestPatchError(t *testing.T) {
	unmarshalable := func() int { return 1 }
	_, err := Patch(unmarshalable, unmarshalable)
	require.Error(t, err)
}
