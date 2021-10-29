package main

import (
	"fmt"
	"io/ioutil"
	"nodepanels/util"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetTimeInfo(tempNo string) {
	var time, _ = exec.Command("sh", "-c", "date +\"%Y-%m-%d %H:%M:%S\"").Output()
	var timestamp, _ = exec.Command("sh", "-c", "date +%s").Output()
	var timezone, _ = exec.Command("sh", "-c", "ls -il /etc | grep localtime | awk '{print $12}' | awk -F zoneinfo/ '{print $2}'").Output()
	var timezoneNum, _ = exec.Command("sh", "-c", "date +\"%z\"").Output()
	fmt.Println("{\"toolType\":\"system-time-info-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":{\"timezoneNum\":\"" + strings.Replace(string(timezoneNum), "\n", "", -1) + "\",\"timezone\":\"" + strings.Replace(string(timezone), "\n", "", -1) + "\",\"time\":\"" + strings.Replace(string(time), "\n", "", -1) + "\",\"timestamp\":\"" + strings.Replace(string(timestamp), "\n", "", -1) + "\"}}")
	fmt.Println("{\"toolType\":\"system-time-info-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-time-info-get-"+tempNo+".temp"))
}

func SetTimeZone(tempNo string) {
	timezone, _ := ioutil.ReadFile(filepath.Join(util.Exepath(), "system-time-zone-set-"+tempNo+".temp"))

	exec.Command("sh", "-c", "ln -snf /usr/share/zoneinfo/"+string(timezone)+" /etc/localtime").Output()

	fmt.Println("{\"toolType\":\"system-time-zone-set\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-time-zone-set-"+tempNo+".temp"))
}

func SetTime(tempNo string) {
	time, _ := ioutil.ReadFile(filepath.Join(util.Exepath(), "system-time-set-"+tempNo+".temp"))

	exec.Command("sh", "-c", "date -s \""+string(time)+"\"").Output()
	exec.Command("sh", "-c", "hwclock -w").Output()
	exec.Command("sh", "-c", "ln -snf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime").Output()

	fmt.Println("{\"toolType\":\"system-time-set\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-time-set-"+tempNo+".temp"))
}
