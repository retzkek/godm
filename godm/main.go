package main

import (
	"fmt"
	"github.com/retzkek/godm"
	"os"
)

func printHelp(args []string) {
	fmt.Printf("\nusage: %s cmd [args]\n", args[0])
	fmt.Printf(`commands:
	help [COMMAND]
		Show help on the program or a command.
	store FILE [TAGS]
		Store file with optional comma-delimited tags.
	get FILE
		Get file by name.
	smite FILE
		Delete file.
	ls QUERY
		List all files with name matching QUERY.
	tag FILE TAGS
		Add TAGS (comma-delimited) to FILE.
	
	search TERMS
		Search for a file.`)
}

type CommandFunc func([]string)

type Command struct {
	Usage     string
	BriefHelp string
	Function  CommandFunc
}

func main() {
	db := godm.Db{"localhost:27017", "test"}
	var cmds = map[string]Command{
		"help": {"help [COMMAND]",
			"Show help on the program or a command.",
			func(args []string) {
				fmt.Printf("\nusage: %s cmd [args]\n", args[0])
			}},
		"store": {"store FILE [TAGS]",
			"Store file with optional comma-delimited tags.",
			func(args []string) {
				if len(args) < 3 {
					printHelp(os.Args)
					os.Exit(1)
				}
				if err := db.Store(args[2]); err != nil {
					panic(err)
				}
			}},
		"get": {"get FILE",
			"Get file by name.",
			func(args []string) {
				if len(args) < 3 {
					printHelp(os.Args)
					os.Exit(1)
				}
				if err := db.Get(args[2]); err != nil {
					panic(err)
				}
			}},
		"smite": {"smite FILE",
			"Delete file.", nil},
		"ls": {"ls QUERY",
			"List all files with name matching QUERY.", nil},
		"tag": {"tag FILE TAGS",
			"Add TAGS (comma-delimited list) to FILE.", nil},
		"search": {"search QUERY",
			"Search for all files matching QUERY.", nil},
	}
	if len(os.Args) < 2 {
		printHelp(os.Args)
		os.Exit(1)
	}
	cmd := os.Args[1]
	if c, ok := cmds[cmd]; ok {
		if c.Function != nil {
			c.Function(os.Args)
		} else {
			fmt.Println("not implemented")
		}
	} else {
		printHelp(os.Args)
		os.Exit(1)
	}
}
