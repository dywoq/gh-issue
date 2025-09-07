package main

import (
	"fmt"
	"os"

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
	case args.CommandClose:
		choice := ""

		fmt.Println("are you sure to close the chosen issues? [y: Yes, n: No]")
		fmt.Print("> ")

		_, err := fmt.Scanf("%s", &choice)
		if err != nil {
			fmt.Printf("github.com/dywoq/gh-issue: error occurred: %v", err)
		}

		if choice == "y" {
			err := processClose(a)
			if !err.Nil() {
				fmt.Printf("github.com/dywoq/gh-issue: error occurred: %v", err)
				os.Exit(1)
			}
		}
		os.Exit(0)
	}
}
