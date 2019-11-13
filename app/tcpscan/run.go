package tcpscan

func Run(cfg *Cfg) {
	ts := NewTcpScan(cfg)
	ts.tcpscan()
}
