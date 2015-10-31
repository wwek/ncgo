package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/wwek/ncgo/app/speedtest"
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
		//a,err := speedtest.GetClientInfo()
	//	fmt.Printf("%s %s", a,err)
		b,err := speedtest.GetServerLists()
		fmt.Printf("%s %s", b,err)
	}

	app.Run(os.Args)
}
