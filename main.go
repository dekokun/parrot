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

	if r.URL.Query()["Content-Type"] != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
	for k, vs := range r.URL.Query() {
		values := strings.Join(vs, ",")
		w.Header().Set(k, values)
	}

	w.WriteHeader(200)
	w.Write(json)
}

func main() {
	var port int
	flag.IntVar(&port, "port", 30000, "listen port")
	flag.Parse()
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(port)), nil))
}
