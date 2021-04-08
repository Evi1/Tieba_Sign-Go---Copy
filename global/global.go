package global

import (
	"net/http/cookiejar"
	"os"
	"path/filepath"
)

var CookieList map[string]*cookiejar.Jar
var ErrorList map[string]bool
var BasePath string
var RunList map[string]map[string]string
var Server string

func init() {
	CookieList = make(map[string]*cookiejar.Jar)
	ErrorList = make(map[string]bool)
	RunList = make(map[string]map[string]string)
	ip := os.Getenv("OPENSHIFT_GO_IP")
	port := os.Getenv("OPENSHIFT_GO_PORT")
	if len(ip) > 0 && len(port) > 0 {
		Server = ip + ":" + port
		BasePath = "../."
	} else {
		Server = ":60080"
		BasePath = filepath.ToSlash(os.Getenv("GOPATH")) + "/src/github.com/rikaaa0928/Tieba_Sign-Go---Copy"
	}
}
