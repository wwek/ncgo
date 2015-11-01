package speedtest

import (
	"fmt"
	//"io/ioutil"
	//"log"
	"net/http"
	"time"
		"encoding/xml"
	//"github.com/kellydunn/golang-geo"
)

const (
	clientConfigUrl  string        = `http://www.speedtest.net/speedtest-config.php` //客户端信息
	serversConfigUrl string        = `http://www.speedtest.net/speedtest-servers-static.php` //SpeedTest 测速服务器列表
	getTimeout       time.Duration = 20 * time.Second //定义超时时间
)

//定义客户端信息结构体
type clientSpeedtestConfig struct {
	XMLName       xml.Name  `xml:"settings"`
	License       string    `xml:"licensekey"`
	ClientConfig  ccconfig  `xml:"client"`
	ServerConfig  csconfig  `xml:"server-config"`
}

type ccconfig struct {
	XMLName  xml.Name `xml:"client"`
	Ip       string   `xml:"ip,attr"`
	Lat      float64  `xml:"lat,attr"`
	Long     float64  `xml:"lon,attr"`
	ISP      string   `xml:"isp,attr"`
	ISPUpAvg uint     `xml:"ispulavg,attr"`
	ISPDlAvg uint     `xml:"ispdlavg,attr"`
}

type csconfig struct {
	XMLName   xml.Name `xml:"server-config"`
	Threads   int      `xml:"threadcount,attr"`
	IgnoreIDs string   `xml:"ignoreids,attr"`
}

//定义服务器结构体
type serverSettings struct {
	XMLName xml.Name `xml:"settings"`
	Servers []server `xml:"servers>server"`
}
type server struct {
	XMLName xml.Name `xml:"server"`
	Url     string   `xml:"url,attr"`
	Url2    string   `xml:"url2,attr"`
	Lat     float64  `xml:"lat,attr"`
	Long    float64  `xml:"lon,attr"`
	Name    string   `xml:"name,attr"`
	Country string   `xml:"country,attr"`
	CC      string   `xml:"cc,attr"`
	Sponsor string   `xml:"sponsor,attr"`
	ID      uint     `xml:"id,attr"`
	Host    string   `xml:"host,attr"`
}


//通过httpget访问clientConfigUrl得到客户端信息
func GetClientInfo() (*clientSpeedtestConfig, error) {
	client := http.Client{
		Timeout : getTimeout, //设置超时时间
	}
	req, err := http.NewRequest("GET", clientConfigUrl, nil)
	if err != nil {
		return nil,err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http状态错误 %s", resp.StatusCode)
	}
	xmlDec := xml.NewDecoder(resp.Body)
	cstc := clientSpeedtestConfig{}
	if err := xmlDec.Decode(&cstc); err != nil {
		return nil,err
	}
	return &cstc, nil
}

//通过httpget访问serversConfigUrl得到服务器列表信息
func GetServerLists() (*serverSettings, error){
	client := http.Client {
		Timeout : getTimeout,
	}
	req, err := http.NewRequest("GET", serversConfigUrl, nil)
	if err !=nil {
		return nil,err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http状态错误 %s", resp.StatusCode)
	}
	xmlDec := xml.NewDecoder(resp.Body)
	ss := serverSettings{}
	if err := xmlDec.Decode(&ss); err != nil {
		return nil, err
	}
	return &ss, nil
	
}

func Run() {
	    a,err := GetClientInfo()
	    fmt.Printf("%s %s", a,err)
		b,err := GetServerLists()
		fmt.Printf("%s %s", b,err)
}