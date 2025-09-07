package issue

import (
	"errors"

	"github.com/dywoq/dywoqlib/err"
	"github.com/google/go-github/v74/github"
)

// Delete deletes the Github issue with given id, repo, token, its owner and token.
// If owner, repo or token is empty, it returns error.
func Delete(owner, repo, token string, id int) err.Context {
	if owner == "" || repo == "" || token == "" {
		return err.NewContext(errors.New("github.com/dywoq/gh-issue/pkg/issue: owner, repo or token is empty"), "source is issue.Delete(string, string, string, int) err.Context")
	}
	c, ctx := client(token)
	req := &github.IssueRequest{State: github.Ptr("closed")}
	_, _, err2 := c.Issues.Edit(ctx, owner, repo, id, req)
	if err2 != nil {
		return err.NewContext(err2, "source is issue.Delete(string, string, string, int) err.Context")
	}
	return err.NoneContext()
}
