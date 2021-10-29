package main

import (
	"fmt"
	"nodepanels/util"
	"os"
	"path/filepath"
)

func GetStartup(tempNo string) {

	fmt.Println("{\"toolType\":\"system-startup-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-startup-get-"+tempNo+".temp"))
}
