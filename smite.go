package main

import (
	"fmt"
)

var smiteCmd = &Command{
	Name:    "smite",
	Usage:   "[OPTION] FILE",
	Summary: "Delete file from database by name,",
	Help: `The smite command deletes a file from the database by filename.
If more than one file exists with the given name, you will 
be prompted which file to delete, unless the --all flag is 
given.`,
}

var (
	smiteAll = smiteCmd.Flag.Bool("all", false,
		"smite all files with the given name")
	smiteForce = smiteCmd.Flag.Bool("force", false,
		"delete files without prompting")
)

func smiteFunc(db *Db, cmd *Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Insufficient arguments.")
	}
	if err := db.Delete(args[0]); err != nil {
		return err
	}
	return nil
}

func init() {
	smiteCmd.Function = smiteFunc
}
