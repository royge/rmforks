package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/r0y3/rmforks/config"
	"golang.org/x/oauth2"
	"os"
	"sort"
)

type cfg struct {
	Username    string
	AccessToken string
	Exclude     []string
}

func contains(hayStack []string, target string) bool {
	sort.Strings(hayStack)
	i := sort.SearchStrings(hayStack, target)
	return i < len(hayStack) && hayStack[i] == target
}

func main() {
	c := &cfg{}
	config.ConfigFunc(func(configFile string) error {
		file, err := os.Open(configFile)
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(file)
		err = decoder.Decode(c)

		if err != nil {
			return err
		}

		return nil
	})("config.json")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.AccessToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	var allRepos []*github.Repository
	var opt = &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	for {
		repos, resp, err := client.Repositories.List(c.Username, opt)
		if err != nil {
			panic(err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}

	for _, repo := range allRepos {
		func(repo *github.Repository) {
			if *repo.Fork == true && !contains(c.Exclude, *repo.Name) {
				fmt.Println("Deleting ", *repo.Name)
				_, err := client.Repositories.Delete(c.Username, *repo.Name)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("Done.")
			}
		}(repo)
	}
}
