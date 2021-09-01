package main

import (
	"fmt"
	"io/ioutil"
	"nodepanels/util"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func ProbeUpgrade(tempNo string) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("{\"toolType\":\"probe-upgrade\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"ERROR\"}")
		}
	}()

	//获取入参
	param, _ := ioutil.ReadFile(filepath.Join(util.Exepath(), "probe-upgrade-"+tempNo+".temp"))

	os.Remove(filepath.Join(util.Exepath(), "probe-upgrade-"+tempNo+".temp"))

	url := strings.Split(string(param), " ")[1]

	util.Download(url, filepath.Join(util.Exepath(), "nodepanels-probe.temp"))

	if runtime.GOOS == "windows" {
		fmt.Println("{\"toolType\":\"probe-upgrade\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"" + util.GetHostId() + "\"}")
		fmt.Println("{\"toolType\":\"probe-upgrade\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")
		exec.Command("cmd", "/C", "net stop Nodepanels-probe").Output()
	}
	if runtime.GOOS == "linux" {
		os.Chmod(util.Exepath()+"/nodepanels-probe.temp", 0777)
		os.Rename(util.Exepath()+"/nodepanels-probe.temp", filepath.Join(util.Exepath(), "/nodepanels-probe"))

		fmt.Println("{\"toolType\":\"probe-upgrade\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"" + util.GetHostId() + "\"}")
		fmt.Println("{\"toolType\":\"probe-upgrade\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")
		exec.Command("sh", "-c", "service nodepanels restart").Output()
	}

}
