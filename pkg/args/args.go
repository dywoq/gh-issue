package args

import (
	"github.com/dywoq/dywoqlib/err"
	"os"
)

// Args is a struct presenting arguments of gh-issue program.
type Args struct {
	Command Command
	Args    []any
}

// New creates and returns a pointer to Args, along with the possible encountered error.
func New() (*Args, err.Context) {
	switch os.Args[1] {
	case string(CommandGet):
		return CommandArgumentsGet()
	case string(CommandClose):
		return CommandArgumentsClose()
	}
	return nil, err.NoneContext()
}
