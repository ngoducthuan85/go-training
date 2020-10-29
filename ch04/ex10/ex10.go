// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"gopl.io/ch4/github"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	issuesMap := map[string][]*github.Issue{}
	for _, item := range result.Items {
		// fmt.Printf("#%-5d %9.9s %.55s %v  %v\n",
		// 	item.Number, item.User.Login, item.Title, item.CreatedAt, time.Duration.Hours(time.Since(item.CreatedAt)) > 24*30)
		t := time.Duration.Hours(time.Since(item.CreatedAt))
		label := "一年以上"
		if t < 24*30 {
			label = "一か月未満"
		} else if t < 24*365 {
			label = "一年未満"
		}
		issuesMap[label] = append(issuesMap[label], item)
	}
	// ("一か月未満", "一年未満", "一年以上")
	keys := reflect.ValueOf(issuesMap).MapKeys()

	for _, l := range keys {
		label := fmt.Sprintf("%s", l)
		fmt.Printf("\n%s\n", label)
		for _, item := range issuesMap[string(label)] {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}

}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/