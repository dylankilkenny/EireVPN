package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		ProxyPort     string `yaml:"ProxyPort"`
		RestPort      string `yaml:"RestPort"`
		ProxyUsername string `yaml:"ProxyUsername"`
		ProxyPassword string `yaml:"ProxyPassword"`
	} `yaml:"App"`
}

var conf Config

func initConfig() Config {
	appPath, _ := os.Getwd()
	filename, _ := filepath.Abs(appPath + "/config.yaml")
	fmt.Println(filename)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println(err)
	}
	return conf
}

func GetConfig() Config {
	return initConfig()
}
