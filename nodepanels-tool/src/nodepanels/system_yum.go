package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"io"
	"io/ioutil"
	"nodepanels/util"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func GetYum(tempNo string) {
	var url string
	var path string
	var file []byte

	platform, _, _, _ := host.PlatformInformation()

	if platform == "centos" {
		path = "/etc/yum.repos.d/CentOS-Base.repo"
		file, _ = ioutil.ReadFile("/etc/yum.repos.d/CentOS-Base.repo")
		r, _ := regexp.Compile("(?m).*(baseurl=http).*")
		url = r.FindAllString(string(file), 1)[0]
		url = strings.Split(url, "baseurl=")[1]
		url = strings.Split(url, "/centos")[0]
	} else if platform == "ubuntu" {
		path = "/etc/apt/sources.list"
		file, _ = ioutil.ReadFile("/etc/apt/sources.list")
		r, _ := regexp.Compile("(?m)^(deb http).*")
		url = r.FindAllString(string(file), 1)[0]
		url = strings.Split(url, "deb ")[1]
		url = strings.Split(url, "/ubuntu")[0]
	} else if platform == "debian" {
		path = "/etc/apt/sources.list"
		file, _ = ioutil.ReadFile("/etc/apt/sources.list")
		r, _ := regexp.Compile("(?m)^(deb http).*")
		url = r.FindAllString(string(file), 1)[0]
		url = strings.Split(url, "deb ")[1]
		url = strings.Split(url, "/debian")[0]
	}
	content := strings.ReplaceAll(strings.ReplaceAll(string(file), "\n", "\\n"), "\"", "\\\"")
	result := "{\"path\":\"" + path + "\",\"file\":\"" + content + "\",\"url\":\"" + url + "\"}"
	fmt.Println("{\"toolType\":\"system-yum-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":" + result + "}")
	fmt.Println("{\"toolType\":\"system-yum-get\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-yum-get-"+tempNo+".temp"))
}

func SetYum(tempNo string) {
	yum, _ := ioutil.ReadFile(filepath.Join(util.Exepath(), "system-yum-set-"+tempNo+".temp"))

	platform, _, platformVersion, _ := host.PlatformInformation()

	if platform == "centos" {
		platformVersion = strings.Split(platformVersion, ".")[0]
		util.Download("https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/yum/"+string(yum)+"/centos/"+platformVersion+"/CentOS-Base.repo", "/etc/yum.repos.d/CentOS-Base.repo")
		os.Chmod("/etc/yum.repos.d/CentOS-Base.repo", 0644)
		fmt.Println("{\"toolType\":\"system-yum-set\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"yum clean\"}")
		exec.Command("sh", "-c", "yum clean all").Output()
		fmt.Println("{\"toolType\":\"system-yum-set\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"yum makecache\"}")
		exec.Command("sh", "-c", "yum makecache").Output()
	} else if platform == "ubuntu" {
		util.Download("https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/yum/"+string(yum)+"/ubuntu/"+platformVersion+"/sources.list", "/etc/apt/sources.list")
		os.Chmod("/etc/apt/sources.list", 0644)
	} else if platform == "debian" {
		platformVersion = strings.Split(platformVersion, ".")[0]
		util.Download("https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/yum/"+string(yum)+"/debian/"+platformVersion+"/sources.list", "/etc/apt/sources.list")
		os.Chmod("/etc/apt/sources.list", 0644)
	}

	fmt.Println("{\"toolType\":\"system-yum-set\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-yum-set-"+tempNo+".temp"))
}

func SetYumFile(tempNo string) {
	content, _ := ioutil.ReadFile(filepath.Join(util.Exepath(), "system-yum-file-set-"+tempNo+".temp"))

	platform, _, _, _ := host.PlatformInformation()

	if platform == "centos" {
		os.WriteFile("/etc/yum.repos.d/CentOS-Base.repo", content, 0644)
		fmt.Println("{\"toolType\":\"system-yum-set\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"yum clean\"}")
		exec.Command("sh", "-c", "yum clean all").Output()
		fmt.Println("{\"toolType\":\"system-yum-set\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"yum makecache\"}")
		exec.Command("sh", "-c", "yum makecache").Output()
	} else if platform == "ubuntu" {
		os.WriteFile("/etc/apt/sources.list", content, 0644)
	} else if platform == "debian" {
		os.WriteFile("/etc/apt/sources.list", content, 0644)
	}

	fmt.Println("{\"toolType\":\"system-yum-file-set\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-yum-file-set-"+tempNo+".temp"))
}

func BackupYum(tempNo string) {
	var srcPath string
	var dstPath string

	platform, _, _, _ := host.PlatformInformation()

	if platform == "centos" {
		srcPath = "/etc/yum.repos.d/CentOS-Base.repo"
		dstPath = filepath.Join(util.Exepath(), "backup", "yum", "CentOS-Base.repo")
	} else if platform == "ubuntu" {
		srcPath = "/etc/apt/sources.list"
		dstPath = filepath.Join(util.Exepath(), "backup", "yum", "sources.list")
	} else if platform == "debian" {
		srcPath = "/etc/apt/sources.list"
		dstPath = filepath.Join(util.Exepath(), "backup", "yum", "sources.list")
	}

	source, _ := os.Open(srcPath)
	defer source.Close()

	if _, err := os.Stat(filepath.Join(util.Exepath(), "backup", "yum")); os.IsNotExist(err) {
		os.MkdirAll(filepath.Join(util.Exepath(), "backup", "yum"), 0777)
	}
	destination, _ := os.Create(dstPath)
	defer destination.Close()

	io.Copy(destination, source)

	fmt.Println("{\"toolType\":\"system-yum-backup\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), "system-yum-backup-"+tempNo+".temp"))
}

func RestoreYum(tempNo string) {
	var srcPath string
	var dstPath string
	var filename string

	platform, _, _, _ := host.PlatformInformation()

	if platform == "centos" {
		srcPath = filepath.Join(util.Exepath(), "backup", "yum", "CentOS-Base.repo")
		dstPath = "/etc/yum.repos.d/CentOS-Base.repo"
		filename = "CentOS-Base.repo"
	} else if platform == "ubuntu" {
		srcPath = filepath.Join(util.Exepath(), "backup", "yum", "sources.list")
		dstPath = "/etc/apt/sources.list"
		filename = "sources.list"
	} else if platform == "debian" {
		srcPath = filepath.Join(util.Exepath(), "backup", "yum", "sources.list")
		dstPath = "/etc/apt/sources.list"
		filename = "sources.list"
	}

	if _, err := os.Stat(filepath.Join(util.Exepath(), "backup", "yum", filename)); os.IsNotExist(err) {
		fmt.Println("{\"toolType\":\"system-yum-restore\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"ERROR\"}")
	} else {
		source, _ := os.Open(srcPath)
		defer source.Close()

		destination, _ := os.Create(dstPath)
		defer destination.Close()

		io.Copy(destination, source)

		if platform == "centos" {
			fmt.Println("{\"toolType\":\"system-yum-restore\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"yum clean\"}")
			exec.Command("sh", "-c", "yum clean all").Output()
			fmt.Println("{\"toolType\":\"system-yum-restore\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"yum makecache\"}")
			exec.Command("sh", "-c", "yum makecache").Output()
		}

		fmt.Println("{\"toolType\":\"system-yum-restore\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")
	}
	os.Remove(filepath.Join(util.Exepath(), "system-yum-restore-"+tempNo+".temp"))
}
