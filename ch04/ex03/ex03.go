package main

import "fmt"

func reverse(arr *[10]int) {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(&arr)
	fmt.Printf("%v\n", arr)
}
