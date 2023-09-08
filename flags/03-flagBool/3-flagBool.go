// A simple example of flag bool
// To run this example, open a terminal and run the following lines
// $ go install 3-flagBool.go
// $ 3-flagBool --help
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var count int

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "foo",
				Usage: "foo greeting",
				Count: &count,
			},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Println("count", count)
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
