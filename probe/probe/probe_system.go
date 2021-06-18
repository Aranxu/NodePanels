package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"nodepanels/util"
)

func GetHostInfo(info ProbeInfo) ProbeInfo {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get host info error : " + fmt.Sprintf("%s", err))
		}
	}()

	infoStat, _ := host.Info()
	uptime, _ := host.Uptime()
	kernelArch, _ := host.KernelArch()
	kernelVersion, _ := host.KernelVersion()
	platform, platformFamily, platformVersion, _ := host.PlatformInformation()

	info.HostInfo.Hostname = infoStat.Hostname
	info.HostInfo.Uptime = uptime
	info.HostInfo.KernelArch = kernelArch
	info.HostInfo.KernelVersion = kernelVersion
	info.HostInfo.Os = infoStat.OS
	info.HostInfo.Platform = platform
	info.HostInfo.PlatformFamily = platformFamily
	info.HostInfo.PlatformVersion = platformVersion

	return info
}
