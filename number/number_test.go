package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSum(t *testing.T) {
	type SumTest[T Number] struct {
		Slice    []T
		Expected T
	}
	for _, test := range []SumTest[int]{
		{[]int{1, 2, 3}, int(6)},
		{[]int{-1, 2, -3}, int(-2)},
	} {
		assert.Equal(t, test.Expected, Sum(test.Slice))
	}
}

func TestReduce(t *testing.T) {
	type ReduceTest[T Number] struct {
		Slice    []T
		Reducer  func(prev, next T) T
		Expected T
	}
	for _, test := range []ReduceTest[int]{
		{[]int{1, 2, 3}, func(p, n int) int { return p + n }, int(6)},
		{[]int{2, 3, 5}, func(p, n int) int { 
			if (p==0) {
				return 1 * n
			} else {
				return p * n 
			}
		}, int(30)},
	} {
		assert.Equal(t, test.Expected, Reduce(test.Slice, test.Reducer))
	}
}
