package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	countMap := map[string]int{}

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		countMap[word]++
	}

	fmt.Printf("text\tcount\n")
	for word, n := range countMap {
		fmt.Printf("%16s\t%d\n", word, n)
	}
}
