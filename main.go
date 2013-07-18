package main

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"log"
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
	db := Db{"localhost:27017", "test"}
	// 'help' is not in the map as it requires
	// the contents of the map.
	var cmds = map[string]Command{
		"store": {"store FILE [TAGS]",
			"Store file with optional comma-delimited tags.",
			func(args []string) {
				if len(args) < 3 {
					log.Fatal("Insufficient arguments. 'help' for usage.")
				}
				if err := db.Store(args[2]); err != nil {
					panic(err)
				}
			},
		},
		"get": {"get FILE",
			"Get file by name.",
			func(args []string) {
				if len(args) < 3 {
					log.Fatal("Insufficient arguments. 'help' for usage.")
				}
				if err := db.Get(args[2]); err != nil {
					log.Fatal("Error: ", err)
				}
			},
		},
		"smite": {"smite FILE",
			"Delete file.",
			func(args []string) {
				if len(args) < 3 {
					log.Fatal("Insufficient arguments. 'help' for usage.")
				}
				if err := db.Delete(args[2]); err != nil {
					log.Fatal("Error: ", err)
				}
			},
		},
		"ls": {"ls REGEX",
			"List all files with name matching regular expression REGEX (^...$ is implied).",
			func(args []string) {
				var query interface{}
				if len(args) > 2 {
					regex := "^" + args[2] + "$" // make the regex match the full filename
					query = bson.M{"filename": bson.RegEx{regex, ""}}
				}
				files, err := db.List(query)
				if err != nil {
					log.Fatal("Error: ", err)
				}
				for _, f := range files {
					fmt.Printf("%10d %16s %20s\n", f.Length,
						f.UploadDate.Format("2006-01-02 15:04"), f.Filename)
				}
			},
		},
		"tag": {"tag FILE TAGS",
			"Add TAGS (comma-delimited list) to FILE.", nil},
		"search": {"search QUERY",
			"Search for all files matching QUERY.", nil},
		"tour": {"tour",
			"Take the godm tour.",
			func(args []string) {
				fmt.Println("Welcome to the godm tour. Where can I get some godm bait?")
			},
		},
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
		log.Fatal("Unknown command. 'help' for usage.")
	}
}
