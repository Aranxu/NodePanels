package probe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nodepanels/config"
	"nodepanels/util"
)

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
				go util.Post(config.AgentUrl+"/warning/cpuWarning", "serverId="+util.GetHostId()+"&msg="+util.Int2string(cpuPercent))
			}
		} else {
			if c.Warning.Rule.Cpu.Count >= c.Warning.Rule.Cpu.Duration {
				go util.Post(config.AgentUrl+"/warning/cpuRecovery", "serverId="+util.GetHostId())
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
				go util.Post(config.AgentUrl+"/warning/memWarning", "serverId="+util.GetHostId()+"&msg="+util.Int2string(memPercent))
			}
		} else {
			if c.Warning.Rule.Mem.Count >= c.Warning.Rule.Mem.Duration {
				go util.Post(config.AgentUrl+"/warning/memRecovery", "serverId="+util.GetHostId())
			}
			c.Warning.Rule.Mem.Count = 0
		}
		data, _ := json.MarshalIndent(c, "", "\t")
		ioutil.WriteFile(util.Exepath()+"/config", data, 0666)
	}
}
