// https://mp.weixin.qq.com/s/OhS_RQZojJbkenOSS_tEng
package main


func main() {
	hostname := flag.String("hostname", "", "hostname to test")
	startPort := flag.Int("start-port", 80, "the port on which the scanning starts")
	endPort := flag.Int("end-port", 100, "the port from which the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond * 200, "timeout")
	flag.Parse()

	ports := []int{}

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
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