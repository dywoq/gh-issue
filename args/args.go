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
	case string(CommandGetConfig):
		return CommandArgumentsGetConfig()

	case string(CommandClose):
		return CommandArgumentsClose()

	case string(CommandGenerateMd):
		return CommandArgumentsGenerateMd()
	case string(CommandGenerateMdConfig):
		return CommandArgumentsGenerateMdConfig()
	}
	return nil, err.NoneContext()
}
