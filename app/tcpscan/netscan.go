package tcpscan

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type TS struct {
	Cfg *Cfg
}

// 设置type
type Cfg struct {
	HostName  string // 地址
	StartPort int    // 开始端口
	EndPort   int    // 结束端口
	Timeout   time.Duration   // 超时时间
}

// 实例化
func NewTcpScan(c *Cfg) *TS {
	if c.HostName == "" {
		c.HostName = "127.0.0.1"
	}
	if c.StartPort == 0 {
		c.StartPort = 80
	}
	if c.EndPort == 0 {
		c.EndPort = 100
	}
	if c.Timeout == 0 {
		c.Timeout = time.Millisecond * 200
	}

	return &TS{
		Cfg: c,
	}
}

func (ts *TS) tcpscan() {

	ports := []int{}

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for port := ts.Cfg.StartPort; port <= ts.Cfg.EndPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(ts.Cfg.HostName, p, ts.Cfg.Timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}

	wg.Wait()
	fmt.Printf("opened ports: %v\n", ports)
}

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}

	return false
}
