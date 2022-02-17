package main

import (
	"fmt"
	"macaque/repl"
	"os"
)

func main() {

	//	user, err := user.Current()
	//	if err != nil {
	//		panic(err)
	//	}
	fmt.Printf("Welcome to Macaque, an interpreter for the Monkey Language\n")
	fmt.Printf("REPL Starting:\n")
	repl.Start(os.Stdin, os.Stdout)

}
