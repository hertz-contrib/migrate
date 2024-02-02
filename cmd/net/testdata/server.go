package main

import (
	"fmt"
	"net/http"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("s", "s")
	//w.Header().Del("s")
	//r.Header.Set("s", "s")
	//r.Header.Del("s")
	//get := r.URL.String()
	//uri := r.RequestURI()
	//file, fileHeader, err := r.FormFile("s")
	//m := r.Method
	//uri := r.RequestURI
	//host := r.Host
	//w.Header()
	//r.Header.Del("s")
	//w.WriteHeader(200)
	//fmt.Fprintf(w, uri, m, host)
	//http.Error(w, "d", http.StatusInternalServerError)
	//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func main() {
	// svc := &service{}
	mux := http.NewServeMux()
	//mux.HandleFunc("/hello", sayhelloName) //设置访问的路由
	svr := http.Server{
		Addr:         fmt.Sprintf(":%d", 8080),
		Handler:      svc.Route(),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	//	svr.ListenAndServe()
	//}
	//
	//func (svc *service) Route() *http.ServeMux {
	//	mux := http.NewServeMux()
	//	mux.HandleFunc("/api/v1/health", svc.healthcheck)
	//	mux.HandleFunc("/api/v1/books", svc.getOrCreateBooks)
	//	mux.HandleFunc("/api/v1/books/", svc.getUpdateOrDeleteBooks)
	//	return mux
}
