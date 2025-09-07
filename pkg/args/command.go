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
	if len(os.Args) != 5 {
		err2 := err.NoneContext()
		err2.SetError(errors.New("github.com/dywoq/gh-issue/pkg/args: len(os.Args) is not 5"))
		err2.SetMore("source is args.CommandProcessGet() (*Args, err.Context)")
		return nil, err2
	}
	args := &Args{}
	args.Args[1] = os.Args[1] // issue ids
	args.Args[2] = os.Args[2] // owner
	args.Args[3] = os.Args[3] // repository
	args.Args[4] = os.Args[4] // token
	return args, err.NoneContext()
}

func CommandArgumentsDelete() (*Args, err.Context) {
	if len(os.Args) != 5 {
		err2 := err.NoneContext()
		err2.SetError(errors.New("github.com/dywoq/gh-issue/pkg/args: len(os.Args) is not 5"))
		err2.SetMore("source is args.CommandProcessDelete() (*Args, err.Context)")
		return nil, err2
	}
	args := &Args{}
	args.Args[1] = os.Args[1] // issue ids
	args.Args[2] = os.Args[2] // owner
	args.Args[3] = os.Args[3] // repository
	args.Args[4] = os.Args[4] // token
	return args, err.NoneContext()
}
