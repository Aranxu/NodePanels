package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"nodepanels/util"
)

func GetMemInfo(info ProbeInfo) ProbeInfo {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get mem info error : " + fmt.Sprintf("%s", err))
		}
	}()

	virtualMemoryStat, _ := mem.VirtualMemory()
	swapMemory, _ := mem.SwapMemory()

	info.MemInfo.Mem = virtualMemoryStat.Total
	info.MemInfo.Swap = swapMemory.Total

	//全局变量，给进程计算实际使用内存
	memTotal = uint(virtualMemoryStat.Total / 1024 / 1024)

	return info
}

func GetMemUsage(probeUsage ProbeUsage) ProbeUsage {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get mem usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	virtualMemoryStat, _ := mem.VirtualMemory()

	JudgeMemWarning(util.String2int(util.Float642string(util.RoundFloat64(virtualMemoryStat.UsedPercent, 0))))

	probeUsage.Mem.Usage = util.RoundFloat64(virtualMemoryStat.UsedPercent, 2)

	return probeUsage
}

func GetSwapUsage(probeUsage ProbeUsage) ProbeUsage {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get swap usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	swapMemory, _ := mem.SwapMemory()

	probeUsage.Swap.Usage = util.RoundFloat64(swapMemory.UsedPercent, 2)

	return probeUsage
}
