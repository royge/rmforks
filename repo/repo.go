package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
	"github.com/royge/rmforks/strings"
)

var register = make(chan *github.Repository)
var done = make(chan bool)

var opt = &github.RepositoryListOptions{
	ListOptions: github.ListOptions{PerPage: 10},
}

// Service struct type
type Service struct {
	User    string
	Client  *github.Client
	Exclude []string
	Timeout time.Duration
}

// Done returns channel of bool indicating the status of the deletion.
func (svc *Service) Done() chan bool {
	return done
}

// Fetch retrieves all Github repositories of the user.
func (svc *Service) Fetch(ctx context.Context) error {
	for {
		repos, resp, err := svc.Client.Repositories.List(ctx, svc.User, opt)
		if err != nil {
			return err
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

	return nil
}

// Delete removes the repository if it is a forked and not in the whitelist.
func (svc *Service) Delete(ctx context.Context) error {
	for {
		select {
		case repo := <-register:
			if *repo.Fork == true && !strings.Contains(svc.Exclude, *repo.Name) {
				fmt.Println("Deleting ", *repo.Name)
				_, err := svc.Client.Repositories.Delete(ctx, svc.User, *repo.Name)
				if err != nil {
					return err
				}
				fmt.Println("Done.")
				return nil
			}
		case <-time.After(time.Second * svc.Timeout):
			select {
			case done <- true:
				return nil
			}
		}
	}
}
