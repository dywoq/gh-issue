package issue

import (
	"context"

	"github.com/google/go-github/v74/github"
	"golang.org/x/oauth2"
)

func client(token string) (*github.Client, context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client, ctx
}
