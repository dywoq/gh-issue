package args

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// Args is a struct presenting arguments of gh-issue program.
type Args struct {
	Issues     []int64
	Owner      string
	Repository string
	Token      string
}

// New creates and returns a pointer to Args, along with the possible encountered error.
// If len(os.Args) is not 5, it returns nil and a error.
func New() (*Args, error) {
	if len(os.Args) != 5 {
		return nil, errors.New("github.com/gh-issue/pkg/args: len(os.Args) is not 5")
	}
	args := &Args{}
	formattedIssues, err := issuesFormat(os.Args[1])
	if err != nil {
		return nil, err
	}

	args.Issues = formattedIssues
	args.Owner = os.Args[2]
	args.Repository = os.Args[3]
	args.Token = os.Args[4]
	return args, nil
}

func issuesFormat(arg string) ([]int64, error) {
	if arg == "" {
		return []int64{}, errors.New("github.com/gh-issue/pkg/args: the argument string is empty")
	}
	split := strings.Split(arg, ",")
	ids := []int64{}
	for _, part := range split {
		id, err := strconv.Atoi(part)
		if err != nil {
			return []int64{}, err
		}
		ids = append(ids, int64(id))
	}
	return ids, nil
}
