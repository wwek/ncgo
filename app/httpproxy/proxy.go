package httpproxy

import (
	"log"
	"net/http"
)

type Pxy struct {
	Cfg Cfg
}

// 设置type
type Cfg struct {
	Addr        string   // 监听地址
	Port        string   // 监听端口
	IsAnonymous bool     // 高匿名模式
	Debug       bool     // 调试模式
	BaseAuth    BaseAuth // 基础认证
}

// 实例化
func NewPxy() *Pxy {
	return &Pxy{
		Cfg: Cfg{
			Addr:        "",
			Port:        "8081",
			IsAnonymous: true,
			Debug:       false,
			BaseAuth: BaseAuth{
				IsBaseAuth: false,
			},
		},
	}
}

// 配置参数
func (p *Pxy) SetPxyCfg(cfg *Cfg) {
	if cfg.Addr != "" {
		p.Cfg.Addr = cfg.Addr
	}
	if cfg.Port != "" {
		p.Cfg.Port = cfg.Port
	}
	if cfg.IsAnonymous != p.Cfg.IsAnonymous {
		p.Cfg.IsAnonymous = cfg.IsAnonymous
	}
	if cfg.Debug != p.Cfg.Debug {
		p.Cfg.Debug = cfg.Debug
	}
	if cfg.BaseAuth.UserName != "" {
		p.Cfg.BaseAuth.UserName = cfg.BaseAuth.UserName
		p.Cfg.BaseAuth.IsBaseAuth = true
		if cfg.BaseAuth.PassWord != "" {
			p.Cfg.BaseAuth.PassWord = cfg.BaseAuth.PassWord
		} else {
			password := GetRandomString(4)
			p.Cfg.BaseAuth.PassWord = password
			log.Printf("httpproxy need BaseAuth username is:%s password is:%s", p.Cfg.BaseAuth.UserName, password)
		}

	}

}

// 运行代理服务
func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// debug
	if p.Cfg.Debug {
		log.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
		// fmt.Println(req)
	}

	// http && https
	if req.Method != "CONNECT" {
		// 处理http

		// http BaseAuth
		if p.Cfg.BaseAuth.IsBaseAuth {
			baseAuth := ChkBaseAuth(req,&p.Cfg)
			if false == baseAuth {
				rw.Header().Set("Proxy-Authorization", `Basic realm=my_realm`)
				rw.Header().Set("Proxy-Connection", `close`)
				rw.WriteHeader(401)
				_, _ = rw.Write([]byte("401 Unauthorized\n"))
				return
			}
		}

		p.HTTP(rw, req)
	} else {
		// 处理https

		// https BaseAuth
		if p.Cfg.BaseAuth.IsBaseAuth {
			baseAuth := ChkBaseAuth(req,&p.Cfg)
			if false == baseAuth {
				rw.Header().Set("Proxy-Authorization", `Basic realm=my_realm`)
				rw.Header().Set("Proxy-Connection", `close`)
				rw.WriteHeader(401)
				_, _ = rw.Write([]byte("401 Unauthorized\n"))
				return
			}
		}
		// 直通模式不做任何中间处理
		p.HTTPS(rw, req)
	}

}
