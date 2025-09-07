package main

import (
	"fmt"
	"os"

	"github.com/dywoq/dywoqlib/lib/attribute"
	"github.com/dywoq/gh-issue/pkg/args"
)

func main() {
	a, err := args.New()
	if !err.Nil() {
		fmt.Printf("github.com/dywoq/gh-issue: error occurred: %v", err)
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
			fmt.Printf("github.com/dywoq/gh-issue: error occurred: %v", err)
			os.Exit(1)
		}
		os.Exit(0)
	case args.CommandDelete:
		attribute.Todo(nil)
		os.Exit(0)
	}
}
