package main

import (
	c "eirevpn/proxy/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	appPath, _ := os.Getwd()
	filename, _ := filepath.Abs(appPath + "/config.yaml")
	c.Init(filename)
	go startProxy()
	startAPI()
}

func startAPI() {
	http.HandleFunc("/update_creds", func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		cred := &credentials{}
		err := d.Decode(cred)
		if err != nil {
			fmt.Println(err)
			return
		}
		config := c.Load()
		config.App.ProxyUsername = cred.Username
		config.App.ProxyPassword = cred.Password
		if err = config.SaveConfig(); err != nil {
			fmt.Println("Error saving config: ", err)
		}
		fmt.Println("Updated Configuration: ", config)
	})
	config := c.Load()
	fmt.Println("REST API Started")
	http.ListenAndServe(":"+config.App.RestPort, nil)
}

func startProxy() {
	config := c.Load()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	auth.ProxyBasic(proxy, "Auth", func(user, passwd string) bool {
		if user == config.App.ProxyUsername && passwd == config.App.ProxyPassword {
			fmt.Println("Authenticated, allowing connection.")
			return true
		}
		fmt.Printf("Wrong Credentials: %s:%s \n", user, passwd)
		return false
	})
	fmt.Println("Proxy Started")
	log.Fatal(http.ListenAndServe(":"+config.App.ProxyPort, proxy))
}
