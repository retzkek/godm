package main

import (
	"flag"
	"log"
	"os"
)

var (
	verbose = flag.Bool("v", true, "turn on verbose output")
)

var commands = []*Command{
	getCmd,
	lsCmd,
	smiteCmd,
	storeCmd,
}

var (
	db = &Db{"localhost:27017", "test"}
)

func main() {
	flag.Parse()
	args := flag.Args()

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
