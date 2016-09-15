package reposervice

import (
	"fmt"
	"github.com/google/go-github/github"
	"github.com/r0y3/rmforks/stringutil"
	"time"
)

var register chan *github.Repository = make(chan *github.Repository)
var done chan bool = make(chan bool)

var opt = &github.RepositoryListOptions{
	ListOptions: github.ListOptions{PerPage: 10},
}

type RepoService struct {
	User    string
	Client  *github.Client
	Exclude []string
	Timeout time.Duration
}

func (svc *RepoService) Done() chan bool {
	return done
}

func (svc *RepoService) Fetch() {
	for {
		repos, resp, err := svc.Client.Repositories.List(svc.User, opt)
		if err != nil {
			panic(err)
		}
		for _, repo := range repos {
			select {
			case register <- repo:
			}
		}
		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}
}

func (svc *RepoService) Delete() {
	for {
		select {
		case repo := <-register:
			if *repo.Fork == true && !stringutil.Contains(svc.Exclude, *repo.Name) {
				fmt.Println("Deleting ", *repo.Name)
				_, err := svc.Client.Repositories.Delete(svc.User, *repo.Name)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("Done.")
			}
		case <-time.After(time.Second * svc.Timeout):
			select {
			case done <- true:
			}
		}
	}
}
