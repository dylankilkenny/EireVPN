package main

import (
	c "eirevpn/proxy/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
)

var config c.Config

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	config = c.GetConfig()
	fmt.Println(config)
	go startProxy()
	startAPI()
}

func startAPI() {
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		cred := &credentials{}
		err := d.Decode(cred)
		if err != nil {
			fmt.Println(err)
			return
		}
		config.App.ProxyUsername = cred.Username
		config.App.ProxyPassword = cred.Password
		fmt.Println(config)
	})
	fmt.Println("REST API Running")
	http.ListenAndServe(":"+config.App.RestPort, nil)
}

func startProxy() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	auth.ProxyBasic(proxy, "Auth", func(user, passwd string) bool {
		fmt.Println(config)
		if user == config.App.ProxyUsername && passwd == config.App.ProxyPassword {
			fmt.Println("correct user pass")
			return true
		}
		fmt.Println("wrong user pass")
		return false
	})
	fmt.Println("Proxy Running")
	log.Fatal(http.ListenAndServe(":"+config.App.ProxyPort, proxy))
}
