// Using a slice flag allow us to pass multiple values
// for a single flag 
// To run this code use: 
// In the terminal type: "go install 6-multipleValues.go"
// Type the command to help: "6-mutipleValues --help"
package main

import(
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)


func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			// The values will be provided as a slice:
			// (i) - Int64SliceFlag, (ii) - IntSliceFlag
			// (iii) - StringSliceFlag
			&cli.StringSliceFlag{
				Name: "greeting",
				Usage: "Pass mulitple greetings",
			},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Println(strings.Join(cCtx.StringSlice("greeting"), `, `))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}