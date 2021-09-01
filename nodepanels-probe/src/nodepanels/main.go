package main

import (
	"encoding/json"
	"fmt"
	"github.com/kardianos/service"
	"io/ioutil"
	"nodepanels/config"
	"nodepanels/probe"
	"nodepanels/util"
	"nodepanels/ws"
	"os"
	"os/exec"
	"runtime"
	"time"
)

//go:generate goversioninfo -icon=favicon.ico

func main() {

	serviceName := ""
	if runtime.GOOS == "windows" {
		serviceName = "Nodepanels-probe"
	}
	if runtime.GOOS == "linux" {
		serviceName = "nodepanels"
	}

	serConfig := &service.Config{
		Name:        serviceName,
		DisplayName: "Nodepanels-probe",
		Description: "Nodepanels探针进程",
	}

	pro := &Program{}
	s, err := service.New(pro, serConfig)
	if err != nil {
		fmt.Println(err, "Create service failed")
	}

	if len(os.Args) > 1 {

		if os.Args[1] == "-install" {
			err = s.Install()
			if err != nil {
				fmt.Println("Install failed", err)
			} else {
				fmt.Println("Install success")
			}
			return
		}

		if os.Args[1] == "-uninstall" {
			err = s.Uninstall()
			if err != nil {
				fmt.Println("Uninstall failed", err)
			} else {
				fmt.Println("Uninstall success")
			}
			return
		}

		if os.Args[1] == "-version" {
			fmt.Println(util.Logo())
			fmt.Println("====================================")
			fmt.Println("App name    : nodepanels-probe")
			fmt.Println("Version     : v1.0.2")
			fmt.Println("Update Time : 20210819")
			fmt.Println("\nMade by     : https://nodepanels.com")
			fmt.Println("====================================")
			return
		}

		if os.Args[1] == "-help" {
			fmt.Println("Available option is :")
			fmt.Println("1) -install    :  Install nodepanels-probe as a service to system.")
			fmt.Println("2) -uninstall  :  Remove nodepanels-probe service from system.")
			fmt.Println("3) -version    :  Check nodepanels-probe version info.")
			fmt.Println("4) -help       :  How to use nodepanels-probe.")
			return
		}
	}

	err = s.Run()
	if err != nil {
		fmt.Println("Run nodepanels-probe failed", err)
	}

}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	fmt.Println("nodepanels-probe start")
	go p.run()
	return nil
}

func (p *Program) run() {
	StartProbe()
}

func (p *Program) Stop(s service.Service) error {
	fmt.Println("nodepanels-probe stop")
	return nil
}

func StartProbe() {
	util.LogFile, _ = os.OpenFile(util.Exepath()+"/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	util.LogDebug(util.Logo())

	runtime.GOMAXPROCS(1)

	if ProbeCheck() {

		sendServerInfo()

		go ws.CreateAgentConn()

		for {
			go sendUsageInfo()
			time.Sleep(60000 * time.Millisecond)
		}

	}

	if runtime.GOOS == "windows" {
		exec.Command("cmd", "/C", "net stop Nodepanels-probe").Output()
	}
	if runtime.GOOS == "linux" {
		exec.Command("sh", "-c", "service nodepanels stop").Output()
	}

}

func ProbeCheck() bool {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Check probe error : " + fmt.Sprintf("%s", err))
		}
	}()

	if util.GetHostId() == "" {
		util.LogError("The program is not completely installed, please reinstall")
		return false
	}
	exist := util.Get(config.AgentUrl + "/server/exist/" + util.GetHostId())

	c := util.GetConfig()
	data, _ := json.MarshalIndent(c, "", "\t")
	ioutil.WriteFile(util.Exepath()+"/config", data, 0666)

	if exist == "1" {
		util.LogDebug("Program started successfully")
		return true
	} else {
		util.LogError("Invalid server ID, please reinstall")
		return false
	}
}

func sendUsageInfo() {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Sending usage data error : " + fmt.Sprintf("%s", err))
		}
	}()

	probeUsage := probe.ProbeUsage{}
	probeUsage = probe.GetCpuUsage(probeUsage)
	probeUsage = probe.GetMemUsage(probeUsage)
	probeUsage = probe.GetSwapUsage(probeUsage)
	probeUsage = probe.GetDiskUsage(probeUsage)
	probeUsage = probe.GetPartitionUsage(probeUsage)
	probeUsage = probe.GetNetUsage(probeUsage)
	probeUsage = probe.GetProcessNum(probeUsage)
	probeUsage = probe.GetProcessUsage(probeUsage)
	probeUsage = probe.GetLoadUsage(probeUsage)
	probeUsage.Unix = time.Now().Unix()

	msg, _ := json.Marshal(probeUsage)

	resultMap := make(map[string]string)
	resultMap["serverId"] = util.GetHostId()
	resultMap["msg"] = string(msg)
	result, _ := json.Marshal(resultMap)
	util.PostJson(config.ApiUrl+"/api/v1", result)
}

func sendServerInfo() {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Sending server info data error : " + fmt.Sprintf("%s", err))
		}
	}()

	probeInfo := probe.ProbeInfo{}
	probeInfo.Version = config.Version
	probeInfo = probe.GetHostInfo(probeInfo)
	probeInfo = probe.GetCpuInfo(probeInfo)
	probeInfo = probe.GetMemInfo(probeInfo)
	probeInfo = probe.GetDiskInfo(probeInfo)
	probeInfo = probe.GetNetInfo(probeInfo)

	msg, _ := json.Marshal(probeInfo)

	resultMap := make(map[string]string)
	resultMap["serverId"] = util.GetHostId()
	resultMap["msg"] = string(msg)
	result, _ := json.Marshal(resultMap)

	go util.PostJson(config.AgentUrl+"/server/info", result)
}
