package main

import (
	"fmt"
	"monkey-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Monkey programming language!\n", user.Username)
	fmt.Printf("Start typing commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
