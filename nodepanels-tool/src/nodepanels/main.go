package main

import (
	"fmt"
	"os"
)

//go:generate goversioninfo -icon=favicon.ico

func main() {

	version := "v1.0.2"

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-performance-net-speedtest-ping":
			SpeedTest("performance-net-speedtest-ping", os.Args[2])
		case "-performance-net-speedtest-all":
			SpeedTest("performance-net-speedtest-all", os.Args[2])
		case "-probe-upgrade":
			ProbeUpgrade(os.Args[2])
		case "-process-list":
			GetProcessesList(os.Args[2])
		case "-process-info":
			GetProcessInfo(os.Args[2])
		case "-process-monitor-switch":
			SetMonitorProcessRule(os.Args[2])
		case "-warning-rule-set":
			SetWarningRule(os.Args[2])
		case "-version":
			fmt.Print(version)
		}
	}

}
