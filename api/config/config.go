package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Port           string   `yaml:"Port"`
		Domain         string   `yaml:"Domain"`
		AllowedOrigins []string `yaml:"AllowedOrigins"`
	} `yaml:"App"`

	DB struct {
		User     string `yaml:"User"`
		Password string `yaml:"Password"`
		Database string `yaml:"Database"`
		Host     string `yaml:"Host"`
		Port     int    `yaml:"Port"`
	} `yaml:"DB"`

	Stripe struct {
		SecretKey         string `yaml:"SecretKey"`
		EndpointSecret    string `yaml:"EndpointSecret"`
		IntegrationActive bool   `yaml:"IntegrationActive"`
		SuccessURL        string `yaml:"SuccessUrl"`
		ErrorURL          string `yaml:"ErrorUrl"`
	} `yaml:"Stripe"`
}

var conf Config

func Init(filename string) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println(err)
	}
}

func GetConfig() Config {
	return conf
}

func UseStripeIntegration(value bool) {
	conf.Stripe.IntegrationActive = value
}
