package main

import (
	"fmt"
	"os"
)

//go:generate goversioninfo -icon=favicon.ico

func main() {

	version := "v1.0.4"

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
		case "-system-hostname-get":
			GetHostname(os.Args[2])
		case "-system-hostname-set":
			SetHostname(os.Args[2])
		case "-system-dns-get":
			GetDns(os.Args[2])
		case "-system-dns-set":
			SetDns(os.Args[2])
		case "-system-dns-backup":
			BackupDns(os.Args[2])
		case "-system-dns-restore":
			RestoreDns(os.Args[2])
		case "-system-yum-get":
			GetYum(os.Args[2])
		case "-system-yum-set":
			SetYum(os.Args[2])
		case "-system-yum-file-set":
			SetYumFile(os.Args[2])
		case "-system-yum-backup":
			BackupYum(os.Args[2])
		case "-system-yum-restore":
			RestoreYum(os.Args[2])
		case "-system-time-info-get":
			GetTimeInfo(os.Args[2])
		case "-system-time-zone-set":
			SetTimeZone(os.Args[2])
		case "-system-time-set":
			SetTime(os.Args[2])
		case "-system-env-get":
			GetEnv(os.Args[2])
		case "-system-startup-get":
			GetStartup(os.Args[2])
		case "-system-service-get":
			GetService(os.Args[2])
		case "-version":
			fmt.Print(version)
		}
	}

}
