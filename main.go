package main


import (
	"b/cmd"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Printf("Error getting current user. %s", err)
	}
	fmt.Printf("|  b   -  v0.0.1  |  user: %s  \n", user.Username)
	fmt.Printf("type in commands to see the output.\n")
	repl.Start(os.Stdin, os.Stdout)
}