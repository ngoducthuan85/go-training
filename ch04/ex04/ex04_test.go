// Copied from @orisano
package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	ts := []struct {
		s        []int
		r        int
		expected []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5},
			1,
			[]int{5, 0, 1, 2, 3, 4},
		},
		{
			[]int{0, 1, 2, 3, 4, 5},
			1,
			[]int{5, 0, 1, 2, 3, 4},
		},
		{
			[]int{0, 1, 2, 3, 4, 5},
			2,
			[]int{4, 5, 0, 1, 2, 3},
		},
		{
			[]int{0, 1, 2, 3, 4, 5},
			3,
			[]int{3, 4, 5, 0, 1, 2},
		},
		{
			[]int{1, 9, 2, 3, 5, 1, 4, 6},
			2,
			[]int{4, 6, 1, 9, 2, 3, 5, 1},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			4,
			[]int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
		},
	}

	for _, tc := range ts {
		// tc.s = Rotate_old(tc.s, tc.r)
		Rotate(tc.s, tc.r)
		if !reflect.DeepEqual(tc.s, tc.expected) {
			t.Errorf("unexpected slice. expected: %v, but got: %v", tc.expected, tc.s)
		}
	}
}
