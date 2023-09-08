// We can set alternate (or short) names for flags by providing a comma-delimited list for the Name
// To run this example, open the terminal and run the following lines:
// $ go install 5-alternateNames.go
// $ 5-alernateNames --help
// Note the short -l option to the --lang flag
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
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
