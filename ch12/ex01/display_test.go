package display

import (
	"testing"
)

// This test ensures that the program terminates without crashing.
func Test(t *testing.T) {
	// a map that contains itself
	type M map[string]M
	m := make(M)
	m[""] = m
	if false {
		Display("m", m)
		// Output:
		// Display m (display.M):
		// ...stuck, no output...
	}
}

func TestMapKeys(t *testing.T) {
	sm := map[struct{ x int }]int{
		{1}: 10,
		{2}: 20,
		{3}: 30,
		{4}: 40,
	}
	Display("sm", sm)
	// Output:
	// Display sm (map[struct { x int }]int):
	// sm[{x: 1}] = 10
	// sm[{x: 2}] = 20
	// sm[{x: 3}] = 30
	// sm[{x: 4}] = 40

	am := map[[3]int]int{
		{1, 2, 3}: 6,
		{4, 5, 6}: 15,
		{7, 8, 9}: 24,
	}
	Display("am", am)
	// Output:
	// Display am (map[[3]int]int):
	// am[{1, 2, 3}] = 6
	// am[{4, 5, 6}] = 15
	// am[{7, 8, 9}] = 24
}