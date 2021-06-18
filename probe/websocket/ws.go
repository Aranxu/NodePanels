package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"nodepanels/config"
	"nodepanels/probe"
	"nodepanels/util"
	"time"
)

func CreateAgentConn() {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Abnormal connection with proxy program：" + fmt.Sprintf("%s", err))
		}
	}()

	dialer := websocket.Dialer{}

	tryCount := 0

TryConn:

	if tryCount != 0 {
		time.Sleep(time.Second * time.Duration(rand.Int63n(90-30)+30))
	}
	tryCount++

	util.LogDebug("Try to establish a connection with the proxy server...")
	wsConnect, _, err := dialer.Dial("wss://"+config.WsUrl+"/ws/v1/"+util.GetHostId(), nil)
	if nil != err {
		util.LogError("Failed to connect to the proxy server, trying to reconnect... ")
		goto TryConn
	}
	util.LogDebug("Successfully connected to the proxy server ")

	go readAgentMsg(wsConnect)

	go wsHeartBeat(wsConnect)

	return
}

func wsHeartBeat(connect *websocket.Conn) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Send websocket heartbeat error :" + fmt.Sprintf("%s", err))
		}
	}()

	for {

		if connect == nil {
			return
		}

		sendMsg(connect, "1")
		time.Sleep(30000 * time.Millisecond)
	}
}

func readAgentMsg(connect *websocket.Conn) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("Receive agent message error :" + fmt.Sprintf("%s", err))
		}
	}()

	for {

		if connect == nil {
			go CreateAgentConn()
			return
		}

		messageType, messageData, err := connect.ReadMessage()
		if nil != err {
			util.LogError("Failed to connect to the proxy server, trying to reconnect... ")
			connect = nil
		}

		switch messageType {
		case websocket.TextMessage:
			handleMsg(string(messageData))
		case websocket.BinaryMessage:

		default:
			sendMsg(connect, "bad request")
		}
	}
}

func sendMsg(connect *websocket.Conn, msg string) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("SendMsg error :" + fmt.Sprintf("%s", err))
		}
	}()

	if connect != nil {
		connect.WriteMessage(websocket.TextMessage, []byte(msg))
	}
}

func handleMsg(message string) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("HandleMsg error :" + fmt.Sprintf("%s", err))
		}
	}()

	util.LogDebug("Receive message from agent : " + message)

	messageMap := make(map[string]string)
	json.Unmarshal([]byte(message), &messageMap)

	command := messageMap["command"]

	commandUUID := messageMap["commandUUID"]

	serverIp := messageMap["serverIp"]

	param := messageMap["param"]
	if command == "update" {
		probe.UpdateProbe(param)
	} else if command == "version" {
		util.SendCommandReceive(commandUUID, serverIp, config.Version)
	} else if command == "restart" {
		probe.RebootProbe()
	} else if command == "processList" {
		util.SendCommandReceive(commandUUID, serverIp, probe.GetProcessesList())
	} else if command == "setWarningRule" {
		probe.SetWarningRule(param)
	} else if command == "getProcessCmd" {
		util.SendCommandReceive(commandUUID, serverIp, probe.GetProcessCmdByPid(util.String2int32(param)))
	} else if command == "setMonitorProcessRule" {
		probe.SetMonitorProcessRule(param)
	} else if command == "getProcessInfo" {
		util.SendCommandReceive(commandUUID, serverIp, probe.GetProcessInfo(util.String2int32(param)))
	}
}
