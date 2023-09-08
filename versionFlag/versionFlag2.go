// Alternatively, the version printer at cli.VersionPrinter may be overriden
// To run this example, open a terminal and execute the following lines:
// $ go install versionFlag2.go
// $ versionFlag2 -v/--version
package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	Revision = "foobar"
)

func main() {
	app := &cli.App{
		Name: "my-app2",
		Version: "v19.99.0",
	}

	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("version=%s revision=%s\n", cCtx.App.Version, Revision)
	}

	app.Run(os.Args)
}