package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	for{
		if runtime.GOOS == "linux" {
			cmd := exec.Command("sh", "-c", "ps -ef |grep 'nodepanels-probe' |grep -v grep  | awk '{print $2}'")
			output, _ := cmd.Output()
			if string(output) == ""{
				fmt.Println("linux")
				exec.Command("sh", "-c", "service nodepanels restart")
			}
		}
		if runtime.GOOS == "windows" {
			cmd := exec.Command("cmd", "/C", "tasklist|findstr nodepanels-probe")
			output, _ := cmd.Output()
			fmt.Println(string(output))
			if string(output) == ""{
				fmt.Println("windows")
			}
		}
		time.Sleep(60000 * time.Millisecond)
	}
}
