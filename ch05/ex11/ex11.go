// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},
	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":      {"discrete math": true},
	"databases":            {"data structures": true},
	"discrete math":        {"intro to programming": true},
	"formal languages":     {"discrete math": true},
	"networks":             {"operating system": true},
	"operating system":     {"data structures": true, "computer organization": true},
	"programming language": {"data structures": true, "computer organization": true},
	"linear algebra":       {"calculus": true},
}

//!-table

//!+main
func main() {
	result, loopIsFound := topoSort(prereqs)
	if loopIsFound {
		fmt.Println("Loop is found!")
		return
	}
	for i, course := range result {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) ([]string, bool) {
	var order []string
	var loopIsFound = false
	seen := make(map[string]int) // 0: 来ていない、 1: 来ているがvisitしていない、 2: visitした
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item := range items {
			if seen[item] == 0 {
				seen[item] = 1
				visitAll(m[item])
				order = append(order, item)
				seen[item] = 2
			} else if seen[item] == 1 {
				loopIsFound = true
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}
	visitAll(keys)
	return order, loopIsFound
}

//!-main
