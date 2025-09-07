package args

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/dywoq/dywoqlib/lib/err"
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
func New() (*Args, err.Context) {
	if len(os.Args) != 5 {
		return nil, err.NewContext(errors.New("github.com/dywoq/gh-issue/pkg/args: len(os.Args) is not 5"), "source is args.New() (*args.Args, err.Context)")
	}
	args := &Args{}
	formattedIssues, err2 := issuesFormat(os.Args[1])
	if err2 != nil {
		copy := err.NoneContext()
		copy.Copy(err2)
		return nil, copy
	}

	args.Issues = formattedIssues
	args.Owner = os.Args[2]
	args.Repository = os.Args[3]
	args.Token = os.Args[4]
	return args, err.NoneContext()
}

func issuesFormat(arg string) ([]int64, err.Context) {
	if arg == "" {
		return []int64{}, err.NewContext(errors.New("github.com/gh-issue/pkg/args: the argument string is empty"), "source is args.issuesFormat(string) ([]int64, err.Context)")
	}
	split := strings.Split(arg, ",")
	ids := []int64{}
	for _, part := range split {
		id, err2 := strconv.Atoi(part)
		if err2 != nil {
			return []int64{}, err.NewContext(err2, "source is args.issuesFormat(string) ([]int64, err.Context)")
		}
		ids = append(ids, int64(id))
	}
	return ids, err.NoneContext()
}
