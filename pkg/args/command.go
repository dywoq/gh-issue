package args

import (
	"errors"
	"os"

	"github.com/dywoq/dywoqlib/lib/err"
)

type Command string

const (
	CommandGet    Command = "get"
	CommandDelete Command = "delete"
)

func CommandArgumentsGet() (*Args, err.Context) {
	if len(os.Args) != 6 {
		err2 := err.NoneContext()
		err2.SetError(errors.New("github.com/dywoq/gh-issue/pkg/args: len(os.Args) is not 5"))
		err2.SetMore("source is args.CommandArgumentsGet() (*Args, err.Context)")
		return nil, err2
	}
	args := &Args{
		CommandGet,
		make([]any, 6),
	}
	args.Args[2] = os.Args[2] // issue ids
	args.Args[3] = os.Args[3] // owner
	args.Args[4] = os.Args[4] // repository
	args.Args[5] = os.Args[5] // token
	return args, err.NoneContext()
}

func CommandArgumentsDelete() (*Args, err.Context) {
	if len(os.Args) != 6 {
		err2 := err.NoneContext()
		err2.SetError(errors.New("github.com/dywoq/gh-issue/pkg/args: len(os.Args) is not 5"))
		err2.SetMore("source is args.CommandArgumentsDelete() (*Args, err.Context)")
		return nil, err2
	}
	args := &Args{
		CommandDelete,
		make([]any, 6),
	}
	args.Args[2] = os.Args[2] // issue ids
	args.Args[3] = os.Args[3] // owner
	args.Args[4] = os.Args[4] // repository
	args.Args[5] = os.Args[5] // token
	return args, err.NoneContext()
}
