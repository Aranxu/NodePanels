package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nodepanels/config"
	"strings"
)

func GetHostId() string {

	defer func() {
		err := recover()
		if err != nil {
			LogError("Get host id error : " + fmt.Sprintf("%s", err))
		}
	}()

	f, err := ioutil.ReadFile(Exepath() + "/config")
	if err != nil {
		return ""
	}

	config := config.Config{}
	err = json.Unmarshal(f, &config)
	if err != nil {
		return ""
	}

	return strings.Split(config.ServerId, "\n")[0]
}

func GetConfig() config.Config {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Get config error : " + fmt.Sprintf("%s", err))
		}
	}()

	f, err := ioutil.ReadFile(Exepath() + "/config")
	if err != nil {
		return config.Config{}
	}

	c := config.Config{}
	err = json.Unmarshal(f, &c)
	if err != nil {
		return config.Config{}
	}

	return c
}
