package main

import (
	"fmt"
	"github.com/armon/go-socks5"
	"github.com/wwek/ncgo/app/httpstat"
	"os"
	"time"

	"github.com/wwek/ncgo/app/httpproxy"
	"github.com/wwek/ncgo/app/speedtest"
	"github.com/wwek/ncgo/app/tcping"
	"github.com/wwek/ncgo/app/tcpscan"
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
			Name:    "tcping",
			Aliases: []string{"tp"},
			Usage:   "tcping 利用tcp端口三次握手来观察ping值",
			Action: func(c *cli.Context) error {
				cfg := &tcping.Cfg{}
				cfg.Host = c.String("host")
				cfg.Port = c.Int("port")
				cfg.Count = c.Int("count")
				cfg.TimeOut = c.Int("timeout")

				//fmt.Println(cfg)
				tcping.Run(cfg)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host,addr",
					Value: "127.0.0.1",
					Usage: "域名或IP",
				},
				cli.IntFlag{
					Name:  "port,p",
					Value: 80,
					Usage: "端口",
				},
				cli.IntFlag{
					Name:  "count,c",
					Value: 10,
					Usage: "tcping次数",
				},
				cli.IntFlag{
					Name:  "timeout,t",
					Value: 3,
					Usage: "连接超时时间默认3s",
				},
			},
		},
		{
			Name:    "httpstat",
			Aliases: []string{"ht"},
			Usage:   "httpstat 可以观察一个http请求各个环节的请求时间分布",
			Action: func(c *cli.Context) error {
				//fmt.Println(cfg)
				httpstat.HttpUrl = c.String("u")
				httpstat.HttpMethod = c.String("X")
				httpstat.PostBody = c.String("d")
				httpstat.FollowRedirects = c.Bool("L")
				httpstat.OnlyHeader = c.Bool("I")
				httpstat.Insecure = c.Bool("k")
				httpstat.HttpHeaders = c.StringSlice("H")
				httpstat.SaveOutput = c.Bool("O")
				httpstat.OutputFile = c.String("o")
				httpstat.ShowVersion = c.Bool("sv")
				httpstat.ClientCertFile = c.String("E")
				httpstat.FourOnly = c.Bool("4")
				httpstat.SixOnly = c.Bool("6")
				httpstat.Loop = c.Bool("lp")
				httpstat.LoopTime = c.Duration("lpt")
				httpstat.Run()
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "u",
					Value: "https://www.baidu.com:443",
					Usage: "HTTP URL",
				},
				cli.StringFlag{
					Name:  "X",
					Value: "GET",
					Usage: "HTTP method to use",
				},
				cli.StringFlag{
					Name:  "d",
					Value: "",
					Usage: "the body of a POST or PUT request; from file use @filename",
				},
				cli.BoolFlag{
					Name:   "L",
					Hidden: false,
					Usage:  "follow 30x redirects",
				},
				cli.BoolFlag{
					Name:   "I",
					Hidden: false,
					Usage:  "don't read body of request",
				},
				cli.StringSliceFlag{
					Name:  "H",
					Usage: "set HTTP header; repeatable: -H 'Accept: ...' -H 'Range: ...'",
				},
				cli.BoolFlag{
					Name:   "O",
					Hidden: false,
					Usage:  "save body as remote filename",
				},
				cli.StringFlag{
					Name:  "o",
					Value: "",
					Usage: "output file for body",
				},
				cli.BoolFlag{
					Name:   "sv",
					Hidden: false,
					Usage:  "print version number",
				},
				cli.StringFlag{
					Name:  "E",
					Value: "",
					Usage: "client cert file for tls config",
				},
				cli.BoolFlag{
					Name:   "4",
					Hidden: false,
					Usage:  "resolve IPv4 addresses only",
				},
				cli.BoolFlag{
					Name:   "6",
					Hidden: false,
					Usage:  "resolve IPv6 addresses only",
				},
				cli.BoolFlag{
					Name:   "lp",
					Hidden: false,
					Usage:  "开启循环运行",
				},
				cli.DurationFlag{
					Name:   "lpt",
					Value:  10 * time.Second,
					Hidden: false,
					Usage:  "循环运行 间隔时间",
				},
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
					Usage: "高匿名代理",
				},
				cli.StringFlag{
					Name:  "username,user",
					Value: "",
					Usage: "认证账号username",
				},
				cli.StringFlag{
					Name:  "password,pwd",
					Value: "",
					Usage: "认证密码password",
				},
				cli.StringFlag{
					Name:  "debug,d",
					Value: "false",
					Usage: "调试模式",
				},
			},
		},
		{
			Name:    "socksproxy",
			Aliases: []string{"socks5"},
			Usage:   "socks5代理服务器",
			Action: func(c *cli.Context) error {
				// Create a SOCKS5 server
				conf := &socks5.Config{}
				server, err := socks5.New(conf)
				if err != nil {
					panic(err)
				}

				// Create SOCKS5 proxy on localhost port 8000
				if err := server.ListenAndServe("tcp", "127.0.0.1:8000"); err != nil {
					panic(err)
				}
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
		{
			Name:    "tcpscan",
			Aliases: []string{"tp"},
			Usage:   "tcpscan tcp网络扫描",
			Action: func(c *cli.Context) error {
				cfg := &tcpscan.Cfg{}
				cfg.HostName = c.String("hostname")
				cfg.StartPort = c.Int("startport")
				cfg.EndPort = c.Int("endport")
				cfg.Timeout = c.Duration("timeout")

				//fmt.Println(cfg)
				tcpscan.Run(cfg)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "hostname,host",
					Value: "127.0.0.1",
					Usage: "域名或IP",
				},
				cli.IntFlag{
					Name:  "startport,sp",
					Value: 80,
					Usage: "开始端口",
				},
				cli.IntFlag{
					Name:  "endport,ep",
					Value: 100,
					Usage: "结束端口",
				},
				cli.DurationFlag{
					Name:  "timeout,t",
					Value: 1000 * 1000000,
					Usage: "连接超时时间默认 1000ms",
				},
			},
		},
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
	_ = app.Run(os.Args)
	//dargs := []string{"", "default"}
	//app.Run(os.Args)
}
