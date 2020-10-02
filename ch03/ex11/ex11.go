// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaBuf(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaBuf(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	count := 3 - n%3
	for i := 0; i < n; i++ {
		fmt.Fprintf(&buf, "%c", s[i])
		count++
		if (count%3 == 0) && (i != n-1) {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

//!-
