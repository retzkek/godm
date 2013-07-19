package main

import (
	"fmt"
)

var getCmd = &Command{
	Name:    "get",
	Usage:   "[OPTION] FILE",
	Summary: "Retrieve file from database by name,",
	Help: `The get command retrieves a file from the database by filename and
saves it in the current working directory. If more than one file 
exists with the given name, you will be prompted which file to get.`,
}

var (
	getPath = getCmd.Flag.String("path", "",
		"path to save file in (default is current working dir)")
)

func getFunc(db *Db, cmd *Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Insufficient arguments.")
	}
	if err := db.Get(args[0]); err != nil {
		return err
	}
	return nil
}

func init() {
	getCmd.Function = getFunc
}
