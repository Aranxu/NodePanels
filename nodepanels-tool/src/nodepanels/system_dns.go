package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"nodepanels/util"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func GetDns(tempNo string) {
	file, _ := ioutil.ReadFile("/etc/resolv.conf")

	r, _ := regexp.Compile("(nameserver)\\s+(.+)")

	nameserver, _ := json.Marshal(r.FindAllString(string(file), -1))
	fmt.Println("{\"toolType\":\"system-dns-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":" + string(nameserver) + "}")
	fmt.Println("{\"toolType\":\"system-dns-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-dns-get-"+tempNo+".temp"))
}

func SetDns(tempNo string) {

	var nameserverList []string
	nameservers, _ := ioutil.ReadFile(filepath.Join(util.Exepath(), "system-dns-set-"+tempNo+".temp"))
	json.Unmarshal(nameservers, &nameserverList)

	_, err := os.Stat("/etc/resolv.conf")
	if err == nil {
		file, _ := ioutil.ReadFile("/etc/resolv.conf")
		r, _ := regexp.Compile("(\\nnameserver)\\s+(.+)")
		resolvFileContent := r.ReplaceAll(file, []byte(""))
		for _, v := range nameserverList {
			resolvFileContent = append(resolvFileContent, []byte("\nnameserver "+v)...)
		}
		ioutil.WriteFile("/etc/resolv.conf", resolvFileContent, 0777)
	}

	_, err = os.Stat("/etc/network/interfaces")
	if err == nil {
		file, _ := ioutil.ReadFile("/etc/network/interfaces")
		r, _ := regexp.Compile("(\\nnameserver)\\s+(.+)")
		resolvFileContent := r.ReplaceAll(file, []byte(""))
		for _, v := range nameserverList {
			resolvFileContent = append(resolvFileContent, []byte("\nnameserver "+v)...)
		}
		ioutil.WriteFile("/etc/network/interfaces", resolvFileContent, 0644)
		exec.Command("sh", "-c", "/etc/init.d/networking restart").Output()
	}

	_, err = os.Stat("/etc/NetworkManager/NetworkManager.conf")
	if err == nil {
		file, _ := ioutil.ReadFile("/etc/NetworkManager/NetworkManager.conf")
		if !strings.Contains(string(file), "dns=none") {
			if strings.Contains(string(file), "[main]") {
				file = []byte(strings.Replace(string(file), "[main]", "[main]\ndns=none", 1))
			}
		}
		ioutil.WriteFile("/etc/NetworkManager/NetworkManager.conf", file, 0644)
		exec.Command("sh", "-c", "systemctl restart network").Output()
	}

	_, err = os.Stat("/etc/sysconfig/network-scripts/ifcfg-eth0")
	if err == nil {
		file, _ := ioutil.ReadFile("/etc/sysconfig/network-scripts/ifcfg-eth0")
		if !strings.Contains(string(file), "PEERDNS=no") {
			file = append(file, []byte("\nPEERDNS=no")...)
		}
		ioutil.WriteFile("/etc/sysconfig/network-scripts/ifcfg-eth0", file, 0644)
		exec.Command("sh", "-c", "systemctl restart network").Output()
	}

	_, err = os.Stat("/etc/resolvconf/resolv.conf.d/base")
	if err == nil {
		file, _ := ioutil.ReadFile("/etc/resolvconf/resolv.conf.d/base")
		r, _ := regexp.Compile("(\\nnameserver)\\s+(.+)")
		resolvFileContent := r.ReplaceAll(file, []byte(""))
		for _, v := range nameserverList {
			resolvFileContent = append(resolvFileContent, []byte("\nnameserver "+v)...)
		}
		ioutil.WriteFile("/etc/resolvconf/resolv.conf.d/base", file, 0644)
		exec.Command("sh", "-c", "resolvconf -u").Output()
		exec.Command("sh", "-c", "/etc/init.d/networking restart").Output()
	}

	fmt.Println("{\"toolType\":\"system-dns-set\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-dns-set-"+tempNo+".temp"))
}

