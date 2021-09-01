package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"nodepanels/util"
)

func GetPartitionUsage(probeUsage ProbeUsage) ProbeUsage {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("get partition usage info error : " + fmt.Sprintf("%s", err))
		}
	}()

	partitionInfo, _ := disk.Partitions(false)

	for _, val := range partitionInfo {

		usage, _ := disk.Usage(val.Mountpoint)

		partitionInfo := Partition{}
		partitionInfo.Device = val.Device
		partitionInfo.Used = usage.Used

		probeUsage.Partition = append(probeUsage.Partition, partitionInfo)
	}

	return probeUsage
}
