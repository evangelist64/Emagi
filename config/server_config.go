package config

import (
	"encoding/json"
	"io/ioutil"
)

type ServerConf struct {
	//配置内容
	Address string `json:"ip"`
	LogPath string `json:"log_path"`
	IsDebug bool `json:"is_debug"`
}

func (p *ServerConf) Init(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, p); err != nil {
		return err
	}
	return nil
}
