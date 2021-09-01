package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
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
		time.Sleep(time.Second * 10)
	}
	tryCount++

	util.LogDebug("Try to establish a connection with the proxy server...")
	wsConnect, _, err := dialer.Dial(config.WsUrl+"/ws/v1/"+util.GetHostId(), nil)
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
			util.LogError("Send ws heartbeat error :" + fmt.Sprintf("%s", err))
		}
	}()

	for {

		if connect == nil {
			return
		}

		SendMsg(connect, "ping")
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
			util.LogError("Failed to connect to the proxy server：" + fmt.Sprintf("%s", err))
			util.LogError("Trying to reconnect... ")
			connect = nil
		}

		switch messageType {
		case websocket.TextMessage:
			go handleMsg(connect, string(messageData))
		case websocket.BinaryMessage:

		default:
			SendMsg(connect, "bad request")
		}
	}
}

func SendMsg(connect *websocket.Conn, msg string) {

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

func handleMsg(connect *websocket.Conn, message string) {

	defer func() {
		err := recover()
		if err != nil {
			util.LogError("HandleMsg error :" + fmt.Sprintf("%s", err))
		}
	}()

	if message != "pong" {

		command := Command{}
		json.Unmarshal([]byte(message), &command)

		if command.Tool.Type == "probe-version" {
			util.LogInfo("[COMMAND] Version probe")
			SendMsg(connect, "{\"toolType\":\"probe-version\",\"serverId\":\""+util.GetHostId()+"\",\"msg\":\""+config.Version+"\"}")
			SendMsg(connect, "{\"toolType\":\"probe-version\",\"serverId\":\""+util.GetHostId()+"\",\"msg\":\"END\"}")
		} else if command.Tool.Type == "probe-upgrade-back" {
			probe.Upgrade(command.Tool.Param)
		} else if command.Tool.Type == "probe-shutdown-back" {
			probe.ShutDown()
		} else if command.Tool.Type == "probe-cmd-back" {
			probe.ExeCmd(command.Tool.Param)
		} else {
			ExeScript(connect, command)
		}
	}

}

type Command struct {
	ServerId string      `json:"serverId"`
	Page     string      `json:"page"`
	Tool     CommandTool `json:"tool"`
}

type CommandTool struct {
	Version string `json:"version"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Timeout int    `json:"timeout"`
}
