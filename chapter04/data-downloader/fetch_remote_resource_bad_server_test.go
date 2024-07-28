package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func startBadTestHttpServer() *httptest.Server {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(30 * time.Second)
		fmt.Fprint(w, "Hello World")
	}))

	return ts
}

func TestBadFetchRemoteResource(t *testing.T) {
	ts := startBadTestHttpServer()

	defer ts.Close()

	expected := "Hello World"

	data, err := fetchRemoteResource(ts.URL)

	if err != nil {
		t.Fatal(err)
	}

	if expected != string(data) {
		t.Errorf("Expected response to be: %s, Got: %s", expected, data)
	}
}
