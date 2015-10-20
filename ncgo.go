package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) {
		println("Hello friend!")
	}

	app.Run(os.Args)
}
