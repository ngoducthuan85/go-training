package main

import "fmt"

func Rotate_old(arr []int, k int) []int {
	if k < 0 || len(arr) == 0 {
		return arr
	}
	r := len(arr) - k%len(arr)
	arr = append(arr[r:], arr[:r]...)
	return arr
}

// Copied from @orisano
func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

// Rotate is copied from @orisano
func Rotate(a []int, r int) {
	L := len(a)
	g := gcd(L, r)
	for i := 0; i < g; i++ {
		x := a[i]
		for j := (i + r) % L; j != i; j = (j + r) % L {
			x, a[j] = a[j], x
		}
		a[i] = x
	}
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr = Rotate_old(arr, 4)
	Rotate(arr, 4)
	fmt.Printf("%v\n", arr)
}
