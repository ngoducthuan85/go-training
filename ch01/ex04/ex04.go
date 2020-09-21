// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	dupFiles := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "", dupFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, dupFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
			fmt.Printf("\t%s\n", dupFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileName string, dupFiles map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if strings.Index(dupFiles[input.Text()], fileName) == -1 {
			dupFiles[input.Text()] += fileName + "; "
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
