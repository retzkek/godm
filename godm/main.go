package main

import (
	"fmt"
	"github.com/retzkek/godm"
	"os"
)

type CommandFunc func([]string)

type Command struct {
	Usage     string
	BriefHelp string
	Function  CommandFunc
}

func printHelp(args []string, cmds map[string]Command) {
	fmt.Printf("\nusage: %s COMMAND [ARGS]\n", args[0])
	fmt.Printf("commands:\n")
	fmt.Printf("\thelp [COMMAND]\n")
	fmt.Printf("\t\tShow this help or get detailed help on a command.\n")
	for _, c := range cmds {
		if c.Function != nil {
			fmt.Printf("\t%s\n", c.Usage)
			fmt.Printf("\t\t%s\n", c.BriefHelp)
		}
	}
}

func main() {
	db := godm.Db{"localhost:27017", "test"}
	// 'help' is not in the map as it requires
	// the contents of the map.
	var cmds = map[string]Command{
		"store": {"store FILE [TAGS]",
			"Store file with optional comma-delimited tags.",
			func(args []string) {
				if len(args) < 3 {
					fmt.Println("Insufficient arguments. 'help' for usage.")
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
					fmt.Println("Insufficient arguments. 'help' for usage.")
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
		printHelp(os.Args, cmds)
		os.Exit(1)
	}
	cmd := os.Args[1]
	if cmd == "help" {
		printHelp(os.Args, cmds)
	} else if c, ok := cmds[cmd]; ok {
		if c.Function != nil {
			c.Function(os.Args)
		} else {
			fmt.Println("not implemented")
		}
	} else {
		printHelp(os.Args, cmds)
		os.Exit(1)
	}
}
