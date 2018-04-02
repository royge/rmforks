package repo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func TestServiceFetch(t *testing.T) {
	server := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		id := int64(123)
		name := "test"
		login := "tester"
		user := github.User{Login: &login, ID: &id}
		repos := []github.Repository{
			github.Repository{
				ID:    &id,
				Name:  &name,
				Owner: &user,
			},
		}

		json.NewEncoder(w).Encode(repos)
	})

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "token"},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	ctx := context.Background()

	svc := Service{
		User:     "royge",
		Exclude:  []string{"test", "test"},
		Timeout:  2,
		Register: make(chan *github.Repository),
		Client:   github.NewClient(tc),
	}

	var err error
	svc.Client.BaseURL, err = url.Parse(server.URL + "/")
	if err != nil {
		t.Fatalf("could not parse url: %v", err)
	}

	repos := []*github.Repository{}
	go svc.Fetch(ctx)
	for v := range svc.Register {
		repos = append(repos, v)
	}
}

func newMockServer(handler func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(handler))
}
