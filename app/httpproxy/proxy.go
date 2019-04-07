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
	if cfg.BaseAuth.IsBaseAuth != p.Cfg.BaseAuth.IsBaseAuth {
		p.Cfg.BaseAuth.IsBaseAuth = cfg.BaseAuth.IsBaseAuth
	}
	if cfg.BaseAuth.UserName != "" && cfg.BaseAuth.PassWord != "" {
		p.Cfg.BaseAuth.UserName = cfg.BaseAuth.UserName
		p.Cfg.BaseAuth.PassWord = cfg.BaseAuth.PassWord
	}

}

// 运行代理服务
func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// debug
	if p.Cfg.Debug {
		log.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
		// fmt.Println(req)
	}

	if req.Method != "CONNECT" {
		p.HTTP(rw, req)
	} else {
		// 处理https
		// 直通模式不做任何中间处理
		p.HTTPS(rw, req)
	}

}
