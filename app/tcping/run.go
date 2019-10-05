package tcping

import (
	//"flag"
	"fmt"
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
}
