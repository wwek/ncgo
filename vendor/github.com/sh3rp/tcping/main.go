package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sh3rp/tcping/tcping"
)

var VERSION = "1.1"

var host string
var ports string
var iface string
var debug bool
var count int
var showVersion bool

func main() {
	flag.StringVar(&ports, "p", "80", "Port(s) to use for the TCP connection; for multiple ports, use a comma separated list")
	flag.BoolVar(&debug, "d", false, "Debug output packet sent and received")
	flag.IntVar(&count, "c", 0, "Number of probes to send")
	flag.StringVar(&iface, "i", "", "Interface to use as the source of the TCP packets")
	flag.BoolVar(&showVersion, "v", false, "Version info")
	flag.Parse()

	if showVersion {
		fmt.Printf("tcping v%s\n", VERSION)
		return
	}

	host = flag.Arg(0)

	if host == "" {
		fmt.Printf("Must supply a host to ping.\n")
		os.Exit(1)
	}

	src := tcping.GetInterface(iface)

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
}
