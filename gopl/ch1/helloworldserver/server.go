package main

import (
	"../gifdemo"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//func main() {
//	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//
//		fmt.Fprintf(writer, "<h1>Hello world %s<h1>",
//			request.FormValue("name"))
//	})
//
//	http.ListenAndServe(":8888", nil)
//}

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/gif", func(w http.ResponseWriter, r *http.Request) {
		cycles, _ := strconv.Atoi(r.Form.Get("cycles"))
		gifdemo.Lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] =%q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAdder = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path= %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "<h1>Count %d\n<h1>", count)
	mu.Unlock()
}
