package main

import (
	"log"
	"net/http"
	"time"
)

type Middleware struct {
}

//LoggingHandler log the time-cosuming of http request
func (m Middleware) LoggingHander(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

//RecoverHandler recover panic
func (m Middleware) RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover from panic: %+v", err)
				http.Error(w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
