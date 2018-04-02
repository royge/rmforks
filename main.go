package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/go-github/github"
	"github.com/royge/rmforks/config"
	"github.com/royge/rmforks/repo"
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
	ctx := context.Background()

	svc := repo.Service{
		User:     cfg.Username,
		Client:   github.NewClient(tc),
		Exclude:  cfg.Exclude,
		Timeout:  cfg.Timeout,
		Register: make(chan *github.Repository),
	}

	go svc.Fetch(ctx)
	go svc.Delete(ctx)

	<-svc.Done()
	os.Exit(0)
}
