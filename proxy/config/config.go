package config

import (
	"fmt"
	"io/ioutil"

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

var configFilename string

func Init(filename string) {
	configFilename = filename
}

func Load() Config {
	conf := Config{}
	yamlFile, err := ioutil.ReadFile(configFilename)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println(err)
	}
	return conf
}

func (c *Config) SaveConfig() error {
	newConf, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("config.yaml", newConf, 0644)
	if err != nil {
		return err
	}
	return nil
}
