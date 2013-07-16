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
		Show help on the progrm or a command.
	add FILE [OPTIONS]
		Store file.
	get FILE
		Get file by name.
	smite FILE
		Delete file.
	search TERMS
		Search for a file.`)
}

func main() {
	if len(os.Args) < 2 {
		printHelp(os.Args)
		os.Exit(1)
	}
	db := godm.Db{"localhost:27017", "test"}
	cmd := os.Args[1]
	switch cmd {
	case "help":
		printHelp(os.Args)
	case "add":
		if len(os.Args) < 3 {
			printHelp(os.Args)
			os.Exit(1)
		}
		if err := db.Add(os.Args[2]); err != nil {
			panic(err)
		}
	case "get":
		if len(os.Args) < 3 {
			printHelp(os.Args)
			os.Exit(1)
		}
		if err := db.Get(os.Args[2]); err != nil {
			panic(err)
		}
	default:
		printHelp(os.Args)
	}
}
