// Copied from @orisano
package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	ts := []struct {
		s        string
		expected string
	}{
		{
			"Hello,世界",
			"Hello,世界",
		},
		{
			"Hello, 世界",
			"Hello, 世界",
		},
		{
			" Hel lo, 世界",
			" Hel lo, 世界",
		},
		{
			"  Hel  lo, 世界",
			" Hel lo, 世界",
		},
		{
			"  　　Hel  lo, 　　世界　",
			" Hel lo, 世界 ",
		},
	}

	for _, tc := range ts {
		b := []byte(tc.s)
		b1 := CompressSpaceInsideArray(b)
		s1 := string(b1)
		if !reflect.DeepEqual(s1, tc.expected) {
			t.Errorf("Unexpected string. expected: '%v', but got: '%v'", tc.expected, s1)
		}
	}
}
