package main

import (
	"encoding/json"
	"net/http"
)

//func sayhelloName(w http.ResponseWriter, r *http.Request) {
//getId := func(r *http.Request, intr int) (int64, error) {
//	println(intr)
//	id, err := strconv.ParseInt(r.URL.Path[len("api/v1/books/"):], 10, 64)
//	if err != nil {
//		http.Error(w, "", http.StatusBadRequest)
//	}
//	return id, err
//}
//r.Form.Get("s")
//form := r.MultipartForm.Value
//form2 := r.MultipartForm.File
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
//header := r.Header
//w.Header()
//r.Header.Del("s")
//w.WriteHeader(200)
//fmt.Fprintf(w, uri, m, host)
//http.Error(w, "d", http.StatusInternalServerError)
//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
//w.Write([]byte("Hello World!"))
//}

//	func __sayhelloName() http.Handler {
//		println()
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			w.Write([]byte("Hello World!"))
//		})
//	}
//func ___sayhelloName() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(400)
//		w.Write([]byte("Hello World!"))
//		fmt.Fprintf(w, "Hello World!")
//	}
//}

type Config struct {
	Addr string
}

//func main() {
//	svc := &Config{}
//	mux := http.NewServeMux()
//	//mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
//	//if err := svc.wj(w, r, map[string]string{"hello": "world"}); err != nil {
//	//	return
//	//}
//	//})
//	mux.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
//		getId := func(r *http.Request) (int64, error) {
//			return 0, nil
//		}
//		switch "s" {
//		case "s":
//			if i, err := getId(r); err != nil {
//				println(i)
//			}
//		}
//		//if r.RequestURI == "localhost" {
//		//
//		//}
//		//if r.RequestURI == "/api/v1/health" {
//		//
//		//}
//		//if r.Method != http.MethodGet {
//		//
//		//}
//		//method := r.Method
//		//w.Write([]byte("Hello World!"))
//		//expr := "s"
//		//
//		//switch expr {
//		//case "S":
//		//	svc.wj(w, r, map[string]string{"hello": "world"})
//		//case "SS":
//		//	svc.wj(w, r, map[string]string{"hello": "world"})
//		//}
//		//if expr == "S" {
//		//if expr == "" {
//		//	svc.wj(w, r, map[string]string{"hello": "world"})
//		//}
//		//if svc.wjbool(w, r, map[string]string{"hello": "world"}) {
//		//	return
//		//}
//	})
//	//cfg := &Config{
//	//	Addr: "8080",
//	//}
//	//svr := http.Server{
//	//	Addr:         cfg.Addr,
//	//	IdleTimeout:  1 * time.Minute,
//	//	ReadTimeout:  10 * time.Second,
//	//	WriteTimeout: 30 * time.Second,
//	//}
//	//svr.ListenAndServe()
//}

//func Route() *http.ServeMux {
//	mux := http.NewServeMux()
//	return mux
//}

// func (svc *Config) wj(w http.ResponseWriter, r *http.Request, data any) error {
// w.Header().Set("Content-Type", "application/json")
// marshal, err := json.Marshal(data)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
// return nil
// }
//
//	func (svc *Config) wjbool(r *http.Request) (error, bool) {
//		getId := func(r *http.Request) (int64, error) {
//			return 0, nil
//		}
//		return nil, false
//	}
func f(w http.ResponseWriter, r *http.Request) {
	err := writeJSON(w, http.StatusCreated, envelope{"book": "book"})
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
