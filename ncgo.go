package main

import (
	//"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/wwek/ncgo/app/speedtest"
)

func main() {
	app          := cli.NewApp()
	app.Name      = "ncgo"
	app.Usage     = "网络瑞士军刀，Golang打造"
	app.Version   = "v1.0"
	app.Copyright = `项目源码： https://github.com/wwek/ncgo
   作者：     wwek|流水理鱼
   作者博客： http://www.iamle.com`
	app.Commands = []cli.Command{
		{
			Name:    "speedtest",
			Aliases: []string{"st"},
			Usage:   "speedtest.net网络带宽测速",
			Action: func(c *cli.Context) {
				speedtest.Run()
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args().First())
			},
		},
		{
			Name:    "template",
			Aliases: []string{"r"},
			Usage:   "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) {
						println("new task template: ", c.Args().First())
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) {
						println("removed task template: ", c.Args().First())
					},
				},
			},
		},
	}


	app.Run(os.Args)
}
