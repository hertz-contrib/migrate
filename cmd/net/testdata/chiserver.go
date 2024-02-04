package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func main() {
	router := newRoute()
	svr := http.Server{
		Addr:         fmt.Sprintf(":%d", 8080),
		Handler:      router,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	svr.ListenAndServe()
}

func newRoute() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/api/v1/health",
		func(writer http.ResponseWriter, request *http.Request) {
			if request.URL.Query().Get("name") == "" {
				http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				return
			}
			writer.Write([]byte("Hello World!"))
		},
	)
	router.Post("/api/v1/books", _sayHello)
	router.Method(http.MethodGet, "/api/v1/books", sayHello())
	return router
}

func sayHello() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!"))
	})
}

func _sayHello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World!"))
}
