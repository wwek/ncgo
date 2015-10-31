package speedtest

import (
	//"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	STCURL = "http://www.speedtest.net/speedtest-config.php"         //客户端信息
	STSS   = "http://www.speedtest.net/speedtest-servers-static.php" //SpeedTest 测速服务器列表
)

func ClientInfo() []byte {
	res, err := http.Get(STCURL)
	if err != nil {
		log.Fatal(err)
	}
	results, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fahital(err)
	}
	return results
}
