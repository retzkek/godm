package main

import (
	"fmt"
)

var helpCmd = &Command{
	Name:    "help",
	Usage:   "help [COMMAND]",
	Summary: "Print command usage and options.",
	Help: `godm is a simple general-purpose document management system.
General Usage:
	godm [OPTION] COMMAND [COMMAND OPTION]`,
}

func helpFunc(db *Db, cmd *Command, args []string) error {
	fmt.Println("Help is not available... good luck.")
	return nil
}

func init() {
	helpCmd.Function = helpFunc
}
