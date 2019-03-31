package main

type Pxy struct {
	Cfg
}

// 设置type
type Cfg struct {
	Addr        string // 监听地址
	Port        string // 监听端口
	IsAnonymous bool   // 高匿名模式
	BaseAuth           // 基础认证
}

// 实例化
func NewPxy() *Pxy {
	return &Pxy{
		Cfg{
			Addr:        "0.0.0.0",
			Port:        "8081",
			IsAnonymous: true,
		},
	}
}

// 配置参数
func (p *Pxy) SetPxyCfg(cfg Cfg) {
	if cfg.Addr != "" {
		p.Cfg.Addr = cfg.Addr
	}
	if cfg.Port != "" {
		p.Cfg.Port = cfg.Port
	}
	if cfg.IsAnonymous != p.Cfg.IsAnonymous {
		p.Cfg.IsAnonymous = cfg.IsAnonymous
	}
	if cfg.IsBaseAuth != p.Cfg.IsBaseAuth {
		p.Cfg.IsBaseAuth = cfg.IsBaseAuth
	}
	if cfg.UserName != "" && cfg.PassWord != "" {
		p.Cfg.UserName = cfg.UserName
		p.Cfg.PassWord = cfg.PassWord
	}

}

func main() {

}
