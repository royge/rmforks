package main

import (
	"github.com/google/go-github/github"
	"github.com/r0y3/rmforks/config"
	"golang.org/x/oauth2"
)

func main() {
	configFunc := config.ConfigFunc(func(configFile string) (config.Config, error) {
		// TODO: Read config file.
		return nil, error.New("Blahh")
	})

	authFunc := config.AuthFunc(func(config config.Config) (interface{}, error) {
		// TODO: Do the decoding
		return nil, error.New("Blahh")
	})

	cred := &struct {
		Username    string
		AccessToken string
	}{}

	var err error
	cred, err = authFunc.GetCredentials(configFunc)

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cred.AccessToken},
	)

	client := github.NewClient(oauth2.NoContext, ts)

	repos, _, err := client.Repositories.List(cred.Username, nil)
}
