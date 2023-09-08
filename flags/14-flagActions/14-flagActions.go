// Handlers can be registered per flag which are triggered after a flag has been processed.
// To run this example, open a terminal and execute the following lines:
// $ go install 14-flagActions.go
// $ 14-flagActions --port 65537
// Note the error, meaning we validate the flag in the moment we recieve the value for --port flag
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
			&cli.IntFlag{
				Name:        "port",
				Usage:       "Use a randomized port",
				Value:       0,
				DefaultText: "random",
				Action: func(ctx *cli.Context, v int) error {
					if v >= 65536 {
						return fmt.Errorf("Flag port value %v out of range[0-65536]", v)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
