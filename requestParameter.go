package main

import (
	"log"
	"math/rand"
	"net/http"
)

type Result struct {
	Name    string
	Omikuji string
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	result := Result{
		Name:    r.FormValue("p"),
		Omikuji: omikuji(),
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func omikuji() string {
	n := rand.Intn(6)
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
