package httpproxy

// BaseAuth http基础认证
type BaseAuth struct {
	IsBaseAuth bool   // 需要BaseAuth认证
	UserName   string // 账号
	PassWord   string // 密码
}
