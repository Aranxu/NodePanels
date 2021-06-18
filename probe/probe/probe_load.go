package probe

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/load"
	"nodepanels/util"
)

func GetLoadUsage(probeUsage ProbeUsage) ProbeUsage {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Get load usage error : " + fmt.Sprintf("%s", err))
		}
	}()

	avg, _ := load.Avg()

	probeUsage.Load.SysLoad = avg.Load1

	return probeUsage

}
