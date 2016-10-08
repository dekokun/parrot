package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var sampleHandler = http.HandlerFunc(rootHandler)

func TestNormal(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	if r.StatusCode != 200 {
		t.Fatalf("Error by status code. %v", r.Status)
	}
}

func TestBody(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()
	client := &http.Client{}
	req, err := http.NewRequest("GET", ts.URL, nil)
	req.Header.Add("Hoge", "Fuga")
	r, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}
	if r.StatusCode != 200 {
		t.Fatalf("Error by status code. %v", r.Status)
	}
	if !strings.Contains(string(data), `"Hoge":["Fuga"]`) {
		t.Fatalf("Error by not contains header. %v", string(data))
	}
}

func TestHeader(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()
	r, err := http.Get(ts.URL + "?hoge=fuga")
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	if r.StatusCode != 200 {
		t.Fatalf("Error by status code. %v", r.Status)
	}
	if r.Header["Hoge"][0] == "Fuga" {
		t.Fatalf("Error by header. %v", r.Header)
	}
}
