package config

const Version = "v1.0.2"

var WsUrl = "wss://ws.nodepanels.com"
var AgentUrl = "https://agent.nodepanels.com"
var ApiUrl = "https://collect.nodepanels.com"

type Config struct {
	ServerId string             `json:"serverId"`
	Warning  Warning            `json:"warning"`
	Monitor  Monitor            `json:"monitor"`
	Command  map[string]Command `json:"command"`
}

type Warning struct {
	Switch int         `json:"switch"`
	Rule   WarningRule `json:"rule"`
}

type WarningRule struct {
	Cpu WarningRuleCpu `json:"cpu"`
	Mem WarningRuleMem `json:"mem"`
}

type WarningRuleCpu struct {
	Switch   int `json:"switch"`
	Value    int `json:"value"`
	Duration int `json:"duration"`
	Count    int `json:"count"`
}

type WarningRuleMem struct {
	Switch   int `json:"switch"`
	Value    int `json:"value"`
	Duration int `json:"duration"`
	Count    int `json:"count"`
}

type Monitor struct {
	Rule MonitorRule `json:"rule"`
}

type MonitorRule struct {
	Process []string `json:"process"`
}

type Command struct {
	Timeout int  `json:"timeout"`
	Stop    bool `json:"stop"`
}
