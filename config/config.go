package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	Config string = "config.json"
)

var (
	configLoaded bool = false
	App          AppConfig
)

type AppConfig struct {
	Server struct {
		Port        string `json:"port"`
		MetricsPort int `json:"metricsport"`
	} `json:"server"`
}

func init() {
	if configLoaded {
		return
	}
	appConfigContents, err := ioutil.ReadFile(Config)
	if err != nil {
		log.Panic("Application configuration file unreadable: ", err)
	}
	appConfig := AppConfig{}
	err = json.Unmarshal(appConfigContents, &appConfig)
	if err != nil {
		log.Panic("Application configuration file unreadable: ", err)
	}
	App = appConfig
	configLoaded = true
}
