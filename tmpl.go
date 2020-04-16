package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var tmpl = template.Must(template.New("msg").
	Parse("<html><body>{{.Name}}さんの運勢は「<b>{{.Omikuji}}</b>」です</body></html>"))

type Result struct {
	Name    string
	Omikuji string
}

func handler(w http.ResponseWriter, r *http.Request) {
	result := Result{
		Name:    r.FormValue("p"),
		Omikuji: omikuji(),
	}
	tmpl.Execute(w, result)
}

func main() {
	// 現在時刻
	t := time.Now().UnixNano()
	rand.Seed(t)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func omikuji() string {
	n := rand.Intn(6) // 0-5
	switch n + 1 {
	case 6:
		return "大吉"
	case 5, 4:
		return "中吉"
	case 3, 2:
		return "小吉"
	default:
		return "凶"
	}
}
