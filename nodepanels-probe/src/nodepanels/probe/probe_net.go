package probe

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"io/ioutil"
	"net/http"
	"nodepanels/util"
)

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

	if ioCounters[0].BytesRecv > netInitRxBytes {
		probeUsage.Net.Rx = ioCounters[0].BytesRecv - netInitRxBytes
	} else {
		probeUsage.Net.Rx = 0
	}
	if ioCounters[0].BytesSent > netInitTxBytes {
		probeUsage.Net.Tx = ioCounters[0].BytesSent - netInitTxBytes
	} else {
		probeUsage.Net.Tx = 0
	}

	netInitRxBytes = ioCounters[0].BytesRecv
	netInitTxBytes = ioCounters[0].BytesSent

	return probeUsage

}
