package conf

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/rikaaa0928/Tieba_Sign-Go---Copy/TiebaSign"
	. "github.com/rikaaa0928/Tieba_Sign-Go---Copy/global"
)

func getCookies(cookieFileName string) (cookieJar *cookiejar.Jar, hasError bool) {
	hasError = true
	cookieJar, _ = cookiejar.New(nil)
	cookies := make([]*http.Cookie, 0)
	if _, err := os.Stat(cookieFileName); err == nil {
		rawCookie, _ := ioutil.ReadFile(cookieFileName)
		rawCookie = bytes.Trim(rawCookie, "\xef\xbb\xbf")
		rawCookieList := strings.Split(strings.Replace(string(rawCookie), "\r\n", "\n", -1), "\n")
		for _, rawCookieLine := range rawCookieList {
			rawCookieInfo := strings.SplitN(rawCookieLine, "=", 2)
			if len(rawCookieInfo) < 2 {
				continue
			}
			cookies = append(cookies, &http.Cookie{
				Name:   rawCookieInfo[0],
				Value:  rawCookieInfo[1],
				Domain: ".baidu.com",
			})
		}
		log.Printf("Verifying imported cookies from %s...", cookieFileName)
		URL, _ := url.Parse("http://baidu.com")
		cookieJar.SetCookies(URL, cookies)
		if TiebaSign.GetLoginStatus(cookieJar) {
			hasError = false
			log.Println("OK")
		} else {
			log.Println("Failed")
		}
	}
	if hasError {
		return nil, true
	}
	hasError = false
	return
}

func StartCookiesWork(cookieList map[string]*cookiejar.Jar, errorList map[string]bool) {
	log.Println("Loading and verifying Cookies from " + BasePath + "/cookies/")
	cookieFiles, e := ioutil.ReadDir(BasePath + "/cookies")
	if e != nil {
		log.Println(e)
	}

	for k := range cookieList {
		delete(cookieList, k)
	}
	for k := range errorList {
		delete(errorList, k)
	}
	for _, file := range cookieFiles {
		profileName := strings.Replace(file.Name(), ".txt", "", 1)
		cookie, hasError := getCookies(BasePath + "/cookies/" + file.Name())
		if hasError {
			log.Printf("Failed to load profile %s, invalid cookie!\n", profileName)
			errorList[profileName] = true
		} else {
			cookieList[profileName] = cookie
			errorList[profileName] = false
		}
	}
}
