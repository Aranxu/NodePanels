package util

import (
	"fmt"
	"net"
	"nodepanels/config"
	"os/exec"
	"runtime"
	"strings"
)

func GetPrivateIp() string {

	defer func() {
		err := recover()
		if err != nil {
			LogError("Get private ip error : " + fmt.Sprintf("%s", err))
		}
	}()

	conn, _ := net.Dial("udp", "baidu.com:80")
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0] + "" //不加空的话入库时会变成其他值
}

func GetDns(privateIp string) string {

	defer func() {
		err := recover()
		if err != nil {
			LogError("Get DNS error : " + fmt.Sprintf("%s", err))
		}
	}()

	result := ""

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", "ipconfig /all")
		output, _ := cmd.Output()

		processStrs := strings.Split(string(output), privateIp)[1]
		processStrs = strings.Split(processStrs, "TCPIP")[0]
		processStrs = strings.Split(processStrs, "DNS")[1]
		processStrs = strings.Split(processStrs, ":")[1]
		processStrs = strings.Replace(processStrs, " ", "", -1)
		processStrs = strings.Replace(processStrs, "\r", "", -1)
		dnsStrs := strings.Split(processStrs, "\n")
		for _, dnsStr := range dnsStrs {
			if dnsStr != "" {
				result += dnsStr + ","
			}
		}
	}

	if runtime.GOOS == "linux" {
		cmd := exec.Command("sh", "-c", "cat /etc/resolv.conf")
		output, _ := cmd.Output()

		processStrs := strings.Split(string(output), "\n")
		for _, processStr := range processStrs {
			if strings.Contains(processStr, "nameserver") {
				result += strings.Split(processStr, "nameserver ")[1] + ","
			}
		}
	}

	return result[0 : len(result)-1]

}

func GetApiIp() string {
	return Get("https://" + config.ApiUrl + "/api/getApiIp")
}

func GetWsIp() string {
	return Get("https://" + config.WsUrl + "/api/getWsIp")
}

func GetAgentIp() string {
	return Get("https://" + config.AgentUrl + "/api/getAgentIp")
}

func GetDomainIp(domain string) string {
	conn, _ := net.Dial("ip:icmp", domain)
	add := conn.RemoteAddr()
	return add.String()
}
