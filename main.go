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
		flag.Usage()
		os.Exit(1)
	}

	cmd, args := args[0], args[1:]

	found := false
	for _, c := range commands {
		if cmd == c.Name {
			c.Run(db, args)
			found = true
			break
		}
	}

	if !found {
		log.Fatal("Unknown command. 'help' for usage.")
	}
}
