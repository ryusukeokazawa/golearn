package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type ResponseData struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	omikujis := []string{"大吉", "中吉", "小吉", "凶"}
	response := ResponseData{http.StatusOK, omikujis[rand.Intn(4)]}

	res, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// "/"
	fmt.Fprintf(w, "omikuji")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/omikuji", omikujiHandler)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
