// Combining bool flags to condensed manner, e.g:
// cmd -s -o -m "message" will be: cmd -som "message"
// To run this example, open a terminal and execute the following lines:
// $ go install combineShortOptions.go
// $ combineShortOption short -som "message"; or,
// $ combineShortOption short --message "message"
// Note that the options `-som` are condesed into "one flag" 
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		// This option allow us to combining flag
		// Note: the use of -option flag will not work instead use --option
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:  "short",
				Usage: "complete a task on the list",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "serve", Aliases: []string{"s"}},
					&cli.BoolFlag{Name: "option", Aliases: []string{"o"}},
					&cli.BoolFlag{Name: "message", Aliases: []string{"m"}},
				},
				Action: func(cCtx *cli.Context) error {
					fmt.Println("serve: ", cCtx.Bool("serve"))
					fmt.Println("option: ", cCtx.Bool("option"))
					fmt.Println("message: ", cCtx.String("message"))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
