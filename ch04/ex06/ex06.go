package main

import (
	"fmt"
	"os"
)

func compressSpace(b []byte) []byte {
	return []byte{}
}

func main() {
	b := []byte(os.Args[1])
	fmt.Printf("%x\n", sha(b))
}
