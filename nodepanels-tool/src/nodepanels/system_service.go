package main

import (
	"fmt"
	"nodepanels/util"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetService(tempNo string) {

	var startup, _ = exec.Command("sh", "-c", "systemctl | grep '.service' | awk '{print $1,$3,$4,$5}'").Output()
	if strings.Index(string(startup), "command not found") < 0 {
		fmt.Println("{\"toolType\":\"system-service-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"" + strings.ReplaceAll(strings.ReplaceAll(string(startup), "\\", "\\\\"), "\n", "\\n") + "\"}")
	}
	fmt.Println("{\"toolType\":\"system-service-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-service-get-"+tempNo+".temp"))
}
