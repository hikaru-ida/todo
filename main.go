package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("todo", "0.1.0")

	c.Args = os.Args[1:]

	c.Commands = Commands()

	exitCode, err := c.Run()
	if err != nil {
		fmt.Printf("Failed to excute: %s\n", err.Error())
	}

	os.Exit(exitCode)
}
