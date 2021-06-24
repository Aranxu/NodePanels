package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"nodepanels/util"
	"runtime"
	"strings"
)

func GetDiskInfo(info ProbeInfo) ProbeInfo {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("get disk info error : " + fmt.Sprintf("%s", err))
		}
	}()

	partitionInfo, _ := disk.Partitions(false)

	for _, val := range partitionInfo {

		usage, _ := disk.Usage(val.Mountpoint)

		diskInfo := DiskInfo{}
		diskInfo.Device = val.Device
		diskInfo.Mountpoint = val.Mountpoint
		diskInfo.Fstype = val.Fstype
		diskInfo.Total = usage.Total
		diskInfo.Used = usage.Used

		info.DiskInfo = append(info.DiskInfo, diskInfo)
	}

	//记录初始读写数据
	ioData := getDiskIOCounters()

	diskInitReadBytes = make(map[string]uint64)
	diskInitWriteBytes = make(map[string]uint64)

	for key, val := range ioData {
		diskInitReadBytes[key] = val.ReadBytes
		diskInitWriteBytes[key] = val.WriteBytes
	}

	return info
}

func GetDiskUsage(probeUsage ProbeUsage) ProbeUsage {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("get disk io usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	ioData := getDiskIOCounters()

	for key, val := range ioData {

		disk := Disk{
			Name:  key,
			Read:  val.ReadBytes - diskInitReadBytes[key],
			Write: val.WriteBytes - diskInitWriteBytes[key],
		}

		diskInitReadBytes[key] = val.ReadBytes
		diskInitWriteBytes[key] = val.WriteBytes

		probeUsage.Disk = append(probeUsage.Disk, disk)
	}

	return probeUsage
}

func getDiskIOCounters() map[string]disk.IOCountersStat {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get disk IO counters error : " + fmt.Sprintf("%s", err))
		}
	}()

	partitionInfo, _ := disk.IOCounters()
	if runtime.GOOS == "linux" {
		disks := make([]string, 0)
		diskMap := make(map[string]string)
		for _, val := range partitionInfo {
			if strings.Index(val.Name, "sd") == 0 || strings.Index(val.Name, "vd") == 0 {
				diskMap[val.Name[0:3]] = "1"
			}
		}
		for key, _ := range diskMap {
			disks = append(disks, key)
		}
		partitionInfo, _ = disk.IOCounters(disks...)
	}
	return partitionInfo
}
