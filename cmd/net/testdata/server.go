package main

import (
	"fmt"
	"net/http"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	uri := r.RequestURI
	w.WriteHeader(200)
	fmt.Fprintf(w, uri, m)
	http.Error(w, "d", http.StatusInternalServerError)
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", sayhelloName) //设置访问的路由
	svr := http.Server{
		Addr:         ":9090",
		Handler:      mux,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	svr.ListenAndServe()
}
