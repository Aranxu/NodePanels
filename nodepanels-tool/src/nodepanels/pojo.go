package main

type MsgMap struct {
	ToolType string `json:"toolType"`
	ServerId string `json:"serverId"`
	Msg      string `json:"msg"`
}

type ResultMap struct {
	ServerId string `json:"serverId"`
	NodeId   string `json:"nodeId"`
	Latency  string `json:"latency"`
	Download string `json:"download"`
	Upload   string `json:"upload"`
}
