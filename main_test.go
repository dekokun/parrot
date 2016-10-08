package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var sampleHandler = http.HandlerFunc(rootHandler)

func TestNormal(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()
	// リクエストの送信先はテストサーバのURLへ。
	r, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	if r.StatusCode != 200 {
		t.Fatalf("Error by status code. %v", r.Status)
	}
}
