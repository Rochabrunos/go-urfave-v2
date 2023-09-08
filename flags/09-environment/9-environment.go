// We can algo have the default value set from the environment via EnvVars
// To run this example, open a terminal and execute the following lines:
// $ go install 9-environment.go
// $ 9-environment --help
// Note in the end of the --lang flag line the [$APP_LANG] environment variable 
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
				Name: "lang",
				Aliases: []string{"1"},
				Value: "english",
				Usage: "language for the greeting",
				EnvVars: []string{"APP_LANG"},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}