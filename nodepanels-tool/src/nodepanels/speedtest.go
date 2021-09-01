package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"nodepanels/util"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func SpeedTest(toolType string, tempNo string) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("speedtest error : " + fmt.Sprintf("%s", err))
		}
	}()

	speedtestFileName := ""
	speedtestDownloadUrl := ""
	if runtime.GOOS == "linux" && runtime.GOARCH == "amd64" {
		speedtestFileName = "speedtest"
		speedtestDownloadUrl = "https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/speedtest/speedtest-linux-64"
	} else if runtime.GOOS == "linux" && runtime.GOARCH == "arm64" {
		speedtestFileName = "speedtest"
		speedtestDownloadUrl = "https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/speedtest/speedtest-linux-arm64"
	} else if runtime.GOOS == "windows" {
		speedtestFileName = "speedtest.exe"
		speedtestDownloadUrl = "https://nodepanels-file-1256221051.cos.accelerate.myqcloud.com/speedtest/speedtest-windows-64.exe"
	}

	if !util.PathExists(filepath.Join(util.Exepath(), speedtestFileName)) {
		util.Download(speedtestDownloadUrl, filepath.Join(util.Exepath(), speedtestFileName))
		if runtime.GOOS == "linux" {
			//linux系统赋予执行权限
			os.Chmod(util.Exepath()+"/speedtest", 0777)
		}
	}

	var cmd *exec.Cmd

	//获取入参
	nodeIdStr, _ := ioutil.ReadFile(filepath.Join(util.Exepath(), toolType+"-"+tempNo+".temp"))

	nodeIds := strings.Split(string(nodeIdStr), " ")

	for _, value := range nodeIds {

		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/C", filepath.Join(util.Exepath(), speedtestFileName+" --accept-license -s "+value))
		}
		if runtime.GOOS == "linux" {
			cmd = exec.Command("/usr/bin/env bash", "-c", "/usr/local/nodepanels/speedtest")
			//cmd = exec.Command("sh", "-c", filepath.Join(util.Exepath(), speedtestFileName+" --accept-license -s "+value))
		}

		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		if err := cmd.Start(); err != nil {
		}

		go asyncLog(toolType, value, stdout)
		go asyncLog(toolType, value, stderr)

		if err := cmd.Wait(); err != nil {
		}
	}
	fmt.Println("{\"toolType\":\"" + toolType + "\",\"serverId\":\"" + util.GetHostId() + "\",\"msg\":\"END\"}")

	os.Remove(filepath.Join(util.Exepath(), toolType+"-"+tempNo+".temp"))
}

func asyncLog(toolType string, nodeId string, std io.ReadCloser) error {
	reader := bufio.NewReader(std)
	for {

		resultMap := ResultMap{}
		resultMap.NodeId = nodeId
		resultMap.ServerId = util.GetHostId()

		readString, err := reader.ReadBytes('\n')

		if err != nil || err == io.EOF {
			return err
		}
		fmt.Println(string(readString))
		if strings.Contains(string(readString), "error") {
			resultMap.Latency = "-1"
			SpeedtestSendBack(toolType, resultMap)
		} else if strings.Contains(string(readString), "ms") {
			latency := strings.Split(strings.Split(strings.ReplaceAll(strings.TrimSpace(string(readString)), " ", ""), ":")[1], "ms")[0]
			resultMap.Latency = latency
			SpeedtestSendBack(toolType, resultMap)
			if "performance-net-speedtest-ping" == toolType {
				if runtime.GOOS == "windows" {
					exec.Command("cmd", "/C", "taskkill /f /im speedtest.exe").Output()
				}
				if runtime.GOOS == "linux" {
					exec.Command("sh", "-c", "ps aux | grep \"speedtest-cli\" |grep -v grep| cut -c 9-15 | xargs kill -9").Output()
				}
			}
		} else if strings.Contains(string(readString), "Download") {
			download := strings.Split(strings.Split(strings.ReplaceAll(strings.TrimSpace(string(readString)), " ", ""), "Download:")[1], "Mbps")[0]
			resultMap.Download = download
			SpeedtestSendBack(toolType, resultMap)
		} else if strings.Contains(string(readString), "Upload") {
			upload := strings.Split(strings.Split(strings.ReplaceAll(strings.TrimSpace(string(readString)), " ", ""), "Upload:")[1], "Mbps")[0]
			resultMap.Upload = upload
			SpeedtestSendBack(toolType, resultMap)
		}

	}
}

func SpeedtestSendBack(toolType string, resultMap ResultMap) {
	msgMap := MsgMap{}
	msgMap.ToolType = toolType
	msgMap.ServerId = util.GetHostId()

	resultMsg, _ := json.Marshal(resultMap)
	msgMap.Msg = string(resultMsg)
	msg, _ := json.Marshal(msgMap)
	fmt.Println(string(msg))
}
