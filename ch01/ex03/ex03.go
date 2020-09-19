package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start1 := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// fmt.Println(s)
	fmt.Printf("%.8fs elapsed - 非効率手法\n", time.Since(start1).Seconds())

	start2 := time.Now()
	// fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf("%.8fs elapsed - 効率手法\n", time.Since(start2).Seconds())
}
