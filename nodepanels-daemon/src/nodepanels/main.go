package main

import (
	"fmt"
	"github.com/kardianos/service"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

//go:generate goversioninfo -icon=favicon.ico

func main() {

	serConfig := &service.Config{
		Name:        "Nodepanels-daemon",
		DisplayName: "Nodepanels-daemon",
		Description: "Nodepanels探针守护进程",
	}

	pro := &Program{}
	s, err := service.New(pro, serConfig)
	if err != nil {
		fmt.Println(err, "Create service failed")
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err = s.Install()
			if err != nil {
				fmt.Println("Install failed", err)
			} else {
				fmt.Println("Install success")
			}
			return
		}

		if os.Args[1] == "uninstall" {
			err = s.Uninstall()
			if err != nil {
				fmt.Println("Uninstall err", err)
			} else {
				fmt.Println("Uninstall success")
			}
			return
		}
	}

	err = s.Run()
	if err != nil {
		fmt.Println("Run nodepanels-daemon failed", err)
	}

}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	fmt.Println("nodepanels-daemon start")
	go p.run()
	return nil
}

func (p *Program) run() {
	for {
		if runtime.GOOS == "linux" {
			cmd := exec.Command("sh", "-c", "ps -ef |grep 'nodepanels-probe' |grep -v grep  | awk '{print $2}'")
			output, _ := cmd.Output()
			if string(output) == "" {
				exec.Command("sh", "-c", "service nodepanels restart").Output()
			}
		}
		if runtime.GOOS == "windows" {
			cmd := exec.Command("cmd", "/C", "tasklist|findstr nodepanels-probe.exe")
			output, _ := cmd.Output()
			if string(output) == "" {

				_, err := os.Stat(Exepath() + "/nodepanels-probe.temp")
				if err == nil {
					os.Rename(Exepath()+"/nodepanels-probe.temp", Exepath()+"/nodepanels-probe.exe")
				}

				exec.Command("cmd", "/C", "net", "start", "Nodepanels-probe").Output()
			}
		}
		time.Sleep(60000 * time.Millisecond)
	}
}

func (p *Program) Stop(s service.Service) error {
	fmt.Println("nodepanels-daemon stop")
	return nil
}

func Exepath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return ""
	}
	return string(path[0 : i+1])
}
