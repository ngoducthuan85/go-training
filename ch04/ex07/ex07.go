package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func rotateLeftBytes(b []byte, numBytes int) {
	reverseBytes(b[:numBytes])
	reverseBytes(b[numBytes:])
	reverseBytes(b)
}

func reverseUTF8(b []byte) {
	n := len(b)
	_, numBytes := utf8.DecodeRune(b)
	for n > 0 {
		rotateLeftBytes(b[:n], numBytes)
		n -= numBytes
		_, numBytes = utf8.DecodeRune(b[:n])
	}
}

func main() {
	s := "Hello, 世界！"
	b := []byte(s)
	reverseUTF8(b)
	s1 := string(b)
	fmt.Println(s, "\t", s1)
}
