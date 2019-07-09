package main

import (
	"fmt"
	cli "github.com/urfave/cli"
	"os"
)

func main() {
	// See http://securitygobyexample.com/urfave-cli-subcommands

	var resFlag bool
	var outFlag string

	app := cli.NewApp()
	app.Name = "terrain"
	app.Usage = "terraform config info"
	app.HideVersion = true

	app.Authors = []cli.Author{
		{"Nick Larsen", "nick@aptiv.co.nz"},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "resources, r",
			Usage:       "List the Terraform resources declared",
			Destination: &resFlag,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "Output format: [text, json, markdown]",
			Value:       "text",
			Destination: &outFlag,
		},
	}

	app.Action = func(c *cli.Context) error {
		if resFlag {
			fmt.Println("Resources will be listed")
		}

		switch outFlag {
		case "json":
			fmt.Println("Gonna output json")
		case "markdown":
			fmt.Println("Gonna output markdown")
		default: // Todo: should this error, or just default to text?
			fmt.Println("Gonna output text")
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Errorf("unable to run command line app: %s", err)
	}

}
