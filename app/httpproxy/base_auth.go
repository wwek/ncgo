package httpproxy

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// BaseAuth http基础认证
type BaseAuth struct {
	IsBaseAuth bool   // 需要BaseAuth认证
	UserName   string // 账号
	PassWord   string // 密码
}

var proxyAuthorizationHeader = "Proxy-Authorization"

// 检查请求认证
// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Proxy-Authorization
func ChkBaseAuth(req *http.Request, cfg *Cfg) bool {
	authheader := strings.SplitN(req.Header.Get(proxyAuthorizationHeader), " ", 2)
	req.Header.Del(proxyAuthorizationHeader)
	if len(authheader) != 2 || authheader[0] != "Basic" {
		return false
	}
	userpassraw, err := base64.StdEncoding.DecodeString(authheader[1])
	if err != nil {
		return false
	}
	userpass := strings.SplitN(string(userpassraw), ":", 2)
	if len(userpass) != 2 {
		return false
	}

	return userpass[0] == cfg.BaseAuth.UserName &&  userpass[1] == cfg.BaseAuth.PassWord
}