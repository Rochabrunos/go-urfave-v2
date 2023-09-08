package main

import (
	"fmt"
	"os"
	"log"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Action: func(cCtx *cli.Context) error {
			fmt.Printf("Hello %q\n", cCtx.Args().Get(0))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil { log.Fatal(err) }
}