package tcping

import (
	//"flag"
	"fmt"
<<<<<<< HEAD
	"net"
	"os"
)

//app 入口
func Run(cfg *Cfg) {

	//hostPtr := flag.String("host", "", "Host or IP address to test")
	//portPtr := flag.Int("port", 80, "Port number to query")
	//countPtr := flag.Int("count", 10, "Number of requests to send")
	//timeoutPtr := flag.Int("timeout", 1, "Timeout for each request, in seconds")
	//var host string
	//
	//flag.Parse()
	//
	//if len(os.Args) == 2 && os.Args[1][:1] != "-" {
	//	host = os.Args[1]
	//} else {
	//	host = *hostPtr
	//}
	//
	//port := *portPtr
	//count := *countPtr
	//timeout := *timeoutPtr
	//
	//if host == "" {
	//	flag.Usage()
	//	os.Exit(1)
	//}

	_, err := net.LookupIP(cfg.Host)

	if err != nil {
		fmt.Printf("error: can't resolve %s\n", cfg.Host)
		os.Exit(2)
	}

	ping(cfg.Host, cfg.Port, cfg.Count, cfg.TimeOut)
=======
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sh3rp/tcping/tcping"
	"gopkg.in/urfave/cli.v1"
)

var VERSION = "1.1"

var host string
var ports string
var iface string
var debug bool
var count int
var showVersion bool

func Run(c *cli.Context) {

	host = c.Args()[0]
	//fmt.Println(host)
	ports = c.String("p")
	iface = c.String("i")
	debug = c.Bool("d")
	count = c.Int("c")
	showVersion = c.Bool("v")

	if host == "" {
		fmt.Printf("Must supply a host to ping.\n")
		os.Exit(1)
	}

	if showVersion {
		fmt.Printf("tcping v%s\n", VERSION)
		return
	}

	src := tcping.GetInterface(iface)

	//fmt.Println(src)

	probe := tcping.NewProbe(src, host, debug)

	strPorts := strings.Split(ports, ",")

	var portList []int

	for _, p := range strPorts {
		prt, err := strconv.Atoi(p)

		if err == nil {
			portList = append(portList, prt)
		}
	}

	if debug {
		fmt.Printf("Src IP: %s\n\n", src)
	}

	if count > 0 {
		for i := 0; i < count; i++ {
			for _, p := range portList {
				sendProbe(probe, p)
			}
			time.Sleep(1 * time.Second)
		}
	} else {
		for {
			for _, p := range portList {
				sendProbe(probe, p)
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func sendProbe(probe tcping.Probe, port int) {
	latency := probe.GetLatency(uint16(port))
	if latency > 0 {
		fmt.Printf("%-15s -> %-15s %d ms\n",
			probe.SrcIP,
			probe.DstIP,
			latency/int64(time.Millisecond))
	} else {
		fmt.Printf("Timeout\n")
	}
>>>>>>> 230edbfd0c6030d382566735eeadeb449aa46dbc
}
