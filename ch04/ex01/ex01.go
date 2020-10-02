package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func popCount(b byte) int {
	count := 0
	for b != 0 {
		count++
		b &= b - 1
	}
	return count
}

func bitDiff(a, b []byte) int {
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		switch {
		case i >= len(a):
			count += popCount(b[i])
		case i >= len(b):
			count += popCount(a[i])
		default:
			count += popCount(a[i] ^ b[i])
		}
	}
	return count
}

func shaBitDiff(a, b []byte) int {
	shaA := sha256.Sum256(a)
	shaB := sha256.Sum256(b)
	return bitDiff(shaA[:], shaB[:])
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Not enough parameters. Add 2 parameters!")
	}
	firstText := []byte(os.Args[1])
	secondText := []byte(os.Args[2])
	fmt.Println(shaBitDiff(firstText, secondText))
}
