## 参考资料
* [HTTP 代理原理和实现](http://cizixs.com/2017/03/21/http-proxy-and-golang-implementation/``)
* [HTTP 隧道代理原理和实现](https://cizixs.com/2017/03/22/http-tunnel-proxy-and-golang-implementation/)
* https://github.com/elazarl/goproxy



## curl测试
```
 curl -x http://proxy:proxy@127.0.0.1:8080/ -I https://www.jd.com/
```