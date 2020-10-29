package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/go-github/v32/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	// repos, _, err := client.Repositories.List(ctx, "", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// for _, rep := range repos {
	// 	fmt.Printf("%v\n", rep)
	// }

	var user, repo string
	flag.StringVar(&user, "user", "ngoducthuan85", "github user")
	flag.StringVar(&repo, "repo", "go-training", "github repository")
	flag.Parse()

	cmds := flag.Args()
	switch cmds[0] {
	case "get":
		if len(cmds) >= 2 {
			number, err := strconv.Atoi(cmds[1])
			if err != nil {
				log.Fatalln(err)
			}
			item, _, err := client.Issues.Get(ctx, user, repo, number)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%v %s %s\n", *item.Number, *item.User.Login, *item.Title)
		} else {
			issues, _, err := client.Issues.ListByRepo(ctx, user, repo, nil)
			if err != nil {
				log.Fatalln(err)
			}
			for _, item := range issues {
				fmt.Printf("%v %s %s\n", *item.Number, *item.User.Login, *item.Title)
			}
		}
	case "create":
		if len(cmds) >= 3 {
			title := cmds[1]
			body := cmds[2]
			item, _, err := client.Issues.Create(ctx, user, repo, &github.IssueRequest{Title: &title, Body: &body})
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%v %s %s\n", *item.Number, *item.User.Login, *item.Title)
		}
	case "edit":
		if len(cmds) >= 4 {
			number, err := strconv.Atoi(cmds[1])
			if err != nil {
				log.Fatalln(err)
			}
			title := cmds[2]
			body := cmds[3]
			item, _, err := client.Issues.Edit(ctx, user, repo, number, &github.IssueRequest{Title: &title, Body: &body})
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%v %s %s\n", *item.Number, *item.User.Login, *item.Title)
		}
	}
}
