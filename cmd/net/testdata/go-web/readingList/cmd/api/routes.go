package main

import "net/http"

func (svc *service) route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health", svc.healthcheck)
	mux.HandleFunc("/api/v1/books", svc.getOrCreateBooks)
	mux.HandleFunc("/api/v1/books/", svc.getUpdateOrDeleteBooks)
	return mux
}
