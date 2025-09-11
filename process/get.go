package process

import (
	"errors"
	"fmt"

	"github.com/dywoq/dywoqlib/container"
	"github.com/dywoq/dywoqlib/err"
	"github.com/dywoq/gh-issue/args"
	"github.com/dywoq/gh-issue/get"
	"github.com/dywoq/gh-issue/issue"
	"github.com/google/go-github/v74/github"
)

func outputIssue(i *github.Issue) {
	if i == nil {
		return
	}
	t := get.Terminal{Issue: i}
	fmt.Println("Retrieved issue: {")
	fmt.Printf("	Title: %v\n", t.Title())
	fmt.Printf("	State: %v\n", t.State())
	fmt.Printf("	Labels: %v\n", container.FormattableSlice[string](t.Labels()))
	fmt.Printf("	Milestone: %v\n", t.Milestone())
	fmt.Printf("	Assignees: %v\n", container.FormattableSlice[string](t.Assignees()))
	fmt.Printf("    Date: %v\n", t.Date())
	fmt.Printf("	Body:\n%v\n", t.Body())
	fmt.Println("}")
}

func getBase(ids any, owner, repo, token string) err.Context {
	switch ids {
	// if ids is wildcard
	case "*":
		gotIds, err2 := issue.GetAllId(owner, repo, token)
		if !err2.Nil() {
			return err2
		}
		for _, elem := range gotIds {
			i, err2 := issue.Get(owner, repo, token, elem)
			if !err2.Nil() {
				return err2
			}
			outputIssue(i)
		}

	// if ids is represented as list (2,3,4)
	default:
		formattedIds, err2 := issue.FormatToIntSlice(ids.(string))
		if !err2.Nil() {
			return err2
		}
		for _, formattedId := range formattedIds {
			i, err2 := issue.Get(owner, repo, token, formattedId)
			if !err2.Nil() {
				return err2
			}
			outputIssue(i)
		}
	}
	return err.NoneContext()
}

func Get(a *args.Args) err.Context {
	failedTypeAssertion := err.NewContext(
		errors.New("github.com/dywoq/gh-issue: failed type assertion"),
		"source is process.go: process.Get(*args.Args) err.Context",
	)
	ids := a.Args[2]
	owner, ok := a.Args[3].(string)
	if !ok {
		return failedTypeAssertion
	}
	repo, ok := a.Args[4].(string)
	if !ok {
		return failedTypeAssertion
	}
	token, ok := a.Args[5].(string)
	if !ok {
		return failedTypeAssertion
	}
	return getBase(ids, owner, repo, token)
}

func GetConfig(a *args.Args) err.Context {
	failedTypeAssertion := err.NewContext(
		errors.New("github.com/dywoq/gh-issue: failed type assertion"),
		"source is process.GetConfig(*args.Args) err.Context",
	)
	configPath, ok := a.Args[2].(string)
	if !ok {
		return failedTypeAssertion
	}
	conf, err2 := newConfig(configPath)
	if !err2.Nil() {
		return err2
	}
	ids := a.Args[2]
	return getBase(ids, conf.Owner, conf.Repository, conf.Token)
}
