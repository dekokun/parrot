package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
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
	type MyType struct {
		Hoge []string
	}

	var mt MyType
	json.Unmarshal(data, &mt)
	if !reflect.DeepEqual(mt.Hoge, []string{"Fuga"}) {
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

func TestSetContentType(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()
	r, err := http.Get(ts.URL + "?Content-Type=text/plain")
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	if r.StatusCode != 200 {
		t.Fatalf("Error by status code. %v", r.Status)
	}
	if r.Header["Content-Type"][0] != "text/plain" {
		t.Fatalf("Error by header. %v", r.Header)
	}
}

func TestStatusCode(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()
	client := &http.Client{}
	req, err := http.NewRequest("GET", ts.URL, nil)
	req.Header.Add("Status", "503")
	r, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}
	if r.StatusCode != 503 {
		t.Fatalf("Error by status code. %v", r.Status)
	}
	type MyType struct {
		Status []string
	}
	var mt MyType
	json.Unmarshal(data, &mt)
	if !reflect.DeepEqual(mt.Status, []string{"503"}) {
		t.Fatalf("Error by not contains header. %v", string(data))
	}
}
