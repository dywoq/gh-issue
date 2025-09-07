package args

import (
	"errors"
	"os"

	"github.com/dywoq/dywoqlib/err"
)

// Args is a struct presenting arguments of gh-issue program.
type Args struct {
	Command Command
	Args    []any
}

// New creates and returns a pointer to Args, along with the possible encountered error.
func New() (*Args, err.Context) {
	if len(os.Args) < 2 {
		return nil, err.NewContext(
			errors.New("github.com/dywoq/gh-issue: no command provided"),
			"source is args.New() (*Args, err.Context)",
		)
	}
	switch os.Args[1] {
	case string(CommandGet):
		return CommandArgumentsGet()
	case string(CommandClose):
		return CommandArgumentsClose()
	}
	return nil, err.NoneContext()
}
