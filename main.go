package main

import (
	"github.com/Evi1/Tieba_Sign-Go---Copy/TiebaSign"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"github.com/Evi1/Tieba_Sign-Go---Copy/conf"
	"time"
	"net/http"
	"github.com/Evi1/Tieba_Sign-Go---Copy/frontend"
	. "github.com/Evi1/Tieba_Sign-Go---Copy/global"
)

var maxRetryTimes int

func main() {
	maxRetryTimes = *flag.Int("retry", 4, "Max retry times for a single tieba")
	flag.Parse()

	go backGroundWork()
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir(BasePath+"/template"))))
	http.HandleFunc("/", frontend.HandleIndex)
	err := http.ListenAndServe(":60080", nil) //设置监听的端口
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}
}

func backGroundWork() {
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(currentDir)

	fmt.Println("Tieba Sign (Go Version) beta")
	fmt.Println("Author: kookxiang <r18@ikk.me>")
	fmt.Println()

	conf.StartCookiesWork(CookieList, ErrorList)
	TiebaSign.StartSign(CookieList, RunList, maxRetryTimes)
	for {
		t := time.Now()
		utc, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			fmt.Println("err: ", err.Error())
		} else {
			t.In(utc)
		}
		if t.Minute() == 30 {
			currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
			os.Chdir(currentDir)

			fmt.Println("Tieba Sign start")

			if t.Hour() == 0 {
				for k := range RunList {
					delete(RunList, k)
				}
			}
			if t.Hour()%3 == 0 {
				conf.StartCookiesWork(CookieList, ErrorList)
			}
			TiebaSign.StartSign(CookieList, RunList, maxRetryTimes)
			time.Sleep(1 * time.Minute)
		} else {
			time.Sleep(1 * time.Second)
		}

	}
}
