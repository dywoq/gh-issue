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
	"fmt"
	"os"

	"github.com/dywoq/dywoqlib/err"
	"github.com/dywoq/gh-issue/args"
	"github.com/dywoq/gh-issue/process"
)

func outputError(err err.Context) {
	if err.Nil() {
		return
	}
	fmt.Printf("github.com/dywoq/gh-issue: error occurred: %v\n", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		text := `
		gh-issue: CLI tool to manage over GitHub issues - v1.2.1

		The project's repository: https://github.com/dywoq/gh-issue

		The syntax:
		gh-issue <command-name> <args...>

		For example, this outputs all your GitHub repository issues in the console:
		gh-issue get * <your-nickname> <your-repository> <your-token>

		The project is licensed under Apache License 2.0, see https://github.com/dywoq/gh-issue/blob/main/LICENSE
		`
		fmt.Println(text)
		os.Exit(0)
	}

	a, err2 := args.New()
	if !err2.Nil() {
		outputError(err2)
	}

	if a == nil {
		fmt.Printf("github.com/dywoq/gh-issue: a is nil\n")
		os.Exit(1)
	}

	commands := map[args.Command]func() err.Context{
		args.CommandGet: func() err.Context {
			return process.Get(a)
		},
		args.CommandGetConfig: func() err.Context {
			return process.GetConfig(a)
		},

		args.CommandClose: func() err.Context {
			choice := ""
			fmt.Printf("are you sure to close the chosen issues? [y: Yes, n: No]:\n>> ")
			_, err1 := fmt.Scanf("%s", &choice)
			if err1 != nil {
				return err.NewContext(err1, "source is main.commands[ags.CommandClose]func() err.Context")
			}
			if choice == "y" {
				return process.Close(a)
			}
			return err.NoneContext()
		},
		args.CommandCloseConfig: func() err.Context {
			choice := ""
			fmt.Printf("are you sure to close the chosen issues? [y: Yes, n: No]:\n>> ")
			_, err1 := fmt.Scanf("%s", &choice)
			if err1 != nil {
				return err.NewContext(err1, "source is main.commands[ags.CommandClose]func() err.Context")
			}
			if choice == "y" {
				return process.CloseConfig(a)
			}
			return err.NoneContext()
		},

		args.CommandGenerateMd: func() err.Context {
			return process.GenerateMd(a)
		},
		args.CommandGenerateMdConfig: func() err.Context {
			return process.GenerateMdConfig(a)
		},
	}

	for cmd, f := range commands {
		if a.Command == cmd {
			err2 := f()
			if !err2.Nil() {
				outputError(err2)
			}
		}
	}
}
