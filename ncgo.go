package main

import (
	"fmt"
	"os"

	"github.com/wwek/ncgo/app/httpproxy"
	"github.com/wwek/ncgo/app/speedtest"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "ncgo"
	app.Usage = "网络工具包"
	app.Version = "v1.0"
	app.Copyright = `项目源码： https://github.com/wwek/ncgo
   作者：     wwek|流水理鱼
   作者博客： http://www.iamle.com`
	app.Commands = []cli.Command{
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
			Name:    "httpproxy",
			Aliases: []string{"hp"},
			Usage:   "http/https代理服务器",
			Action: func(c *cli.Context) error {
				cfg := &httpproxy.Cfg{}
				cfg.Addr = c.String("addr")
				cfg.Port = c.String("port")
				cfg.IsAnonymous = c.Bool("anonymous")
				cfg.BaseAuth.UserName = c.String("user")
				cfg.BaseAuth.PassWord = c.String("pwd")
				cfg.Debug = c.Bool("debug")
				// fmt.Println(cfg)
				httpproxy.Run(cfg)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "port,p",
					Value: "8080",
					Usage: "监听端口",
				},
				cli.StringFlag{
					Name:  "addr,a",
					Value: "0.0.0.0",
					Usage: "监听IP",
				},
				cli.StringFlag{
					Name:  "anonymous,n",
					Value: "true",
					Usage: "high匿名",
				},
				cli.StringFlag{
					Name:  "username,user",
					Value: "",
					Usage: "username",
				},
				cli.StringFlag{
					Name:  "password,pwd",
					Value: "",
					Usage: "password",
				},
				cli.StringFlag{
					Name:  "debug,d",
					Value: "false",
					Usage: "调试模式(默认否)",
				},
			},
		},
		{
			Name:    "socksproxy",
			Aliases: []string{"s5p"},
			Usage:   "socks5代理服务器",
			Action: func(c *cli.Context) error {
				//httpproxy.Run()
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "port,p",
					Value: "1080",
					Usage: "监听端口",
				},
				cli.StringFlag{
					Name:  "bind,b",
					Value: "0.0.0.0",
					Usage: "监听IP",
				},
			},
		},
		{
			Name:    "httpfile",
			Aliases: []string{"hf"},
			Usage:   "基于http的文件下载和上传",
			Action: func(c *cli.Context) error {
				fmt.Println("boom httpfile!")
				fmt.Println(c)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "port,p",
					Value: "8080",
					Usage: "监听端口",
				},
				cli.StringFlag{
					Name:  "docroot,dr",
					Value: "./",
					Usage: "目录",
				},
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
