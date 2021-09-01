package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nodepanels/util"
	"os"
	"path/filepath"
)

func SetMonitorProcessRule(tempNo string) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Set monitor rule error : " + fmt.Sprintf("%s", err))
		}
	}()

	//获取入参
	processCmdListJsonStr, _ := ioutil.ReadFile(filepath.Join(util.Exepath(), "process-monitor-switch-"+tempNo+".temp"))

	var processCmdList []string
	json.Unmarshal(processCmdListJsonStr, &processCmdList)

	c := util.GetConfig()
	c.Monitor.Rule.Process = processCmdList

	data, _ := json.MarshalIndent(c, "", "\t")
	ioutil.WriteFile(util.Exepath()+"/config", data, 0666)

	fmt.Println("{\"toolType\":\"process-monitor-switch\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "process-monitor-switch-"+tempNo+".temp"))
}
