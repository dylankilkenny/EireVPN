module eirevpn/proxy

go 1.12

require (
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5
	github.com/elazarl/goproxy v0.0.0-20190711103511-473e67f1d7d2
	github.com/elazarl/goproxy/ext v0.0.0-20190711103511-473e67f1d7d2
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	gopkg.in/yaml.v2 v2.2.7
)

replace github.com/elazarl/goproxy => github.com/dylankilkenny/goproxy v0.0.0-20200109204127-1c107847a855

replace github.com/elazarl/goproxy/ext => github.com/dylankilkenny/goproxy/ext v0.0.0-20200109204127-1c107847a855
