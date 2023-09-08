// An explicit exit code may be set by returning a non-nil error that fullfils
// "cli.ExitCoder", or a cli.MultiError
// To run this example, open a termina and execute de followin lines:
// $ go install exitCodes.g
// $ exitCodes --ginger-crouton
// Note the exit code 0
// $ exitCodes false
// No the error been shown "Ginger croutons are not in the soup"
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "ginger-crouton",
				Usage: "is it in the soup?",
			},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Println(cCtx.Bool("ginger-crouton"))
			if !cCtx.Bool("ginger-crouton") {
				return cli.Exit("Ginger croutons are not in the soup", 86)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
