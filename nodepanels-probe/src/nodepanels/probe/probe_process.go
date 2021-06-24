package probe

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/process"
	"nodepanels/util"
	"sort"
	"strings"
)

func GetProcessUsage(probeUsage ProbeUsage) ProbeUsage {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get process usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	if initProcessUsageMap == nil {
		initProcessUsageMap = make(map[string]ProcessUsage)
	}

	processCmdList := util.GetConfig().Monitor.Rule.Process
	processCmdMap := make(map[string]int)
	for _, val := range processCmdList {
		processCmdMap[val] = 1
	}

	//临时map
	initProcessUsageTempMap := make(map[string]ProcessUsage)

	processUsageList := []ProcessUsage{}

	processes, _ := process.Processes()
	for _, val := range processes {

		cmd, _ := val.Cmdline()

		//判断是否需要监控
		if _, ok := processCmdMap[cmd]; ok {

			pid := val.Pid
			name, _ := val.Name()
			cpuPercent, _ := val.CPUPercent()
			memPercent, _ := val.MemoryPercent()
			ioData, _ := val.IOCounters()

			processUsage := ProcessUsage{}
			processUsage.Pid = pid
			processUsage.Name = name
			processUsage.Cmd = cmd
			processUsage.CpuPercent = util.RoundFloat64(cpuPercent, 1)
			processUsage.MemPercent = util.RoundFloat32(memPercent, 1)

			if _, ok := initProcessUsageMap[cmd]; ok {
				if ioData != nil {
					processUsage.DiskRead = ioData.ReadBytes - initProcessUsageMap[cmd].DiskRead
					processUsage.DiskWrite = ioData.WriteBytes - initProcessUsageMap[cmd].DiskWrite
				}
				processUsageList = append(processUsageList, processUsage)
			} else {
				if ioData != nil {
					processUsage.DiskRead = ioData.ReadBytes
					processUsage.DiskWrite = ioData.WriteBytes
				}
			}

			if ioData != nil {
				initProcessUsageTempMap[cmd] = ProcessUsage{
					DiskRead:  ioData.ReadBytes,
					DiskWrite: ioData.WriteBytes,
				}
			}
		}
	}

	probeUsage.Process.ProcessList = processUsageList

	initProcessUsageMap = initProcessUsageTempMap

	return probeUsage
}

func GetProcessNum(probeUsage ProbeUsage) ProbeUsage {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get process num error : " + fmt.Sprintf("%s", err))
		}
	}()

	infoStat, _ := host.Info()

	probeUsage.Process.Num = infoStat.Procs

	return probeUsage
}

func GetProcessesList() string {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get process list error : " + fmt.Sprintf("%s", err))
		}
	}()

	resultStr := ""

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
	for _, val := range list {
		resultStr += val.Name + ";" + util.Float642string(util.RoundFloat64(float64(val.CpuPercent), 1)) + ";" + util.Float642string(util.RoundFloat64(float64(val.MemPercent), 1)) + ";" + util.Int322string(val.Pid) + ";cmd:" + val.Cmd + "|"
	}

	return resultStr
}

func GetProcessCmdByPid(pid int32) string {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get process cmd by pid error : " + fmt.Sprintf("%s", err))
		}
	}()

	process, _ := process.NewProcess(pid)
	if process != nil {
		cmdline, _ := process.Cmdline()
		name, _ := process.Name()
		return cmdline + "$" + name
	}
	return ""
}

func GetProcessInfo(pid int32) string {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get process info error : " + fmt.Sprintf("%s", err))
		}
	}()

	process, _ := process.NewProcess(pid)
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
		processInfo.NumCtxSwitchesVoluntary = numCtxSwitches.Voluntary
		processInfo.NumCtxSwitchesInvoluntary = numCtxSwitches.Involuntary
		processInfo.NumThreads = numThreads
		processInfo.OpenFiles = len(openFiles)
		processInfo.Status = status[0]
		processInfo.Username = username

		msg, _ := json.Marshal(processInfo)

		return string(msg)
	}
	return ""
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
