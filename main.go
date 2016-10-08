package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

	headers := r.Header
	headers["Host"] = []string{r.Host}
	json, _ := json.Marshal(r.Header)
	log.Print(r.URL.Query())
	for k, vs := range r.URL.Query() {
		values := strings.Join(vs, ",")
		w.Header().Set(k, values)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(json)

	log.Print("get request")
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe("localhost:30000", nil))
}
