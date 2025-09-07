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

	"github.com/dywoq/gh-issue/pkg/args"
)

func main() {
	if len(os.Args) < 2 {
		text := `
		gh-issue: CLI tool to manage over GitHub issues

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

	a, err := args.New()
	if !err.Nil() {
		fmt.Printf("github.com/dywoq/gh-issue: error occurred: %v\n", err)
		os.Exit(1)
	}

	if a == nil {
		fmt.Printf("github.com/dywoq/gh-issue: a is nil\n")
		os.Exit(1)
	}

	switch a.Command {
	case args.CommandGet:
		err := processGet(a)
		if !err.Nil() {
			fmt.Printf("github.com/dywoq/gh-issue: error occurred: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	case args.CommandClose:
		choice := ""

		fmt.Println("are you sure to close the chosen issues? [y: Yes, n: No]")
		fmt.Print("> ")

		_, err := fmt.Scanf("%s", &choice)
		if err != nil {
			fmt.Printf("github.com/dywoq/gh-issue: error occurred: %v\n", err)
		}

		if choice == "y" {
			err := processClose(a)
			if !err.Nil() {
				fmt.Printf("github.com/dywoq/gh-issue: error occurred: %v\n", err)
				os.Exit(1)
			}
		}
		os.Exit(0)
	}
}
