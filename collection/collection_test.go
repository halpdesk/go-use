package collection

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCollection(t *testing.T) {
	collection := New[int]()
	assert.Equal(t, Collection[int]{}, collection)
}

func TestCollect(t *testing.T) {
	ints := []int{1, 2, 3}
	collection := Collect(ints)
	value := reflect.ValueOf(collection).FieldByName("items")
	assert.Equal(t, reflect.Slice, value.Kind())
	assert.Equal(t, 3, value.Len())
	for i, v := range ints {
		assert.Equal(t, int64(v), value.Index(i).Int())
	}
}

func TestAdd(t *testing.T) {
	ints := []int{3, 2, 1}
	collection := Collection[int]{items: ints}
	collection.Add(int(5))
	collection.Add(int(7))
	ints = append(ints, 5, 7)
	value := reflect.ValueOf(collection).FieldByName("items")
	assert.Equal(t, reflect.Slice, value.Kind())
	assert.Equal(t, 5, value.Len())
	for i, v := range ints {
		assert.Equal(t, int64(v), value.Index(i).Int())
	}
}

func TestItems(t *testing.T) {
	ints := []int{1, 2, 3}
	collection := Collection[int]{items: ints}
	assert.Equal(t, ints, collection.Items())
}
