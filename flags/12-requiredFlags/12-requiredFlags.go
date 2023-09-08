// We can make a flag required by serring the Required field as trus
// If a user does not provide a required flag, they will be shown an error message
// To run this example, open a terminal and execute the following lines:
// $ go install 12-requiredFlags.go
// $ 12-requiredFlags
// Note, the last line indicating that you must provide a value for "lang" flag
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
			&cli.StringFlag{
				Name:     "lang",
				Value:    "english",
				Usage:    "language for the greeting",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			output := "Hello"
			if cCtx.String("lang") == "spanish" {
				output = "Hola"
			}
			fmt.Println(output)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
