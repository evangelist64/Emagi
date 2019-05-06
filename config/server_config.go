package config

import (
	"io/ioutil"
	"encoding/json"
	)

type ServerConf struct {
	//配置内容
	Address string `json:"ip"`
	LogPath string `json:"log_path"`
	LogLevel string `json:"log_level"`
}

func (p *ServerConf) Init(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data,p);err != nil{
		return err
	}
	return nil
}