// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {

	fmt.Println(url)

	tokens <- struct{}{}
	list, err := Extract(url)
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	depth := flag.Int("depth", 0, "depth")
	flag.Parse()

	type Links struct {
		List  []string
		Depth int
	}

	type Link struct {
		URL   string
		Depth int
	}

	worklist := make(chan *Links)   // lists of URLs, may have duplicates
	unseenLinks := make(chan *Link) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- &Links{flag.Args(), 0} }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link.URL)
				go func(depth int) {
					worklist <- &Links{foundLinks, depth + 1}
				}(link.Depth)
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		if list.Depth > *depth {
			continue
		}
		for _, link := range list.List {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- &Link{link, list.Depth}
			}
		}
	}
}

//!-

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
