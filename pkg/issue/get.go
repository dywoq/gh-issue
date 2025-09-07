package issue

import (
	"errors"

	"github.com/dywoq/dywoqlib/lib/err"
	"github.com/google/go-github/v74/github"
)

// Get gets the issue from GitHub with given repo, token,
// its owner and id, and returns it.
// If owner, repo or token is empty, it returns nil and error.
func Get(owner, repo, token string, id int) (*github.Issue, err.Context) {
	if owner == "" || repo == "" || token == "" {
		return nil, err.NewContext(errors.New("github.com/dywoq/gh-issue/pkg/issue: owner, repo or token is empty"), "source is issue.Get(string, string, string, int) (*github.Issue, err.Context)")
	}
	c, ctx := client(token)
	issue, _, err2 := c.Issues.Get(ctx, owner, repo, id)
	if err2 != nil {
		return nil, err.NewContext(err2, "source is issue.Get(string, string, string, int) (*github.Issue, err.Context)")
	}
	return issue, err.NoneContext()
}
