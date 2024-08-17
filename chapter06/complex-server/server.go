package main

import (
	"github.com/dasd412/complex-server/config"
	"github.com/dasd412/complex-server/handlers"
	"github.com/dasd412/complex-server/middleware"
	"io"
	"log"
	"net/http"
	"os"
)

func setupServer(mux *http.ServeMux, w io.Writer) http.Handler {
	conf := config.InitConfig(w)

	handlers.Register(mux, conf)
	return middleware.RegisterMiddleware(mux, conf)
}

func main() {

	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	mux := http.NewServeMux()
	wrappedMux := setupServer(mux, os.Stdout)

	log.Fatal(http.ListenAndServe(listenAddr, wrappedMux))
}
