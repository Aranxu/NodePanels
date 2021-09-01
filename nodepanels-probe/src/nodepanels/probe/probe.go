package probe

import (
	"fmt"
	"nodepanels/util"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func Upgrade(url string) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Update probe error : " + fmt.Sprintf("%s", err))
		}
	}()

	util.LogInfo("[COMMAND] Upgrade probe - back")

	util.Download(url, filepath.Join(util.Exepath(), "nodepanels-probe.temp"))

	if runtime.GOOS == "windows" {

		exec.Command("cmd", "/C", "net stop Nodepanels-probe").Output()
	}
	if runtime.GOOS == "linux" {
		os.Chmod(util.Exepath()+"/nodepanels-probe.temp", 0777)

		os.Rename(util.Exepath()+"/nodepanels-probe.temp", util.Exepath()+"/nodepanels-probe")

		exec.Command("sh", "-c", "service nodepanels restart").Output()
	}

}

func ShutDown() {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("ShutDown probe error : " + fmt.Sprintf("%s", err))
		}
	}()

	util.LogInfo("[COMMAND] Shutdown probe - back")

	if runtime.GOOS == "windows" {

		exec.Command("cmd", "/C", "net stop Nodepanels-probe").Output()
	}
	if runtime.GOOS == "linux" {

		exec.Command("sh", "-c", "service nodepanels stop").Output()

	}

}

func ExeCmd(cmd string) string {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Execute command error : " + fmt.Sprintf("%s", err))
		}
	}()

	if runtime.GOOS == "windows" {
		output, _ := exec.Command("cmd", "/C", cmd).Output()
		return string(output)
	}
	if runtime.GOOS == "linux" {
		output, _ := exec.Command("sh", "-c", cmd).Output()
		return string(output)
	}

	return ""
}

type ProbeUsage struct {
	Cpu       Cpu         `json:"cpu"`
	Mem       Mem         `json:"mem"`
	Swap      Swap        `json:"swap"`
	Disk      []Disk      `json:"disk"`
	Partition []Partition `json:"partition"`
	Net       Net         `json:"net"`
	Process   Process     `json:"process"`
	Load      Load        `json:"load"`
	Unix      int64       `json:"unix"`
}

type Cpu struct {
	Total float64   `json:"total"`
	Per   []float64 `json:"per"`
}

type Mem struct {
	Usage float64 `json:"usage"`
}

type Swap struct {
	Usage float64 `json:"usage"`
}

type Disk struct {
	Name  string `json:name`
	Read  uint64 `json:read`
	Write uint64 `json:"write"`
}

type Partition struct {
	Device string `json:"device""`
	Used   uint64 `json:"used"`
}

type Net struct {
	Rx uint64 `json:"rx"`
	Tx uint64 `json:"tx"`
}

type Process struct {
	Num         uint64         `json:"num"`
	ProcessList []ProcessUsage `json:"list"`
}

type Load struct {
	SysLoad float64 `json:"sysLoad"`
}

type ProbeInfo struct {
	Version  string     `json:"version"`
	HostInfo HostInfo   `json:"host"`
	CpuInfo  CpuInfo    `json:"cpu"`
	MemInfo  MemInfo    `json:"mem"`
	DiskInfo []DiskInfo `json:"disk"`
	NetInfo  NetInfo    `json:"net"`
}

type HostInfo struct {
	Hostname        string `json:"hostname"`
	Uptime          uint64 `json:"uptime"`
	KernelArch      string `json:"kernelArch"`
	KernelVersion   string `json:"kernelVersion"`
	Os              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformFamily  string `json:"platformFamily"`
	PlatformVersion string `json:"platformVersion"`
}

type CpuInfo struct {
	CpuNums       int     `json:"num"`
	PhysicalCores int     `json:"physical"`
	LogicCore     int     `json:"logic"`
	Model         string  `json:"model"`
	VendorID      string  `json:"vendor"`
	Mhz           float64 `json:"mhz"`
	Cache         int32   `json:"cache"`
}

type MemInfo struct {
	Mem  uint64 `json:"mem"`
	Swap uint64 `json:"swap"`
}

type DiskInfo struct {
	Device     string `json:"device"`
	Mountpoint string `json:"mountpoint"`
	Fstype     string `json:"fstype"`
	Total      uint64 `json:"total"`
}

type NetInfo struct {
	PublicIp   string `json:"publicIp"`
	PrivateIp  string `json:"privateIp"`
	WsIp       string `json:"wsIp"`
	AgentIp    string `json:"agentIp"`
	ApiIp      string `json:"apiIp"`
	DetailInfo IpInfo `json:"detail"`
}

type IpInfo struct {
	Status        string  `json:"status"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Region        string  `json:"region"`
	RegionName    string  `json:"regionName"`
	City          string  `json:"city"`
	District      string  `json:"district"`
	Zip           string  `json:"zip"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	Timezone      string  `json:"timezone"`
	Isp           string  `json:"isp"`
	Org           string  `json:"org"`
	As            string  `json:"as"`
	Asname        string  `json:"asname"`
	Query         string  `json:"query"`
}
