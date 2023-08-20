package main

import (
	"fmt"
	"keepass-tui/tui"
	"os"
)

func main() {
	if len(os.Args) != 2 && len(os.Getenv("KEEPASSDB")) == 0 {
		fmt.Println("Usage: <db>")
		fmt.Println("or use KEEPASSDB enviromental variable for db path")
		os.Exit(1)
	}

	dbPath := os.Getenv("KEEPASSDB")

	if len(os.Args) == 2 {
		dbPath = os.Args[1]
	}

	if err := tui.CreateTui(dbPath).Run(); err != nil {
		panic(err)
	}
}
