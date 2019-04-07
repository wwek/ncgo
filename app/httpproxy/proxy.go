package httpproxy

import (
	"io"
	"log"
	"net"
	"net/http"
	"strings"
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

// http
func (p *Pxy) HTTP(rw http.ResponseWriter, req *http.Request) {

	transport := http.DefaultTransport

	// 新建一个请求outReq
	outReq := new(http.Request)
	// 复制客户端请求到outReq上
	*outReq = *req // 复制请求

	//  处理匿名代理
	if p.Cfg.IsAnonymous == false {
		if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
			if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
				clientIP = strings.Join(prior, ", ") + ", " + clientIP
			}
			outReq.Header.Set("X-Forwarded-For", clientIP)
		}
	}

	// outReq请求放到传送上
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		rw.Write([]byte(err.Error()))
		return
	}

	// 回写http头
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}
	// 回写状态码
	rw.WriteHeader(res.StatusCode)
	// 回写body
	io.Copy(rw, res.Body)
	res.Body.Close()
}

// https
func (p *Pxy) HTTPS(rw http.ResponseWriter, req *http.Request) {

	// Step 1
	host := req.URL.Host
	hij, ok := rw.(http.Hijacker)
	if !ok {
		log.Printf("HTTP Server does not support hijacking")
	}

	client, _, err := hij.Hijack()
	if err != nil {
		return
	}

	// 连接远程
	server, err := net.Dial("tcp", host)
	if err != nil {
		return
	}
	client.Write([]byte("HTTP/1.0 200 Connection Established\r\n\r\n"))

	// 直通双向复制
	go io.Copy(server, client)
	go io.Copy(client, server)
}
