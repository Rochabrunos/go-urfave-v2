// For additional organization in apps that have many subcommands, we can associate
// a category for each command to gruop them together in the help output
// To execute this example, open a terminal and run the following lines:
// $ go install 02-categories.go
// $ 02-categories
// Note that "add" and "remove" options are in the same level of indentation,
// which means that they belong to the same category
package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "noop",
			},
			{
				Name:     "add",
				Category: "template",
			},
			{
				Name:     "remove",
				Category: "template",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
