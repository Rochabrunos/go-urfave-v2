// Using timestamp flag is simple
// To run this example, open a terminal and execute the following lines:
// $ go install timestampFlag.go
// $ timestampFlag --meeting 2023-08-12T15:21:05
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	brazilTime, _ := time.LoadLocation("America/Sao_Paulo")
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.TimestampFlag{Name: "meeting", Layout: "2006-01-02T15:04:05", Timezone: brazilTime},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Printf("%v", cCtx.Timestamp("meeting").Location().String())
			fmt.Println()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
