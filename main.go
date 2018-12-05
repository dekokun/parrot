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
	statusCode := 200
	if r.Header["X-Parrot-Status"] != nil {
		var err error
		statusCode, err = strconv.Atoi(r.Header["X-Parrot-Status"][0])
		if err != nil {
			log.Print("invalid status code. set 200.")
			statusCode = 200
		}
	}

	if r.URL.Query()["Content-Type"] != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
	for k, vs := range r.URL.Query() {
		values := strings.Join(vs, ",")
		w.Header().Set(k, values)
	}

	w.WriteHeader(statusCode)
	w.Write(json)
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "listen port")
	flag.Parse()
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(port)), nil))
}
