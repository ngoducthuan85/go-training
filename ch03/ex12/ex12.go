package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Not enough parameters. Add 2 strings to the parameters!")
	}
	firstText := os.Args[1]
	secondText := os.Args[2]
	fmt.Printf("Is Anagram? %v\n", anagram(firstText, secondText))
}

func anagram(s1, s2 string) bool {
	return sortStr(s1) == sortStr(s2)
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
