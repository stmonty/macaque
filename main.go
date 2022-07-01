package main

import (
	"fmt"
	"simian/repl"
	"os"
)

func main() {

	//	user, err := user.Current()
	//	if err != nil {
	//		panic(err)
	//	}
	fmt.Printf("Welcome to Simian, an interpreter for the Monkey Language\n")
	fmt.Printf("REPL Starting:\n")
	repl.Start(os.Stdin, os.Stdout)

}
