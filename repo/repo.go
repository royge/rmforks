package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
	"github.com/royge/rmforks/strings"
)

var done = make(chan bool)

var opt = &github.RepositoryListOptions{
	ListOptions: github.ListOptions{PerPage: 10},
}

// Service struct type
type Service struct {
	User     string
	Client   *github.Client
	Exclude  []string
	Timeout  time.Duration
	Register chan *github.Repository
}

// Done returns channel of bool indicating the status of the deletion.
func (svc *Service) Done() chan bool {
	return done
}

// Fetch retrieves all Github repositories of the user.
func (svc *Service) Fetch(ctx context.Context) {
	for {
		repos, resp, err := svc.Client.Repositories.List(ctx, svc.User, opt)
		if err != nil {
			close(svc.Register)
			panic(err)
		}
		for _, repo := range repos {
			select {
			case svc.Register <- repo:
			}
		}
		if resp.NextPage == 0 {
			close(svc.Register)
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}
}

// Delete removes the repository if it is a forked and not in the whitelist.
func (svc *Service) Delete(ctx context.Context) {
	for {
		select {
		case repo := <-svc.Register:
			if *repo.Fork == true && !strings.Contains(svc.Exclude, *repo.Name) {
				fmt.Println("Deleting ", *repo.Name)
				_, err := svc.Client.Repositories.Delete(ctx, svc.User, *repo.Name)
				if err != nil {
					panic(err)
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
