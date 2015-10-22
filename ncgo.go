package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "ncgo"
	app.Usage = "网络瑞士军刀，Golang打造"
	app.Version = "v1.0"
	app.Copyright = `
	项目源码： https://github.com/wwek/ncgo

	wwek|流水理鱼
	http://www.iamle.com`
	app.Action = func(c *cli.Context) {
		name := "someone"
		if len(c.Args()) > 0 {
			name = c.Args()[0]
		}
		if c.String("lang") == "spanish" {
			println("Hola", name)
		} else {
			println("Hello", name)
		}
	}

	app.Run(os.Args)
}
