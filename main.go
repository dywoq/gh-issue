package main

import (
	"fmt"
	"os"

	"github.com/dywoq/gh-issue/pkg/args"
)

func main() {
	if len(os.Args) == 0 {
		text := `
		
		
		gh-issue: CLI tool to manage over GitHub issues
		
		The project's repository: https://github.com/dywoq/gh-issue

		The syntax:
		gh-issue <command-name> <args...>
		`
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
