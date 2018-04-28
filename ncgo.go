package main

import (
	"fmt"
	"os"

	"github.com/wwek/ncgo/app/speedtest"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "ncgo"
	app.Usage = "网络工具包"
	app.Version = "v1.0"
	app.Copyright = `项目源码： https://github.com/wwek/ncgo
   作者：     wwek|流水理鱼
   作者博客： http://www.iamle.com`
	app.Commands = []*cli.Command{
		{
			Name:    "default",
			Aliases: []string{"df"},
			Usage:   "支持的命令",
			Action: func(c *cli.Context) error {
				fmt.Println("boom! I say!")
				return nil
			},
		},
		{
			Name:    "speedtest",
			Aliases: []string{"st"},
			Usage:   "speedtest.net网络带宽测速",
			Action: func(c *cli.Context) error {
				speedtest.Run()
				return nil
			},
		},
		// {
		// 	Name:    "netscan",
		// 	Aliases: []string{"ns"},
		// 	Usage:   "网络扫描/端口检查",
		// 	Action: func(c *cli.Context) {
		// 		netscan.Run()
		// 		//println("网络扫描: ", c.Args().First())
		// 	},
		// },
		// {
		// 	Name:    "dtest",
		// 	Aliases: []string{"dt"},
		// 	Usage:   "ddos压力测试",
		// 	Action: func(c *cli.Context) {
		// 		dtest.Run()
		// 	},
		// },
		// {
		// 	Name:    "complete",
		// 	Aliases: []string{"c"},
		// 	Usage:   "complete a task on the list",
		// 	Action: func(c *cli.Context) {
		// 		println("completed task: ", c.Args().First())
		// 	},
		// },
		// {
		// 	Name:    "template",
		// 	Aliases: []string{"r"},
		// 	Usage:   "options for task templates",
		// 	Subcommands: []cli.Command{
		// 		{
		// 			Name:  "add",
		// 			Usage: "add a new template",
		// 			Action: func(c *cli.Context) {
		// 				println("new task template: ", c.Args().First())
		// 			},
		// 		},
		// 		{
		// 			Name:  "remove",
		// 			Usage: "remove an existing template",
		// 			Action: func(c *cli.Context) {
		// 				println("removed task template: ", c.Args().First())
		// 			},
		// 		},
		// 	},
		// },
	}
	app.Run(os.Args)
	//dargs := []string{"", "default"}
	//app.Run(os.Args)
}
