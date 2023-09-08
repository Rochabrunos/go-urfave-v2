// Sometimes it's useful to specify a flag's value within the usage string itself
// Such placeholders are incated with back quotes ``
// To run this examples, open a terminal and run the following lines:
// $ go install 4-placeholders.go
// $ 4-placeholder --help
// In the line of GLOBAL OPTION see the placeholder FILE, and 
// try to find where we define it on the code
// Note that only the fist placeholder is used. 
// Subsequente back-quoted words will be left as-is
package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				Aliases: []string{"c"},
				Usage: "Load configuration from  `FILE`",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}