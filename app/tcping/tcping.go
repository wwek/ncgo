package tcping

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"net"
	"os"
	"time"
)

// 配置
type Cfg struct {
	Host    string // 域名或者ip
	Port    int    // 端口
	Count   int    // 次数
	TimeOut int    // 超时
}

func NewCfg() *Cfg {
	return &Cfg{
		Host:    "127.0.0.1",
		Port:    80,
		Count:   10,
		TimeOut: 1,
	}
}

func ping(host string, port int, count int, timeout int) {
	successfulProbes := 0
	i := 1
	timeTotal := time.Duration(0)
	var responseTimes []float64

	addr := fmt.Sprintf("%s:%d", host, port)

	for i = 1; count >= i; i++ {
		timeStart := time.Now()
		_, err := net.DialTimeout("tcp", addr, time.Second*time.Duration(timeout))
		responseTime := time.Since(timeStart)
		if err != nil {
			fmt.Println(fmt.Sprintf("Received timeout while connecting to %s on port %d.", host, port))
		} else {
			fmt.Println(fmt.Sprintf("Probe %v: Connected to %s:%d, RTT=%.2fms", i, host, port, float32(responseTime)/1e6))
			timeTotal += responseTime
			successfulProbes++
			responseTimes = append(responseTimes, float64(responseTime))
		}

		time.Sleep(time.Second - responseTime)
	}

	// Print results
	output(successfulProbes, timeTotal, host, port, responseTimes, i)

}

func output(successfulProbes int, timeTotal time.Duration, host string, port int, responseTimes []float64, i int) {
	// Let's calculate and spill some results
	// 1. Average response time
	timeAverage := time.Duration(1)
	if successfulProbes > 0 {
		timeAverage = time.Duration(int64(timeTotal) / int64(successfulProbes))
	} else {
		fmt.Printf("\nAll the requests have failed. The host %s is not replying to connections on %d\n", host, port)
		os.Exit(1)
	}
	// 2. Min and Max response times
	var biggest float64

	smallest := float64(1000000000)

	for _, v := range responseTimes {

		if v > biggest {
			biggest = v
		}

		if v < smallest {
			smallest = v
		}

	}

	// 3. Median response time
	median, _ := stats.Median(responseTimes)

	// 4. Percentile
	percentile90, _ := stats.Percentile(responseTimes, float64(90))
	percentile75, _ := stats.Percentile(responseTimes, float64(75))
	percentile50, _ := stats.Percentile(responseTimes, float64(50))
	percentile25, _ := stats.Percentile(responseTimes, float64(25))

	fmt.Println("\nProbes sent:", i-1, "\nSuccessful responses:", successfulProbes,
		"\n% of requests failed:", float64(100-(successfulProbes*100)/(i-1)),
		"\nMin response time:", time.Duration(smallest),
		"\nAverage response time:", timeAverage,
		"\nMedian response time:", time.Duration(median),
		"\nMax response time:", time.Duration(biggest))

	fmt.Println("\n90% of requests were faster than:", time.Duration(percentile90),
		"\n75% of requests were faster than:", time.Duration(percentile75),
		"\n50% of requests were faster than:", time.Duration(percentile50),
		"\n25% of requests were faster than:", time.Duration(percentile25))

}
