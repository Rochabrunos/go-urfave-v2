// We can additionally query the environment query search for some of many flags
// The first environment variable that resolves is used
// To run this example, open a terminal and execute the following lines:
// $ go instal 9-environment2.go
// $ 9-environment2 --help
// Note that the same occurs with the --lang flag, 
// it is displayed a list of environment varible that will be queried.
// The fist value found will be used
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
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "english",
				Usage:   "language for the greeting",
				EnvVars: []string{
					"LEGACY_COMPACT_LANG",
					"APP_LANG",
					"LANG",
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
