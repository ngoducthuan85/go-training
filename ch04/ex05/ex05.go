package main

import (
	"fmt"
)

func unique(strArr []string) []string {
	w := 0
	for _, s := range strArr {
		if strArr[w] == s {
			continue
		}
		w++
		strArr[w] = s
	}
	return strArr[:w+1]
}

func main() {
	arr := []string{"aa", "bb", "bb", "aaaa", "dd"}
	fmt.Printf("%v\n", unique(arr))
}
