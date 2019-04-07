package httpproxy

import (
	"log"
	"net/http"
)

func Run(cfg *Cfg) {
	pxy := NewPxy()
	pxy.SetPxyCfg(cfg)
	log.Printf("httpPxoy is runing on %s:%s \n", cfg.Addr, cfg.Port)
	http.Handle("/", pxy)
	bindAddr := cfg.Addr + ":" + cfg.Port
	log.Fatalln(http.ListenAndServe(bindAddr, nil))
}
