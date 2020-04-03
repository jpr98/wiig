package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/jpr98/wiig/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("Use !!exit to end the interactive session")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
	fmt.Println("Thanks for using the Monkey REPL!")
}
