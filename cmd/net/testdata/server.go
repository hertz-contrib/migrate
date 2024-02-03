package main

import (
	"net/http"
)

//func sayhelloName(w http.ResponseWriter, r *http.Request) {
//	//r.Form.Get("s")
//	//form := r.MultipartForm.Value
//	//form2 := r.MultipartForm.File
//	//w.Header().Add("s", "s")
//	//w.Header().Del("s")
//	//r.Header.Set("s", "s")
//	//r.Header.Del("s")
//	//get := r.URL.String()
//	//uri := r.RequestURI()
//	//file, fileHeader, err := r.FormFile("s")
//	//m := r.Method
//	//uri := r.RequestURI
//	//host := r.Host
//	//header := r.Header
//	//w.Header()
//	//r.Header.Del("s")
//	//w.WriteHeader(200)
//	//fmt.Fprintf(w, uri, m, host)
//	//http.Error(w, "d", http.StatusInternalServerError)
//	//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
//	w.Write([]byte("Hello World!"))
//}

//func __sayhelloName() http.Handler {
//	println()
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("Hello World!"))
//	})
//}
//
//func ___sayhelloName() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("Hello World!"))
//	}
//}

func main() {
	//svc := &service{}
	mux := http.NewServeMux()
	//router := chi.NewRouter()
	//router.Get("/api/v1/health", ___sayhelloName())
	//router.Method(http.MethodGet, "/api/v1/books", __sayhelloName())
	//mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello World!"))
	//})
	mux.HandleFunc("/api/v1/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}))
	//svr := http.Server{
	//	Addr:         fmt.Sprintf(":%d", 8080),
	//	Handler:      svc.Route(),
	//	IdleTimeout:  1 * time.Minute,
	//	ReadTimeout:  10 * time.Second,
	//	WriteTimeout: 30 * time.Second,
	//}
	//svr.ListenAndServe()
}

func (svc *service) Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health", svc.healthcheck)
	mux.HandleFunc("/api/v1/books", svc.getOrCreateBooks)
	mux.HandleFunc("/api/v1/books/", svc.getUpdateOrDeleteBooks)
	return mux
}
