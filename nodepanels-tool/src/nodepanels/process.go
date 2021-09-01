package main

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"io/ioutil"
	"nodepanels/util"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func GetProcessesList(tempNo string) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("{\"toolType\":\"process-list\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"ERROR\"}")
		}
	}()

	var list []ProcessUsage

	processes, _ := process.Processes()
	for _, val := range processes {
		name, _ := val.Name()
		cmd, _ := val.Cmdline()
		pid := val.Pid
		cpuPercent, _ := val.CPUPercent()
		memPercent, _ := val.MemoryPercent()

		processUsage := ProcessUsage{}
		processUsage.Name = name
		processUsage.Cmd = cmd
		if strings.Index(util.Float642string(cpuPercent), "0.0") == 0 {
			processUsage.CpuPercent = float64(0)
		} else {
			processUsage.CpuPercent = cpuPercent
		}
		processUsage.MemPercent = memPercent
		processUsage.Pid = pid

		list = append(list, processUsage)
	}
	sort.Sort(ProcessUsageSlice(list))
	if len(list) >= 30 {
		list = list[0:30]
	}

	result, _ := json.Marshal(list)

	fmt.Println("{\"toolType\":\"process-list\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":" + string(result) + "}")
	fmt.Println("{\"toolType\":\"process-list\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "process-list-"+tempNo+".temp"))
}

func GetProcessInfo(tempNo string) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get process info error : " + fmt.Sprintf("%s", err))
		}
	}()

	//获取入参
	processPid, _ := ioutil.ReadFile(filepath.Join(util.Exepath(), "process-info-"+tempNo+".temp"))

	process, _ := process.NewProcess(util.String2int32(string(processPid)))

	if process != nil {

		processInfo := ProcessInfo{}
		cmd, _ := process.Cmdline()
		name, _ := process.Name()
		cwd, _ := process.Cwd()
		exe, _ := process.Exe()
		createTime, _ := process.CreateTime()
		foreground, _ := process.Foreground()
		nice, _ := process.Nice()
		numCtxSwitches, _ := process.NumCtxSwitches()
		numThreads, _ := process.NumThreads()
		openFiles, _ := process.OpenFiles()
		status, _ := process.Status()
		username, _ := process.Username()

		processInfo.Cmd = cmd
		processInfo.Name = name
		processInfo.Cwd = cwd
		processInfo.Exe = exe
		processInfo.CreateTime = createTime
		processInfo.Foreground = foreground
		processInfo.Nice = nice
		if numCtxSwitches != nil {
			processInfo.NumCtxSwitchesVoluntary = numCtxSwitches.Voluntary
			processInfo.NumCtxSwitchesInvoluntary = numCtxSwitches.Involuntary
		}
		processInfo.NumThreads = numThreads
		processInfo.OpenFiles = len(openFiles)
		processInfo.Status = status[0]
		processInfo.Username = username

		msg, _ := json.Marshal(processInfo)

		fmt.Println("{\"toolType\":\"process-info\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":" + string(msg) + "}")
		fmt.Println("{\"toolType\":\"process-info\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

		os.Remove(filepath.Join(util.Exepath(), "process-info-"+tempNo+".temp"))
	}
}

type ProcessUsage struct {
	Name       string  `json:"name"`
	CpuPercent float64 `json:"cpu"`
	MemPercent float32 `json:"mem"`
	DiskWrite  uint64  `json:"write"`
	DiskRead   uint64  `json:"read"`
	Cmd        string  `json:"cmd"`
	Pid        int32   `json:"pid"`
}

type ProcessUsageSlice []ProcessUsage

func (p ProcessUsageSlice) Len() int {
	return len(p)
}

func (p ProcessUsageSlice) Less(i, j int) bool {
	if p[i].CpuPercent == p[j].CpuPercent {
		return p[i].MemPercent > p[j].MemPercent
	} else {
		return p[i].CpuPercent > p[j].CpuPercent
	}
}

func (p ProcessUsageSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ProcessInfo struct {
	Name                      string `json:"name"`
	Cmd                       string `json:"cmd"`
	Cwd                       string `json:"cwd"`
	Exe                       string `json:"exe"`
	CreateTime                int64  `json:"createTime"`
	Foreground                bool   `json:"foreground"`
	Nice                      int32  `json:"nice"`
	NumCtxSwitchesVoluntary   int64  `json:"voluntary"`
	NumCtxSwitchesInvoluntary int64  `json:"involuntary"`
	NumThreads                int32  `json:"numThreads"`
	OpenFiles                 int    `json:"openFiles"`
	Status                    string `json:"status"`
	Username                  string `json:"username"`
}
