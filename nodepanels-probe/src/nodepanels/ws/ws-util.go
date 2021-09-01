package ws

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"nodepanels/config"
	"nodepanels/util"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

func StopScript(msg string) {

	messageMap := make(map[string]string)
	json.Unmarshal([]byte(msg), &messageMap)

	toolType := messageMap["toolType"]

	c := util.GetConfig()
	command := config.Command{
		Timeout: 0,
		Stop:    true,
	}
	c.Command[toolType] = command

	data, _ := json.MarshalIndent(c, "", "\t")
	ioutil.WriteFile(util.Exepath()+"/config", data, 0666)
}

func ExeScript(connect *websocket.Conn, command Command) {

	//toolTimeout := command.Tool.Timeout
	toolParam := command.Tool.Param
	toolType := command.Tool.Type
	toolVersion := command.Tool.Version
	toolName := command.Tool.Name
	toolDownloadUrl := command.Tool.Url

	//工具类不存在或者版本号不一致的话，下载工具类
	if !util.PathExists(filepath.Join(util.Exepath(), toolName)) || toolVersion != GetToolVersion() {
		util.Download(toolDownloadUrl, filepath.Join(util.Exepath(), toolName))
		if runtime.GOOS == "linux" {
			//linux系统赋予执行权限
			os.Chmod(util.Exepath()+"/nodepanels-tool", 0777)
		}
	}

	//生成入参临时文件
	tempNo := util.Int642string(time.Now().UnixNano())
	newFile, _ := os.Create(filepath.Join(util.Exepath(), toolType+"-"+tempNo+".temp"))
	newFile.Write([]byte(toolParam))
	newFile.Close()

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", filepath.Join(util.Exepath(), toolName)+" -"+toolType+" "+tempNo)
	}
	if runtime.GOOS == "linux" {
		cmd = exec.Command("sh", "-c", filepath.Join(util.Exepath(), toolName)+" -"+toolType+" "+tempNo)
	}

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		util.LogError("Error starting command: " + fmt.Sprintf("%s", err))
	}

	go asyncLog(connect, stdout)
	go asyncLog(connect, stderr)

	if err := cmd.Wait(); err != nil {
		util.LogError("Error waiting for command execution: " + fmt.Sprintf("%s", err))
	}

}

func asyncLog(connect *websocket.Conn, std io.ReadCloser) error {
	reader := bufio.NewReader(std)
	for {
		readString, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			return err
		}
		SendMsg(connect, string(readString))
	}
}

func GetToolVersion() string {
	if runtime.GOOS == "windows" {
		output, _ := exec.Command("cmd", "/C", filepath.Join(util.Exepath(), "nodepanels-tool.exe")+" -version").Output()
		return string(output)
	}
	if runtime.GOOS == "linux" {
		output, _ := exec.Command("sh", "-c", filepath.Join(util.Exepath(), "nodepanels-tool")+" -version").Output()
		return string(output)
	}
	return ""
}
