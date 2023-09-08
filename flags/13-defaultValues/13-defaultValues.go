// Sometimes it's useful to specify a flag's default help-text value within the flag declaration
// The default value can be set via the DefaultText struct field
// To run this example, open a terminal executing the following lines:
// $ go install 13-defaultValues.go
// $ 13-defaultValues
// Note that port flag shows "default:random" as we defined before in the code
package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Usage:       "Use a randomized port",
				Value:       0,
				DefaultText: "random",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
