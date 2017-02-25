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

func init() {
	CookieList = make(map[string]*cookiejar.Jar)
	ErrorList = make(map[string]bool)
	RunList = make(map[string]map[string]string)
	BasePath = filepath.ToSlash(os.Getenv("GOPATH")) + "/src/github.com/Evi1/Tieba_Sign-Go---Copy"
}
