package httpproxy

type Pxy struct {
	Cfg Cfg
}

// 设置type
type Cfg struct {
	Addr        string   // 监听地址
	Port        string   // 监听端口
	IsAnonymous bool     // 高匿名模式
	BaseAuth    BaseAuth // 基础认证
}

// 实例化
func NewPxy() *Pxy {
	return &Pxy{
		Cfg: Cfg{
			Addr:        "",
			Port:        "8081",
			IsAnonymous: true,
			BaseAuth: BaseAuth{
				IsBaseAuth: false,
			},
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
	if cfg.BaseAuth.IsBaseAuth != p.Cfg.BaseAuth.IsBaseAuth {
		p.Cfg.BaseAuth.IsBaseAuth = cfg.BaseAuth.IsBaseAuth
	}
	if cfg.BaseAuth.UserName != "" && cfg.BaseAuth.PassWord != "" {
		p.Cfg.BaseAuth.UserName = cfg.BaseAuth.UserName
		p.Cfg.BaseAuth.PassWord = cfg.BaseAuth.PassWord
	}

}
