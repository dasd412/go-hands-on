package handlers

import (
	"bytes"
	"github.com/dasd412/complex-server/config"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/api", nil)
	w := httptest.NewRecorder()

	b := new(bytes.Buffer)
	c := config.InitConfig(b)

	apiHandler(w, r, c)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected response status: %v, Got: %v\n", http.StatusOK, resp.StatusCode)
	}

	expectedResponseBody := "Hello, world!"

	if string(body) != expectedResponseBody {
		t.Errorf("Expected response: %s, Got: %s\n", expectedResponseBody, string(body))
	}
}