func BackupDns(tempNo string) {
	if _, err := os.Stat(filepath.Join(util.Exepath(), "backup", "dns")); os.IsNotExist(err) {
		os.MkdirAll(filepath.Join(util.Exepath(), "backup", "dns"), 0777)
	}

	_, err := os.Stat("/etc/resolv.conf")
	if err == nil {
		destination, _ := os.Create(filepath.Join(util.Exepath(), "backup", "dns", "resolv.conf"))
		source, _ := os.Open("/etc/resolv.conf")
		io.Copy(destination, source)
	}

	_, err = os.Stat("/etc/network/interfaces")
	if err == nil {
		destination, _ := os.Create(filepath.Join(util.Exepath(), "backup", "dns", "interfaces"))
		source, _ := os.Open("/etc/network/interfaces")
		io.Copy(destination, source)
	}

	_, err = os.Stat("/etc/NetworkManager/NetworkManager.conf")
	if err == nil {
		destination, _ := os.Create(filepath.Join(util.Exepath(), "backup", "dns", "NetworkManager.conf"))
		source, _ := os.Open("/etc/NetworkManager/NetworkManager.conf")
		io.Copy(destination, source)
	}

	_, err = os.Stat("/etc/NetworkManager/NetworkManager.conf")
	if err == nil {
		destination, _ := os.Create(filepath.Join(util.Exepath(), "backup", "dns", "NetworkManager.conf"))
		source, _ := os.Open("/etc/NetworkManager/NetworkManager.conf")
		io.Copy(destination, source)
	}

	_, err = os.Stat("/etc/sysconfig/network-scripts/ifcfg-eth0")
	if err == nil {
		destination, _ := os.Create(filepath.Join(util.Exepath(), "backup", "dns", "ifcfg-eth0"))
		source, _ := os.Open("/etc/sysconfig/network-scripts/ifcfg-eth0")
		io.Copy(destination, source)
	}

	_, err = os.Stat("/etc/resolvconf/resolv.conf.d/base")
	if err == nil {
		destination, _ := os.Create(filepath.Join(util.Exepath(), "backup", "dns", "base"))
		source, _ := os.Open("/etc/resolvconf/resolv.conf.d/base")
		io.Copy(destination, source)
	}

	fmt.Println("{\"toolType\":\"system-dns-backup\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-dns-backup-"+tempNo+".temp"))
}

func RestoreDns(tempNo string) {

	_, srcErr := os.Stat(filepath.Join(util.Exepath(), "backup", "dns", "resolv.conf"))
	if srcErr == nil {
		source, _ := os.ReadFile(filepath.Join(util.Exepath(), "backup", "dns", "resolv.conf"))
		os.WriteFile("/etc/resolv.conf",source,0777)
	}

	_, srcErr = os.Stat(filepath.Join(util.Exepath(), "backup", "dns", "interfaces"))
	if srcErr == nil {
		source, _ := os.ReadFile(filepath.Join(util.Exepath(), "backup", "dns", "interfaces"))
		os.WriteFile("/etc/network/interfaces",source,0777)
	}

	_, srcErr = os.Stat(filepath.Join(util.Exepath(), "backup", "dns", "NetworkManager.conf"))
	if srcErr == nil {
		source, _ := os.ReadFile(filepath.Join(util.Exepath(), "backup", "dns", "NetworkManager.conf"))
		os.WriteFile("/etc/NetworkManager/NetworkManager.conf",source,0777)
	}

	_, srcErr = os.Stat(filepath.Join(util.Exepath(), "backup", "dns", "NetworkManager.conf"))
	if srcErr == nil {
		source, _ := os.ReadFile(filepath.Join(util.Exepath(), "backup", "dns", "NetworkManager.conf"))
		os.WriteFile("/etc/NetworkManager/NetworkManager.conf",source,0777)
	}

	_, srcErr = os.Stat(filepath.Join(util.Exepath(), "backup", "dns", "ifcfg-eth0"))
	if srcErr == nil {
		source, _ := os.ReadFile(filepath.Join(util.Exepath(), "backup", "dns", "ifcfg-eth0"))
		os.WriteFile("/etc/sysconfig/network-scripts/ifcfg-eth0",source,0777)
	}

	_, srcErr = os.Stat(filepath.Join(util.Exepath(), "backup", "dns", "base"))
	if srcErr == nil {
		source, _ := os.ReadFile(filepath.Join(util.Exepath(), "backup", "dns", "base"))
		os.WriteFile("/etc/resolvconf/resolv.conf.d/base",source,0777)
	}

	fmt.Println("{\"toolType\":\"system-dns-restore\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-dns-restore-"+tempNo+".temp"))
}
