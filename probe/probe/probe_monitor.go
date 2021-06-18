package probe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nodepanels/util"
)

func SetMonitorProcessRule(msg string) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Set monitor rule error : " + fmt.Sprintf("%s", err))
		}
	}()

	var processCmdList []string
	json.Unmarshal([]byte(msg), &processCmdList)

	c := util.GetConfig()
	c.Monitor.Rule.Process = processCmdList

	data, _ := json.MarshalIndent(c, "", "\t")
	ioutil.WriteFile(util.Exepath()+"/config", data, 0666)

}
