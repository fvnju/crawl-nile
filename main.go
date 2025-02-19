package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/fvnju/crawl-nile/internals/logic"
)

func main() {
	var isTimed bool
	flag.BoolVar(&isTimed, "t", false, "Measure execution time")
	start := time.Now()

	var isCliMode bool
	flag.BoolVar(&isCliMode, "c", false, "Enable CLI mode")

	flag.Parse()

	// Get the remaining arguments (username and password)
	args := flag.Args()

	if len(args) < 2 || !isCliMode {
		fmt.Printf("Usage: %s [-c] <username> <password>\n", os.Args[0])
		fmt.Printf("\nUse \"%s -h\" for help\n", os.Args[0])
		os.Exit(1)
	}

	username := args[0]
	password := args[1]

	logic.RunCli(username, password)

	if isTimed {
		fmt.Println("\nExecution time:", time.Since(start))
	}
}
