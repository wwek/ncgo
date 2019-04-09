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


func ChkBaseAuth(w http.ResponseWriter, r *http.Request, cfg *Cfg) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ",2)
	if len(s) != 2 {
		return false
	}
	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	return pair[0] == cfg.BaseAuth.UserName && pair[1] == cfg.BaseAuth.PassWord
}