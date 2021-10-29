package main

import (
	"fmt"
	"nodepanels/util"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetEnv(tempNo string) {
	var env, _ = exec.Command("sh", "-c", "env").Output()

	fmt.Println("{\"toolType\":\"system-env-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"" + strings.ReplaceAll(string(env), "\n", "\\n") + "\"}")
	fmt.Println("{\"toolType\":\"system-env-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-env-get-"+tempNo+".temp"))
}
