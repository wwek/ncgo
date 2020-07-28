# ncgo [![Build Status](https://travis-ci.org/wwek/ncgo.svg?branch=master)](https://travis-ci.org/wwek/ncgo)
ncgo 网络工具包

# 下载
* Github (Windows/Linux/Mac) - https://github.com/wwek/ncgo/releases
* Mirror (Windows/Linux/Mac) - 

# 使用
```bash
ncgo tcping www.baidu.com
ncgo httpstat https://www.baidu.com
ncgo httpstat -u www.qq.com -lp
```


## 帮助
ncgo help

# 源码安装
go get -u github.com/wwek/ncgo

# 功能清单

- [x] 网络带宽测速 SpeedTest Cli命令行模式
- [x] http 和 socks代理
- [x] tcping
- [x] httpstat
- [ ] 检查网络TCP UDP端口是否打开
- [ ] 扫描IP
- [ ] 反向tcp隧道
- [ ] Ping
- [ ] Mtr
- [ ] Whois IP或域名信息查询

# 感谢
github.com/davecheney/httpstat
