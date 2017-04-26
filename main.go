package main

import (
	"flag"
	"os"

	"github.com/google/go-github/github"
	"github.com/r0y3/rmforks/config"
	"github.com/r0y3/rmforks/repo"
	"golang.org/x/oauth2"
)

var conf = flag.String("conf", "config.json", "Configuration file.")

func main() {
	flag.Parse()

	cfg, err := config.GetConfig(*conf)

	if err != nil {
		panic(err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.AccessToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	svc := repo.Service{
		User:    cfg.Username,
		Client:  github.NewClient(tc),
		Exclude: cfg.Exclude,
		Timeout: cfg.Timeout,
	}

	go svc.Fetch()
	go svc.Delete()

	<-svc.Done()
	os.Exit(0)
}
