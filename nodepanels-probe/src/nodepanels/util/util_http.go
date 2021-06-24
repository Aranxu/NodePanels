package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func SendCommandReceive(commandUUID string, commandIp string, msg string) {

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
	PostJson("https://"+commandIp+"/command/receive", result)
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
