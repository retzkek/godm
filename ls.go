package main

import (
	"fmt"
	"labix.org/v2/mgo/bson"
)

var lsCmd = &Command{
	Name:    "ls",
	Usage:   "[OPTION] [REGEX]",
	Summary: "List all files, optionally only those matching REGEX.",
	Help: `The ls command lists files stored in the database, optionally
listing only files whose name matches the given regular expression.

By default the match is case-sensitive and matches the full file name.`,
}

var (
	lsCase    = lsCmd.Flag.Bool("i", false, "case-insensitive match")
	lsPartial = lsCmd.Flag.Bool("partial", false, "find partial matches")
)

func lsFunc(db *Db, cmd *Command, args []string) error {
	var query interface{}
	if len(args) > 0 {
		var regex string
		if *lsPartial {
			regex = args[0]
		} else {
			regex = "^" + args[0] + "$" // make the regex match the full filename
		}
		var opt string
		if *lsCase {
			opt += "i" // case-insensitive match
		}
		query = bson.M{"filename": bson.RegEx{regex, opt}}
	}
	files, err := db.List(query)
	if err != nil {
		return err
	}
	for _, f := range files {
		fmt.Printf("%10d %16s %20s\n", f.Length,
			f.UploadDate.Format("2006-01-02 15:04"), f.Filename)
	}
	return nil
}

func init() {
	lsCmd.Function = lsFunc
}
