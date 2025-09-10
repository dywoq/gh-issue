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

package args

import (
	"errors"
	"os"

	"github.com/dywoq/dywoqlib/err"
)

type Command string

const (
	CommandGet       Command = "get"
	CommandGetConfig Command = "get-config"

	CommandClose       Command = "close"
	CommandCloseConfig Command = "close-config"

	CommandGenerateMd       Command = "generate-md"
	CommandGenerateMdConfig Command = "generate-md-config"
)

func CommandArgumentsGet() (*Args, err.Context) {
	if len(os.Args) != 6 {
		err2 := err.NoneContext()
		err2.SetError(errors.New("github.com/dywoq/gh-issue/pkg/args: len(os.Args) is not 6"))
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

func CommandArgumentsClose() (*Args, err.Context) {
	if len(os.Args) != 6 {
		err2 := err.NoneContext()
		err2.SetError(errors.New("github.com/dywoq/gh-issue/pkg/args: len(os.Args) is not 6"))
		err2.SetMore("source is args.CommandArgumentsClose() (*Args, err.Context)")
		return nil, err2
	}
	args := &Args{
		CommandClose,
		make([]any, 6),
	}
	args.Args[2] = os.Args[2] // issue ids
	args.Args[3] = os.Args[3] // owner
	args.Args[4] = os.Args[4] // repository
	args.Args[5] = os.Args[5] // token
	return args, err.NoneContext()
}

func CommandArgumentsCloseConfig() (*Args, err.Context) {
	if len(os.Args) != 4 {
		err2 := err.NoneContext()
		err2.SetError(errors.New("github.com/dywoq/gh-issue/pkg/args: len(os.Args) is not 4"))
		err2.SetMore("source is args.CommandArgumentsCloseConfig() (*Args, err.Context)")
		return nil, err2
	}
	args := &Args{
		CommandCloseConfig,
		make([]any, 4),
	}
	args.Args[2] = os.Args[2] // issue ids
	args.Args[3] = os.Args[3] // config
	return args, err.NoneContext()
}

func CommandArgumentsGenerateMd() (*Args, err.Context) {
	if len(os.Args) != 7 {
		err2 := err.NoneContext()
		err2.SetError(errors.New("github.com/dywoq/gh-issue/pkg/args: len(os.Args) is not 7"))
		err2.SetMore("source is args.CommandArgumentsGenerateMd() (*Args, err.Context)")
		return nil, err2
	}
	args := &Args{
		CommandGenerateMd,
		make([]any, 7),
	}
	args.Args[2] = os.Args[2] // issue ids
	args.Args[3] = os.Args[3] // owner
	args.Args[4] = os.Args[4] // repository
	args.Args[5] = os.Args[5] // token
	args.Args[6] = os.Args[6] // filename
	return args, err.NoneContext()
}

func CommandArgumentsGenerateMdConfig() (*Args, err.Context) {
	if len(os.Args) != 5 {
		err2 := err.NoneContext()
		err2.SetError(errors.New("github.com/dywoq/gh-issue/pkg/args: len(os.Args) is not 5"))
		err2.SetMore("source is args.CommandArgumentsGenerateMdConfig() (*Args, err.Context)")
		return nil, err2
	}
	args := &Args{
		CommandGenerateMdConfig,
		make([]any, 5),
	}
	args.Args[2] = os.Args[2] // config path
	args.Args[3] = os.Args[3] // issues ids
	args.Args[4] = os.Args[4] // filename
	return args, err.NoneContext()
}

func CommandArgumentsGetConfig() (*Args, err.Context) {
	if len(os.Args) != 4 {
		err2 := err.NoneContext()
		err2.SetError(errors.New("github.com/dywoq/gh-issue/pkg/args: len(os.Args) is not 4"))
		err2.SetMore("source is args.CommandArgumentsGetConfig() (*Args, err.Context)")
		return nil, err2
	}
	args := &Args{
		CommandGetConfig,
		make([]any, 4),
	}
	args.Args[2] = os.Args[2] // config path
	args.Args[3] = os.Args[3] // issues ids
	return args, err.NoneContext()
}
