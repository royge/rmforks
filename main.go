package main

import (
	"github.com/google/go-github/github"
	"github.com/r0y3/rmforks/config"
	"github.com/r0y3/rmforks/reposervice"
	"golang.org/x/oauth2"
)

func waitThenQuit(svc *reposervice.RepoService) {
	done := false
	for {
		select {
		case done = <-svc.Done():
			if done {
				break
			}
		default:
			continue
		}
		if done {
			break
		}
	}
}

func main() {
	cfg, err := config.GetConfig("config.json")

	if err != nil {
		panic(err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.AccessToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	svc := reposervice.RepoService{
		User:    cfg.Username,
		Client:  github.NewClient(tc),
		Exclude: cfg.Exclude,
		Timeout: cfg.Timeout,
	}

	go svc.Fetch()
	go svc.Delete()

	waitThenQuit(svc)
}
