package main

import (
	"errors"
	"fmt"
	"sync"

	"github.com/dywoq/dywoqlib/lib/container"
	"github.com/dywoq/dywoqlib/lib/err"
	"github.com/dywoq/gh-issue/pkg/args"
	"github.com/dywoq/gh-issue/pkg/issue"
	"github.com/google/go-github/v74/github"
)

func outputIssue(i *github.Issue) {
	if i == nil {
		return
	}
	fmt.Println("Retrieved issue: {")
	fmt.Printf("	Title: %v\n", getTitle(i))
	fmt.Printf("	Labels: %v\n", container.FormattableSlice[string](getLabels(i)))
	fmt.Printf("	Milestone: %v\n", getMilestone(i))
	fmt.Printf("	Assignees: %v\n", container.FormattableSlice[string](getAssignees(i)))
	fmt.Printf("    Date: %v\n", getDate(i))
	fmt.Printf("	Body: %v\n", getFormattedMarkdownBody(i))
	fmt.Println("}")
}

func processGet(a *args.Args) err.Context {
	if a.Args[1] != args.CommandGet {
		return err.NewContext(
			errors.New("github.com/dywoq/gh-issue: wrong command, expected args.CommandGet"),
			"source is process.go: processGet(*args.Args) err.Context",
		)
	}

	failedTypeAssertion := err.NewContext(
		errors.New("github.com/dywoq/gh-issue: failed type assertion"),
		"source is process.go: processGet(*args.Args) err.Context",
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

func issueCloseAsync(wg *sync.WaitGroup, errch chan error, owner, repo, token string, id int) {
	defer wg.Done()
	fmt.Printf("deleting github issue %d in %s/%s repository...\n", id, owner, repo)
	err2 := issue.Close(owner, repo, token, id)
	if !err2.Nil() {
		errch <- err2.Error()
		fmt.Printf("failed deleting github issue %d in %s/%s repository\n", id, owner, repo)
	}
}

func processClose(a *args.Args) err.Context {
	failedTypeAssertion := err.NewContext(
		errors.New("github.com/dywoq/gh-issue: failed type assertion"),
		"source is process.go: processDelete(*args.Args) err.Context",
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

	var wg sync.WaitGroup
	errch := make(chan error, 1)

	switch ids {
	// if ids is wildcard
	case "*":
		gotIds, err2 := issue.GetAllId(owner, repo, token)
		if !err2.Nil() {
			return err2
		}
		wg.Add(len(gotIds))
		for _, id := range gotIds {
			go issueCloseAsync(&wg, errch, owner, repo, token, id)
		}

	// if ids is represented as list (2,3,4)
	default:
		formattedIds, err2 := issue.FormatToIntSlice(ids.(string))
		if !err2.Nil() {
			return err2
		}
		wg.Add(len(formattedIds))
		for _, formattedId := range formattedIds {
			go issueCloseAsync(&wg, errch, owner, repo, token, formattedId)
		}
	}

	go func() {
		wg.Wait()
		close(errch)
	}()

	if err2, ok := <-errch; ok {
		return err.NewContext(
			err2,
			"source is process.go: processDelete(*args.Args) err.Context",
		)
	}

	return err.NoneContext()
}
