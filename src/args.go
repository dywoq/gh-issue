package main

import "os"

// argsGet returns the program arguments, excluding the executable path.
// It expects four arguments from the command line:
//
// - A comma-separated list of issue IDs or a wildcard.
//
// - The repository owner's name.
//
// - The repository's name.
//
// - A GitHub personal access token.
func argsGet() []string {
	args := os.Args
	if len(args) != 5 {
		panic("github.com/dywoq/gh-issue-deleter: expected 4 arguments, but got " + string(rune(len(os.Args)-1)))
	}
	return args[1:]
}
