package main

import (
	"fmt"
	"github.com/meriy100/goInterpreter/monkey/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! this is the Monky programming language!\n", u.Username)

	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
