package frontend

import (
	"net/http"
	"text/template"
	. "github.com/Evi1/Tieba_Sign-Go---Copy/global"
	"bytes"
	"strconv"
	"time"
	"log"
	"fmt"
	"sort"
)

type menuT struct {
	Name string
	Url  string
}

type indexT struct {
	Location string
	Time     string
	Menu     string
	Errors   int
	Counts   int
	Users    int
	Body     string
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//getTime
	ti := ""
	location := ""
	utc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println("err: ", err.Error())
		ti = err.Error()
	} else {
		tt := time.Now().In(utc)
		location = tt.Location().String()
		ti = tt.Format("2006-01-02 15:04:05")
	}


	//Creat menu
	menu := ""
	isUser := false
	keyList := make([]string, 0)
	for k := range CookieList {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, x := range keyList {
		m := menuT{Name: x, Url: "?n=" + x}
		buf := new(bytes.Buffer)
		fName := ""
		if len(r.Form["n"]) > 0 && r.Form["n"][0] == x {
			fName = BasePath + "/template/menuSelect.gtpl"
			isUser = true
		} else {
			fName = BasePath + "/template/menu.gtpl"
		}
		t, e := template.ParseFiles(fName)
		if e != nil {
			log.Println(e)
			continue
		}
		e = t.Execute(buf, m)
		if e != nil {
			log.Println(e)
			continue
		}
		menu += buf.String()
	}
	in := indexT{Menu: menu, Time: ti, Location: location}
	//Create Panel
	in.Users = len(RunList)
	in.Errors = 0
	in.Counts = 0
	for _, x1 := range RunList {
		for _, x2 := range x1 {
			in.Counts++
			if x2 == "Failed" {
				in.Errors++
			}
		}
	}
	//Create body
	if isUser {
		in.Body = userBody(r.Form["n"][0])
	} else {
		in.Body = indexBody()
	}

	t, e := template.ParseFiles(BasePath + "/template/index.html")
	if e != nil {
		log.Println(e)
		fmt.Fprintln(w, "error:"+e.Error())
		return
	}
	t.Execute(w, in)
}

func indexBody() (b string) {
	str := ""
	keyList := make([]string, 0)
	for k := range ErrorList {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		v := ErrorList[k]
		if v {
			str += makeListI(k, "Error !", "fa-user")
		} else {
			str += makeListI(k, "Fine !", "fa-user")
		}
	}
	b = makeList("UserList", str)

	str = ""
	keyList = make([]string, 0)
	for k := range RunList {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		v := RunList[k]
		n := 0
		m := 0
		for _, q := range v {
			if q != "none" && q != "Failed" {
				n++
			}
			m++
		}
		str += makeProgressI(float64(n)/float64(m)*100, strconv.Itoa(n)+"/"+strconv.Itoa(m), k)
	}
	b += makeProgress("Finished", str)
	return
}

func userBody(user string) (b string) {
	str := ""
	b = ""
	for k, v := range RunList {
		if k == user {
			keyList := make([]string, 0)
			for k := range v {
				keyList = append(keyList, k)
			}
			sort.Strings(keyList)
			for _, tb := range keyList {
				st := v[tb]
				str += makeListI(tb, st, "fa-comment")
			}
		}
	}
	b = makeList("TiebaList", str)
	return
}
