package main

import "fmt"

func rotate(arr []int, k int) []int {
	if k < 0 || len(arr) == 0 {
		return arr
	}
	r := len(arr) - k%len(arr)
	arr = append(arr[r:], arr[:r]...)

	return arr
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	newArr := rotate(arr, 2)
	fmt.Printf("%v\n", newArr)
}
