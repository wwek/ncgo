package httpstat

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

//app 入口
func Run() {

	if ShowVersion {
		fmt.Printf("%s %s (runtime: %s)\n", os.Args[0], Version, runtime.Version())
		os.Exit(0)
	}

	if FourOnly && SixOnly {
		fmt.Fprintf(os.Stderr, "%s: Only one of -4 and -6 may be specified\n", os.Args[0])
		os.Exit(-1)
	}

	if (HttpMethod == "POST" || HttpMethod == "PUT") && PostBody == "" {
		log.Fatal("must supply post body using -d when POST or PUT is used")
	}

	if OnlyHeader {
		HttpMethod = "HEAD"
	}

	url := parseURL(HttpUrl)

	visit(url)

}
