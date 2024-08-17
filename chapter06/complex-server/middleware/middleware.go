package middleware

import (
	"fmt"
	"github.com/dasd412/complex-server/config"
	"net/http"
	"time"
)

func loggingMiddleware(h http.Handler, c config.AppConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		h.ServeHTTP(w, r)
		requestDuration := time.Now().Sub(t1).Seconds()
		c.Logger.Printf("protocol=%s path=%s method=%s duration=%f", r.Proto, r.URL.Path, r.Method, requestDuration)
	})
}

func panicMiddleware(h http.Handler, c config.AppConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rValue := recover(); rValue != nil {
				c.Logger.Println("panic detected", rValue)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Unexpected server error occured")
			}
		}()
		h.ServeHTTP(w, r)
	})
}
