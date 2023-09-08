// A separate package ALTSRC that adds support for getting flag values from other file input sources
// Currently supported input source formats: YAML, JSON, TOML.
package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func main() {
	flags := []cli.Flag{
		altsrc.NewIntFlag(&cli.IntFlag{Name: "test"}),
		&cli.StringFlag{Name: "load"},
	}

	app := &cli.App{
		Action: func(cCtx *cli.Context) error {
			fmt.Println("--test value.*default:0")
			return nil
		},
		Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("load")),
		Flags:  flags,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Errorf("\a%v", err)
	}

}
