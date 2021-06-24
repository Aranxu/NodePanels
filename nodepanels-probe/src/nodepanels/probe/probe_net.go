package probe

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"io/ioutil"
	"net/http"
	"nodepanels/config"
	"nodepanels/util"
	"strings"
)

func InitConfigIp() {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Init config ip error : " + fmt.Sprintf("%s", err))
		}
	}()

	url := "http://ip-api.com/json?fields=status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,isp,org,as,asname,query&lang=zh-CN"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	detailInfo := make(map[string]string)
	json.Unmarshal(body, &detailInfo)

	//设置域名前缀
	prefixDomain := "america"
	if strings.Index(detailInfo["continent"], "亚洲") > -1 {
		prefixDomain = "asian"
	} else if strings.Index(detailInfo["continent"], "欧洲") > -1 {
		prefixDomain = "europe"
	} else if strings.Index(detailInfo["continent"], "美洲") > -1 {
		prefixDomain = "america"
	}

	util.LogInfo("This server will connect to " + prefixDomain + " datacenter")

	config.AgentUrl = prefixDomain + "-" + config.AgentUrl
	config.WsUrl = prefixDomain + "-" + config.WsUrl
	config.ApiUrl = prefixDomain + "-" + config.ApiUrl

}

func GetNetInfo(info ProbeInfo) ProbeInfo {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get net info error : " + fmt.Sprintf("%s", err))
		}
	}()

	ioCounters, _ := net.IOCounters(false)
	//初始化流量信息
	netInitRxBytes = ioCounters[0].BytesRecv
	netInitTxBytes = ioCounters[0].BytesSent

	url := "http://ip-api.com/json?fields=status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,isp,org,as,asname,query&lang=zh-CN"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	detailInfo := IpInfo{}
	json.Unmarshal(body, &detailInfo)

	privateIp := util.GetPrivateIp()

	info.NetInfo.PublicIp = detailInfo.Query
	info.NetInfo.PrivateIp = privateIp
	info.NetInfo.Dns = util.GetDns(privateIp)
	info.NetInfo.DetailInfo = detailInfo

	info.NetInfo.AgentIp = util.GetAgentIp()
	info.NetInfo.ApiIp = util.GetApiIp()

	return info
}

func GetNetUsage(probeUsage ProbeUsage) ProbeUsage {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("get net usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	ioCounters, _ := net.IOCounters(false)

	probeUsage.Net.Rx = ioCounters[0].BytesRecv - netInitRxBytes
	probeUsage.Net.Tx = ioCounters[0].BytesSent - netInitTxBytes

	netInitRxBytes = ioCounters[0].BytesRecv
	netInitTxBytes = ioCounters[0].BytesSent

	return probeUsage

}
