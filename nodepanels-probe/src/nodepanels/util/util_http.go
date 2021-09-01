package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"nodepanels/config"
	"os"
	"strings"
	"unsafe"
)

func Post(url string, param string) {

	defer func() {
		err := recover()
		if err != nil {
			LogError("Sending POST error : " + fmt.Sprintf("%s", err))
		}
	}()

	http.DefaultTransport.(*http.Transport).DisableKeepAlives = true
	resp, _ := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(param))
	resp.Body.Close()
}

func PostJson(url string, jsonParam []byte) string {

	defer func() {
		err := recover()
		if err != nil {
			LogError("Post json error : " + fmt.Sprintf("%s", err))
		}
	}()

	request, err := http.NewRequest("POST", url, bytes.NewReader(jsonParam))
	if err != nil {
		LogError(err.Error())
		return ""
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		LogError(err.Error())
		return ""
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogError(err.Error())
		return ""
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	return *str
}

func SendCommandReceive(commandUUID string, msg string) {

	defer func() {
		err := recover()
		if err != nil {
			LogError("Sending command receive error : " + fmt.Sprintf("%s", err))
		}
	}()

	resultMap := make(map[string]string)
	resultMap["serverId"] = GetHostId()
	resultMap["commandUUID"] = commandUUID
	resultMap["msg"] = msg
	result, _ := json.Marshal(resultMap)
	PostJson(config.AgentUrl+"/command/receive", result)
}

func Get(url string) string {

	defer func() {
		err := recover()
		if err != nil {
			LogError("Sending GET error : " + fmt.Sprintf("%s", err))
		}
	}()

	http.DefaultTransport.(*http.Transport).DisableKeepAlives = true
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return string(bytes)
}

func Download(url string, target string) {

	defer func() {
		err := recover()
		if err != nil {
			LogError("Download file error : " + fmt.Sprintf("%s", err))
		}
	}()

	res, _ := http.Get(url)
	newFile, _ := os.Create(target)
	io.Copy(newFile, res.Body)
	defer res.Body.Close()
	defer newFile.Close()
}
