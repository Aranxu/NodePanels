package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nodepanels/util"
	"os"
	"path/filepath"
	"strings"
)

func SetWarningRule(tempNo string) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Set warning rule error : " + fmt.Sprintf("%s", err))
		}
	}()

	//获取入参
	msg, _ := ioutil.ReadFile(filepath.Join(util.Exepath(), "warning-rule-set-"+tempNo+".temp"))

	totalSwitch := strings.Split(string(msg), "|")[0]

	cpuSwitch := strings.Split(strings.Split(string(msg), "|")[1], ";")[0]
	cpuValue := strings.Split(strings.Split(string(msg), "|")[1], ";")[1]
	cpuDuration := strings.Split(strings.Split(string(msg), "|")[1], ";")[2]

	c := util.GetConfig()

	c.Warning.Switch = util.String2int(totalSwitch)
	c.Warning.Rule.Cpu.Switch = util.String2int(cpuSwitch)
	c.Warning.Rule.Cpu.Value = util.String2int(cpuValue)
	c.Warning.Rule.Cpu.Duration = util.String2int(cpuDuration)
	c.Warning.Rule.Cpu.Count = 0

	memSwitch := strings.Split(strings.Split(string(msg), "|")[2], ";")[0]
	memValue := strings.Split(strings.Split(string(msg), "|")[2], ";")[1]
	memDuration := strings.Split(strings.Split(string(msg), "|")[2], ";")[2]

	c.Warning.Rule.Mem.Switch = util.String2int(memSwitch)
	c.Warning.Rule.Mem.Value = util.String2int(memValue)
	c.Warning.Rule.Mem.Duration = util.String2int(memDuration)
	c.Warning.Rule.Mem.Count = 0

	data, _ := json.MarshalIndent(c, "", "\t")
	ioutil.WriteFile(util.Exepath()+"/config", data, 0666)

	fmt.Println("{\"toolType\":\"warning-rule-set\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"" + util.GetHostId() + "\"}")
	fmt.Println("{\"toolType\":\"warning-rule-set\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "warning-rule-set-"+tempNo+".temp"))
}
