package issue

import (
	"errors"

	"github.com/dywoq/dywoqlib/err"
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

// GetAllId gets all issue and returns their IDs.
// its owner and id, and returns it.
// If owner, repo or token is empty, it returns nil and error.
func GetAllId(owner, repo, token string) ([]int, err.Context) {
	if owner == "" || repo == "" || token == "" {
		return []int{}, err.NewContext(errors.New("github.com/dywoq/gh-issue/pkg/issue: owner, repo or token is empty"), "source is issue.GetAllId(string, string, string, int) ([]int, err.Context)")
	}
	c, ctx := client(token)
	result := []int{}
	options := &github.IssueListByRepoOptions{
		State: "all",
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}
	for {
		issues, resp, err2 := c.Issues.ListByRepo(ctx, owner, repo, options)
		if err2 != nil {
			return []int{}, err.NewContext(err2, "source is issue.GetAllId(string, string, string) ([]int64, err.Context)")
		}
		for _, issue := range issues {
			result = append(result, issue.GetNumber())
		}
		if resp.NextPage == 0 {
			break
		}
		options.ListOptions.Page = resp.NextPage
	}
	return result, err.NoneContext()
}
