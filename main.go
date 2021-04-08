package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/rikaaa0928/Tieba_Sign-Go---Copy/TiebaSign"
	"github.com/rikaaa0928/Tieba_Sign-Go---Copy/conf"
	"github.com/rikaaa0928/Tieba_Sign-Go---Copy/frontend"
	. "github.com/rikaaa0928/Tieba_Sign-Go---Copy/global"
)

var maxRetryTimes int

var f *os.File

func main() {
	maxRetryTimes = *flag.Int("retry", 4, "Max retry times for a single tieba")
	flag.Parse()

	var err error

	os.Remove(BasePath + "/logfile.log")
	f, err = os.OpenFile(BasePath+"/logfile.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v \n", err)
	}
	defer f.Close()

	log.SetOutput(f)

	go backGroundWork()
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir(BasePath+"/template"))))
	http.HandleFunc("/", frontend.HandleIndex)
	err = http.ListenAndServe(Server, nil) //设置监听的端口
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}
}

func backGroundWork() {
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(currentDir)

	log.Println("Tieba Sign (Go Version) beta")
	log.Println("Author: kookxiang <r18@ikk.me>")
	log.Println()

	conf.StartCookiesWork(CookieList, ErrorList)
	TiebaSign.StartSign(CookieList, RunList, maxRetryTimes)
	for {
		t := time.Now()
		utc, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			log.Println("err: ", err.Error())
		} else {
			t.In(utc)
		}
		if t.Minute() == 30 {
			currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
			os.Chdir(currentDir)

			log.Println("Tieba Sign start")

			if t.Hour() == 0 {
				for k := range RunList {
					delete(RunList, k)
				}
				if t.Day() == 1 {
					log.SetOutput(os.Stdout)
					f.Close()
					os.Remove(BasePath + "/logfile.log")
					f, err = os.OpenFile(BasePath+"/logfile.log", os.O_RDWR|os.O_CREATE, 0666)
					if err != nil {
						fmt.Printf("error opening file: %v \n", err)
					}
					log.SetOutput(f)
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
