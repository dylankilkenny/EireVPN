package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var configFilename string

type Config struct {
	App struct {
		Port                string   `yaml:"Port"`
		Domain              string   `yaml:"Domain"`
		JWTSecret           string   `yaml:"JWTSecret"`
		AllowedOrigins      []string `yaml:"AllowedOrigins"`
		EnableCSRF          bool     `yaml:"EnableCSRF"`
		EnableSubscriptions bool     `yaml:"EnableSubscriptions"`
		EnableAuth          bool     `yaml:"EnableAuth"`
		AuthCookieAge       int      `yaml:"AuthCookieAge"`
		RefreshCookieAge    int      `yaml:"RefreshCookieAge"`
		AuthCookieName      string   `yaml:"AuthCookieName"`
		RefreshCookieName   string   `yaml:"RefreshCookieName"`
		AuthTokenExpiry     int      `yaml:"AuthTokenExpiry"`
		RefreshTokenExpiry  int      `yaml:"RefreshTokenExpiry"`
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

	SendGrid struct {
		APIKey            string `yaml:"APIKey"`
		IntegrationActive bool   `yaml:"IntegrationActive"`
		Templates         struct {
			Registration   string `yaml:"Registration"`
			SupportRequest string `yaml:"SupportRequest"`
		} `yaml:"Templates"`
	} `yaml:"SendGrid"`
}

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
