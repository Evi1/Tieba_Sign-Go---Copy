package frontend

import (
	"bytes"
	"text/template"
	. "github.com/Evi1/Tieba_Sign-Go---Copy/global"
	"log"
)

type listT struct {
	Name   string
	Inside string
}

type listIT struct {
	X string
	Y string
	T string
}
type proT struct {
	Title  string
	Inside string
}

type proIT struct {
	P float64
	N string
	T string
}

func makeListI(x, y, tp string) (b string) {
	a := listIT{X: x, Y: y, T: tp}
	buf := new(bytes.Buffer)
	t, e := template.ParseFiles(BasePath + "/template/listInside.gtpl")
	if e != nil {
		log.Println(e)
	}
	e = t.Execute(buf, a)
	if e != nil {
		log.Println(e)
	}
	b = buf.String()
	return
}
func makeList(name, inside string) (b string) {
	a := listT{Name: name, Inside: inside}
	buf := new(bytes.Buffer)
	t, e := template.ParseFiles(BasePath + "/template/list.gtpl")
	if e != nil {
		log.Println(e)
		return ""
	}
	e = t.Execute(buf, a)
	if e != nil {
		log.Println(e)
		return ""
	}
	b = buf.String()
	return
}
func makeProgressI(p float64, n string, ti string) (b string) {
	a := proIT{P: p, N: n, T: ti}
	buf := new(bytes.Buffer)
	t, e := template.ParseFiles(BasePath + "/template/proInside.gtpl")
	if e != nil {
		log.Println(e)
	}
	e = t.Execute(buf, a)
	if e != nil {
		log.Println(e)
	}
	b = buf.String()
	return
}
func makeProgress(title, inside string) (b string) {
	a := proT{Title: title, Inside: inside}
	buf := new(bytes.Buffer)
	t, e := template.ParseFiles(BasePath + "/template/progress.gtpl")
	if e != nil {
		log.Println(e)
		return ""
	}
	e = t.Execute(buf, a)
	if e != nil {
		log.Println(e)
		return ""
	}
	b = buf.String()
	return
}
