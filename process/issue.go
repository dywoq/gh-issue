package process

import (
	"errors"

	"github.com/dywoq/dywoqlib/err"
	"github.com/dywoq/gh-issue/issue"
	"github.com/google/go-github/v74/github"
)

func issuesGetFromIds(owner, repo, token string, ids []int) ([]*github.Issue, err.Context) {
	if len(ids) == 0 {
		return []*github.Issue{}, err.NewContext(
			errors.New("github.com/dywoq/gh-issue/process: len(ids) is 0"),
			"source is process.issuesGetFromIds(string, string, string, []int) ([]*github.Issue, err.Context)",
		)
	}
	issues := []*github.Issue{}
	for _, id := range ids {
		i, err2 := issue.Get(owner, repo, token, id)
		if !err2.Nil() {
			return []*github.Issue{}, err2
		}
		issues = append(issues, i)
	}
	return issues, err.NoneContext()
}
