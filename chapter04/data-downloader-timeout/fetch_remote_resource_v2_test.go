package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func startBadTestHttpServerV2(shutdownServer chan struct{}) *httptest.Server {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-shutdownServer
		fmt.Fprint(w, "Hello World")
	}))

	return ts
}

func TestBadFetchRemoteResourceV2(t *testing.T) {
	shutdownServer := make(chan struct{})

	ts := startBadTestHttpServerV2(shutdownServer)

	defer ts.Close()

	defer func() {
		shutdownServer <- struct{}{}
	}()

	client := createHttpClientWithTimeout(200 * time.Millisecond)

	_, err := fetchRemoteResource(client, ts.URL)

	if err == nil {
		t.Fatal("Expected non-nil error")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("Expected error to contain: context deadline exceeded, Got: %v", err.Error())
	}
}
