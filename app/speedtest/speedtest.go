package speedtest

import (
	"errors"
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"io"
	"os"
	"encoding/xml"
	"net/http"
	"time"
	"github.com/kellydunn/golang-geo"
	"github.com/Bowery/prompt"
	"github.com/bndr/gotabulate"
	"github.com/joliv/spark"
)

const (
	clientConfigUrl  string        = `http://www.speedtest.net/speedtest-config.php`         //客户端信息
	serversConfigUrl string        = `http://www.speedtest.net/speedtest-servers-static.php` //SpeedTest 测速服务器列表
	getTimeout       time.Duration = 30 * time.Second                                        //定义超时时间
)

const (
	tableFormat      = "simple"
	maxFailureCount  = 3
	initialTestCount = 5
	basePingCount    = 5
	fullTestCount    = 20
	speedtestDuration = 3 
)



//定义客户端信息结构体
type clientSpeedtestConfig struct {
	XMLName      xml.Name `xml:"settings"`
	License      string   `xml:"licensekey"`
	ClientConfig ccconfig `xml:"client"`
	ServerConfig csconfig `xml:"server-config"`
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

type Config struct {
	LicenseKey string
	IP         net.IP
	Lat        float64
	Long       float64
	ISP        string
	Servers    []Testserver
}

type Testserver struct {
	Name     string
	Sponsor  string
	Country  string
	Lat      float64
	Long     float64
	Distance float64 //距离服务器的距离单位KM
	URLs     []string
	Host     string
	Latency  time.Duration //latency in ms
}
type testServerlist []Testserver



//通过httpget访问clientConfigUrl得到客户端信息
func GetClientInfo() (*clientSpeedtestConfig, error) {
	client := http.Client{
		Timeout: getTimeout, //设置超时时间
	}
	req, err := http.NewRequest("GET", clientConfigUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http状态错误 %s", resp.StatusCode)
	}
	xmlDec := xml.NewDecoder(resp.Body)
	cstc := clientSpeedtestConfig{}
	if err := xmlDec.Decode(&cstc); err != nil {
		return nil, err
	}
	return &cstc, nil
}

//通过httpget访问serversConfigUrl得到服务器列表信息
func GetServerLists() (*serverSettings, error) {
	client := http.Client{
		Timeout: getTimeout,
	}
	req, err := http.NewRequest("GET", serversConfigUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
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

func populateServers(cfg *Config, srvs []server, ignore map[uint]bool) error {
	for i := range srvs {
		//checking if we are ignoring this server
		_, ok := ignore[srvs[i].ID]
		if ok {
			continue
		}
		srv := Testserver{
			Name:    srvs[i].Name,
			Sponsor: srvs[i].Sponsor,
			Country: srvs[i].Country,
			Lat:     srvs[i].Lat,
			Long:    srvs[i].Long,
			Host:    srvs[i].Host,
		}
		if srvs[i].Url != "" {
			srv.URLs = append(srv.URLs, srvs[i].Url)
		}
		if srvs[i].Url2 != "" {
			srv.URLs = append(srv.URLs, srvs[i].Url2)
		}
		p := geo.NewPoint(cfg.Lat, cfg.Long)
		if p == nil {
			return errors.New("Invalid client lat/long")
		}
		sp := geo.NewPoint(srvs[i].Lat, srvs[i].Long)
		if sp == nil {
			return errors.New("Invalid server lat/long")
		}
		srv.Distance = p.GreatCircleDistance(sp)
		cfg.Servers = append(cfg.Servers, srv)
	}
	sort.Sort(testServerlist(cfg.Servers))
	return nil
}

func (tsl testServerlist) Len() int           { return len(tsl) }
func (tsl testServerlist) Swap(i, j int)      { tsl[i], tsl[j] = tsl[j], tsl[i] }
func (tsl testServerlist) Less(i, j int) bool { return tsl[i].Distance < tsl[j].Distance }


func testLatency(server Testserver) error {
	//perform a full latency test
	durs, err := server.Ping(fullTestCount)
	if err != nil {
		return err
	}
	var avg, max, min uint64
	var latencies []float64
	for i := range durs {
		ms := uint64(durs[i].Nanoseconds() / 1000000)
		latencies = append(latencies, float64(ms))
		avg += ms
		if ms > max {
			max = ms
		}
		if ms < min || min == 0 {
			min = ms
		}
	}
	avg = avg / uint64(len(durs))
	median := durs[len(durs)/2].Nanoseconds() / 1000000
	sparkline := spark.Line(latencies)
	fmt.Printf("Latency: %s\t%dms avg\t%dms median\t%dms max\t%dms min\n", sparkline, avg, median, max, min)
	return nil
}

func testDownstream(server Testserver) error {
	bps, err := server.Downstream(speedtestDuration)
	if err != nil {
		return err
	}
	fmt.Printf("Download: %s\n", HumanSpeed(bps))
	return nil
}

func testUpstream(server Testserver) error {
	bps, err := server.Upstream(speedtestDuration)
	if err != nil {
		return err
	}
	fmt.Printf("Upload:   %s\n", HumanSpeed(bps))
	return nil
}

func fullTest(server Testserver) error {
	if err := testLatency(server); err != nil {
		return err
	}
	if err := testDownstream(server); err != nil {
		return err
	}
	if err := testUpstream(server); err != nil {
		return err
	}
	return nil
}

func autoGetTestServers(cfg *Config) ([]Testserver, error) {
	//get the first 5 closest servers
	testServers := []Testserver{}
	failures := 0
	for i := range cfg.Servers {
		if failures >= maxFailureCount {
			if len(testServers) > 0 {
				return testServers, nil
			}
			return nil, fmt.Errorf("Failed to perform latency test on closest servers\n")
		}
		if len(testServers) >= initialTestCount {
			return testServers, nil
		}
		//get a latency from the server, the last latency will also be store in the
		//server structure
		if _, err := cfg.Servers[i].MedianPing(basePingCount); err != nil {
			failures++
			continue
		}
		testServers = append(testServers, cfg.Servers[i])
	}
	return testServers, nil
}

func getSearchServers(cfg *Config, query string) ([]Testserver, error) {
	//get the first 5 closest servers
	testServers := []Testserver{}
	for i := range cfg.Servers {
		if strings.Contains(strings.ToLower(cfg.Servers[i].Name), strings.ToLower(query)) {
			testServers = append(testServers, cfg.Servers[i])
		}
	}
	if len(testServers) == 0 {
		return nil, errors.New("no servers found")
	}
	return testServers, nil
}

func Run() {
	fmt.Printf("开始测速...\n")
	fmt.Printf("获取客户端信息和服务器列表开始...\n")
	cc, err := GetClientInfo()
	if err != nil {
		fmt.Printf("错误： %v\n", err)
		return
	}
	//fmt.Printf("%s %s", cc, err)
	srvs, err := GetServerLists()
	if err != nil {
		fmt.Printf("错误： %v\n", err)
		return
	}
	fmt.Printf("获取客户端信息和服务器列表完成...\n")
	cfg := Config{
		LicenseKey: cc.License,
		IP:         net.ParseIP(cc.ClientConfig.Ip),
		Lat:        cc.ClientConfig.Lat,
		Long:       cc.ClientConfig.Long,
		ISP:        cc.ClientConfig.ISP,
	}
	ignoreIDs := make(map[uint]bool, 1)
	strIDs := strings.Split(cc.ServerConfig.IgnoreIDs, ",")
	for i := range strIDs {
		x, err := strconv.ParseUint(strIDs[i], 10, 32)
		if err != nil {
			continue
		}
		ignoreIDs[uint(x)] = false
	}
	if err := populateServers(&cfg, srvs.Servers, ignoreIDs); err != nil {
		fmt.Printf("错误： %v\n", err)
		return
	}
	
		if len(cfg.Servers) <= 0 {
		fmt.Printf("没找到合适的测速服务器！\n")
		return
	}
	var headers []string
	var data [][]string
	var testServers []Testserver
//	if *search == "" {
		fmt.Printf("获取距离最近的服务器...\n")
		if testServers, err = autoGetTestServers(&cfg); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(-1)
		}
		fmt.Printf("%d Closest responding servers:\n", len(testServers))
		for i := range testServers {
			data = append(data, []string{fmt.Sprintf("%d", i),
				testServers[i].Name, testServers[i].Sponsor,
				fmt.Sprintf("%.02f", testServers[i].Distance),
				fmt.Sprintf("%s", testServers[i].Latency)})
		}
		headers = []string{"ID", "Name", "Sponsor", "Distance (km)", "Latency (ms)"}
//	} else {
//		if testServers, err = getSearchServers(cfg, *search); err != nil {
//			fmt.Fprintf(os.Stderr, "%s\n", err)
//			os.Exit(-1)
//		}
//		headers = []string{"ID", "Name", "Sponsor", "Distance (km)"}
//		fmt.Printf("%d Matching servers:\n", len(testServers))
//		for i := range testServers {
//			data = append(data, []string{fmt.Sprintf("%d", i),
//				testServers[i].Name, testServers[i].Sponsor,
//				fmt.Sprintf("%.02f", testServers[i].Distance)})
//		}

//	}
	t := gotabulate.Create(data)
	t.SetHeaders(headers)
	t.SetWrapStrings(false)
	fmt.Printf("%s", t.Render(tableFormat))
	fmt.Printf("Enter server ID for bandwidth test, or \"quit\" to exit\n")
	for {
		s, err := prompt.Basic("ID> ", true)
		if err != nil {
			fmt.Printf("input failure \"%v\"\n", err)
			os.Exit(-1)
		}
		//be REALLY forgiving on exit logic
		if strings.HasPrefix(strings.ToLower(s), "exit") {
			os.Exit(0)
		}
		if strings.HasPrefix(strings.ToLower(s), "quit") {
			os.Exit(0)
		}

		//try to convert the string to a number
		id, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\"%s\" is not a valid id\n", s)
			continue
		}
		if id > uint64(len(testServers)) {
			fmt.Fprintf(os.Stderr, "No server with ID \"%d\" available\n", id)
			continue
		}
		if err = fullTest(testServers[id]); err != nil {
			if err == io.EOF {
				fmt.Fprintf(os.Stderr, "Error, the remote server kicked us.\n")
				fmt.Fprintf(os.Stderr, "Maximum request size may have changed\n")
			} else {
				fmt.Fprintf(os.Stderr, "Test failed with unknown 错误： %v\n", err)
			}
			os.Exit(-1)
		} else {
			break //we are done
		}
	}
}
