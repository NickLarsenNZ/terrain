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
	err := config.Validate()

	if err != nil {
		fmt.Errorf("There were errors parsing the flags")
		return
	}

	terrain.Run(config)
}
