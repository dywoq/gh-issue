package process

import (
	"errors"
	"fmt"
	"sync"

	"github.com/dywoq/dywoqlib/err"
	"github.com/dywoq/gh-issue/args"
	"github.com/dywoq/gh-issue/issue"
)

func asyncClose(owner, repo, token string, id int, wg *sync.WaitGroup, errch chan err.Context) {
	defer wg.Done()
	fmt.Printf("closing https://github.com/%s/%s/issues/%d issue...\n", owner, repo, id)
	err2 := issue.Close(owner, repo, token, id)
	if !err2.Nil() {
		fmt.Printf("failed closing https://github.com/%s/%s/issues/%d issue\n", owner, repo, id)
		errch <- err2
		return
	}
	fmt.Printf("closed https://github.com/%s/%s/issues/%d issue\n", owner, repo, id)
}

func Close(a *args.Args) err.Context {
	failedTypeAssertion := err.NewContext(
		errors.New("github.com/dywoq/gh-issue: failed type assertion"),
		"source is process.go: process.Close(*args.Args) err.Context",
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

	var (
		errch = make(chan err.Context, 1)
		wg    = sync.WaitGroup{}
	)

	switch ids {
	case "*":
		ids, err2 := issue.GetAllId(owner, repo, token)
		if !err2.Nil() {
			return err2
		}

		wg.Add(len(ids))

		for _, id := range ids {
			go asyncClose(owner, repo, token, id, &wg, errch)
		}
	}

	wg.Wait()
	close(errch)

	if err2, ok := <-errch; ok {
		if !err2.Nil() {
			return err2
		}
	}

	return err.NoneContext()
}
