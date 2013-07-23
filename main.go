package main

import (
	"flag"
	"log"
	"os"
)

var (
	verbose = flag.Bool("v", true, "turn on verbose output")
	host    = flag.String("host", "localhost:27017",
		"MongoDB database server. Can also set GODMHOST env variable.")
	database = flag.String("db", "test",
		"MongoDB database. Can also set GODMDB env variable.")
)

var commands = []*Command{
	getCmd,
	helpCmd,
	lsCmd,
	smiteCmd,
	storeCmd,
	tourCmd,
}

func findCommand(cmdName string) *Command {
	for _, c := range commands {
		if cmdName == c.Name {
			return c
		}
	}
	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()

	// check environment variables
	if *host == "localhost:27017" {
		if env := os.Getenv("GODMHOST"); env != "" {
			flag.Set("host", env)
		}
	}
	if *database == "test" {
		if env := os.Getenv("GODMDB"); env != "" {
			flag.Set("db", env)
		}
	}

	db := &Db{*host, *database}

	if len(args) == 0 {
		helpFunc(db, nil, args)
		os.Exit(1)
	}

	cmdName, args := args[0], args[1:]
	if cmd := findCommand(cmdName); cmd != nil {
		cmd.Run(db, args)
	} else {
		log.Fatal("Unknown command. 'help' for usage.")
	}
}
