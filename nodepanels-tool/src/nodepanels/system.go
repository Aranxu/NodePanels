package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"io/ioutil"
	"nodepanels/util"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func GetHostname(tempNo string) {
	infoStat, _ := host.Info()

	fmt.Println("{\"toolType\":\"system-hostname-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":" + infoStat.Hostname + "}")
	fmt.Println("{\"toolType\":\"system-hostname-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-hostname-get-"+tempNo+".temp"))
}

func SetHostname(tempNo string) {
	hostname, _ := ioutil.ReadFile(filepath.Join(util.Exepath(), "system-hostname-set-"+tempNo+".temp"))

	if runtime.GOOS == "linux" {
		exec.Command("sh", "-c", "hostnamectl set-hostname "+string(hostname)).Output()
	} else if runtime.GOOS == "windows" {
		exec.Command("cmd", "/C", "WMIC computersystem where caption=\"%computername%\" rename \""+string(hostname)+"\"").Output()
	}
	fmt.Println("{\"toolType\":\"system-hostname-set\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-hostname-set-"+tempNo+".temp"))
}