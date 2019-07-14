package main

import (
	"fmt"
	terrain "github.com/nicklarsennz/terrain/internal"
)

func main() {
	// This is the main cli command `terrain`

	// Have the "FlagParser" "Parse" the commands flag and return a "Config"
	config := ParseFlags()

	// Validate the config
	errs := config.Validate()

	if len(errs) != 0 {
		fmt.Println("There were errors parsing the flags")
		for i, err := range errs {
			fmt.Printf("\t%d. %s\n", i+1, err)
		}
		return
	}

	terrain.Run(config)
}
