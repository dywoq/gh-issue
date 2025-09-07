// Copyright 2025 dywoq
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"fmt"
	"sync"

	"github.com/dywoq/dywoqlib/container"
	"github.com/dywoq/dywoqlib/err"
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
	fmt.Printf("	Body:\n%v\n", getFormattedMarkdownBody(i))
	fmt.Println("}")
}

func processGet(a *args.Args) err.Context {
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
	fmt.Printf("closing github issue %d in %s/%s repository...\n", id, owner, repo)
	err2 := issue.Close(owner, repo, token, id)
	if !err2.Nil() {
		errch <- err2.Error()
		fmt.Printf("failed closing github issue %d in %s/%s repository\n", id, owner, repo)
	}
}

func processClose(a *args.Args) err.Context {
	failedTypeAssertion := err.NewContext(
		errors.New("github.com/dywoq/gh-issue: failed type assertion"),
		"source is process.go: processClose(*args.Args) err.Context",
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
			"source is process.go: processClose(*args.Args) err.Context",
		)
	}

	return err.NoneContext()
}
