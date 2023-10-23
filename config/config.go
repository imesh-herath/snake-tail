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
		MetricsPort int    `json:"metricsport"`
	} `json:"server"`

	Firebase struct {
		Url               string `json:"url"`
		ApiKey            string `json:"apikey"`
		AuthDomain        string `json:"authdomain"`
		DatabaseURL       string `json:"databaseurl"`
		ProjectId         string `json:"projectid"`
		StorageBucket     string `json:"storagebucket"`
		MessagingSenderId string `json:"messagingsenderid"`
		AppId             string `json:"appid"`
		MeasurementId     string `json:"measurementid"`
	} `json:"firebase"`
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
