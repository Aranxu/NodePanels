package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"nodepanels/util"
	"runtime"
	"strings"
)

func GetCpuInfo(info ProbeInfo) ProbeInfo {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("get cpu info error : " + fmt.Sprintf("%s", err))
		}
	}()

	infoStat, _ := cpu.Info()

	cpuNums := 0
	physicalIds := ""
	if runtime.GOOS == "linux" {
		for _, val := range infoStat {
			if !strings.Contains(physicalIds, val.PhysicalID+",") {
				physicalIds += val.PhysicalID + ","
				cpuNums++
			}
		}
	}
	if runtime.GOOS == "windows" {
		cpuNums = len(infoStat)
	}

	physicalCores, _ := cpu.Counts(false)

	logicCore, _ := cpu.Counts(true)

	info.CpuInfo.CpuNums = cpuNums
	info.CpuInfo.PhysicalCores = physicalCores
	info.CpuInfo.LogicCore = logicCore

	info.CpuInfo.Model = infoStat[0].ModelName
	info.CpuInfo.VendorID = infoStat[0].VendorID
	info.CpuInfo.Mhz = infoStat[0].Mhz
	info.CpuInfo.Cache = infoStat[0].CacheSize

	return info
}

func GetCpuUsage(probeUsage ProbeUsage) ProbeUsage {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("get cpu usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	combineCpuUsage, _ := cpu.Percent(0, false)

	for _, ccu := range combineCpuUsage {
		probeUsage.Cpu.Total = util.RoundFloat64(ccu, 2)
		JudgeCpuWarning(util.String2int(fmt.Sprintf("%.0f", ccu)))
	}

	logicCore, _ := cpu.Counts(true)
	if logicCore != 1 {
		perCpuUsage, _ := cpu.Percent(0, true)
		perCpuList := []float64{}
		for _, pcu := range perCpuUsage {
			perCpuList = append(perCpuList, util.RoundFloat64(pcu, 1))
		}
		probeUsage.Cpu.Per = perCpuList
	}

	return probeUsage
}
