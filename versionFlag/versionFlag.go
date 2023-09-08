// The default version flag(-v/--version) is defined as cli.VersionFlag
// To run this example, open a terminal and execute the following lines:
// $ go install versionFlag.go
// $ versionFlag -V
package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		// Instead of -v flag, we can provide other names for version flag
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:    "my-app",
		Version: "v19.99.2",
	}

	app.Run(os.Args)
}
