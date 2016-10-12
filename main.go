package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

	headers := r.Header
	headers["Host"] = []string{r.Host}
	json, _ := json.Marshal(r.Header)
	log.Print(r.URL.Query())

	if _, ok := r.URL.Query()["Content-Type"]; ok {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
	for k, vs := range r.URL.Query() {
		values := strings.Join(vs, ",")
		w.Header().Set(k, values)
	}

	w.WriteHeader(200)
	w.Write(json)

	log.Print("get request")
}

func main() {
	var port int
	flag.IntVar(&port, "port", 30000, "listen port")
	flag.Parse()
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", strconv.Itoa(port)), nil))
}
