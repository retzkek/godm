package main

import (
	"fmt"
)

var tourCmd = &Command{
	Name:    "tour",
	Usage:   "",
	Summary: "Take the godm tour.",
}

func tourFunc(db *Db, cmd *Command, args []string) error {
	fmt.Println("Welcome to the godm tour... Where can I get some godm bait?")
	return nil
}

func init() {
	tourCmd.Function = tourFunc
}
