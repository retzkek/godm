package main

import (
	"fmt"
)

var storeCmd = &Command{
	Name:    "store",
	Usage:   "[OPTION] FILE",
	Summary: "Store file in database,",
	Help: `The store command adds a file to the database, optionally
adding tags and/or other metadata.`,
}

var (
	storeTags = storeCmd.Flag.String("tags", "", "list of tags to be added to file (comma-delimited)")
)

func storeFunc(db *Db, cmd *Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Insufficient arguments.")
	}
	if err := db.Store(args[0]); err != nil {
		return err
	}
	return nil
}

func init() {
	storeCmd.Function = storeFunc
}
