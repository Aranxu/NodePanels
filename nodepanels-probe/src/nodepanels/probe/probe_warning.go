package probe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nodepanels/config"
	"nodepanels/util"
	"strings"
)

func SetWarningRule(msg string) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Set warning rule error : " + fmt.Sprintf("%s", err))
		}
	}()

	totalSwitch := strings.Split(msg, "|")[0]

	cpuSwitch := strings.Split(strings.Split(msg, "|")[1], ";")[0]
	cpuValue := strings.Split(strings.Split(msg, "|")[1], ";")[1]
	cpuDuration := strings.Split(strings.Split(msg, "|")[1], ";")[2]

	c := util.GetConfig()

	c.Warning.Switch = util.String2int(totalSwitch)
	c.Warning.Rule.Cpu.Switch = util.String2int(cpuSwitch)
	c.Warning.Rule.Cpu.Value = util.String2int(cpuValue)
	c.Warning.Rule.Cpu.Duration = util.String2int(cpuDuration)
	c.Warning.Rule.Cpu.Count = 0

	memSwitch := strings.Split(strings.Split(msg, "|")[2], ";")[0]
	memValue := strings.Split(strings.Split(msg, "|")[2], ";")[1]
	memDuration := strings.Split(strings.Split(msg, "|")[2], ";")[2]

	c.Warning.Rule.Mem.Switch = util.String2int(memSwitch)
	c.Warning.Rule.Mem.Value = util.String2int(memValue)
	c.Warning.Rule.Mem.Duration = util.String2int(memDuration)
	c.Warning.Rule.Mem.Count = 0

	data, _ := json.MarshalIndent(c, "", "\t")
	ioutil.WriteFile(util.Exepath()+"/config", data, 0666)
}

func JudgeCpuWarning(cpuPercent int) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Judge CPU warning error : " + fmt.Sprintf("%s", err))
		}
	}()

	c := util.GetConfig()
	if c.Warning.Switch == 1 && c.Warning.Rule.Cpu.Switch == 1 {
		if cpuPercent >= c.Warning.Rule.Cpu.Value {
			c.Warning.Rule.Cpu.Count++
			if c.Warning.Rule.Cpu.Count == c.Warning.Rule.Cpu.Duration {
				go util.Post("https://"+config.AgentUrl+"/warning/cpuWarning", "serverId="+util.GetHostId()+"&msg="+util.Int2string(cpuPercent))
			}
		} else {
			if c.Warning.Rule.Cpu.Count >= c.Warning.Rule.Cpu.Duration {
				go util.Post("https://"+config.AgentUrl+"/warning/cpuRecovery", "serverId="+util.GetHostId())
			}
			c.Warning.Rule.Cpu.Count = 0
		}
		data, _ := json.MarshalIndent(c, "", "\t")
		ioutil.WriteFile(util.Exepath()+"/config", data, 0666)
	}
}

func JudgeMemWarning(memPercent int) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Judge MEM error : " + fmt.Sprintf("%s", err))
		}
	}()

	c := util.GetConfig()
	if c.Warning.Switch == 1 && c.Warning.Rule.Mem.Switch == 1 {
		if memPercent >= c.Warning.Rule.Mem.Value {
			c.Warning.Rule.Mem.Count++
			if c.Warning.Rule.Mem.Count == c.Warning.Rule.Mem.Duration {
				go util.Post("https://"+config.AgentUrl+"/warning/memWarning", "serverId="+util.GetHostId()+"&msg="+util.Int2string(memPercent))
			}
		} else {
			if c.Warning.Rule.Mem.Count >= c.Warning.Rule.Mem.Duration {
				go util.Post("https://"+config.AgentUrl+"/warning/memRecovery", "serverId="+util.GetHostId())
			}
			c.Warning.Rule.Mem.Count = 0
		}
		data, _ := json.MarshalIndent(c, "", "\t")
		ioutil.WriteFile(util.Exepath()+"/config", data, 0666)
	}
}
