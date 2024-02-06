package main

import "net/http"

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
//svc := &Config{}
//mux := http.NewServeMux()
//mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
//	//if err := svc.wj(w, r, map[string]string{"hello": "world"}); err != nil {
//	//	return
//	//}
//})
//mux.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
//if err != nil {
//	slog.Info("unable to get csrf_token cookie" + err.Error())
//}
//
//token := r.FormValue("csrf_token")

//})

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

//func (svc *service) _getUpdateOrDeleteBooks(w http.ResponseWriter, r *http.Request) {
//
//	getId := func(r *http.Request) (int64, error) {
//		id, err := strconv.ParseInt(r.URL.Path[len("api/v1/books/"):], 10, 64)
//		if err != nil {
//			http.Error(w, "", http.StatusBadRequest)
//		}
//		return id, err
//	}
//
//	switch r.Method {
//	case http.MethodGet:
//		if id, err := getId(r); err == nil {
//			svc.getBook(id, w, r)
//		}
//	case http.MethodPut:
//		if id, err := getId(r); err == nil {
//			svc.updateBook(id, w, r)
//		}
//	case http.MethodDelete:
//		if id, err := getId(r); err == nil {
//			svc.deleteBook(id, w, r)
//		}
//	default:
//		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
//	}
//}

//	func f(w http.ResponseWriter, r *http.Request) {
//		err := writeJSON(w, http.StatusCreated, envelope{"book": "book"})
//	}
//
//	func writeJSON(w http.ResponseWriter, status int, data any) error {
//		w.WriteHeader(status)
//		w.Header().Set("Content-Type", "application/json")
//		return json.NewEncoder(w).Encode(data)
//	}
//
//	func writeNotFoundOrBadRequestIfHasError(err error, w http.ResponseWriter, r *http.Request) bool {
//		if err != nil {
//			switch {
//			case errors.Is(err, data.NotFoundError):
//				writeProblemDetails(w, r, "Not Found", http.StatusNotFound, "matching book not found")
//			default:
//				writeProblemDetails(w, r, "server error", http.StatusInternalServerError, err.Error())
//			}
//			return true
//		}
//		return false
//	}
func writeProblemDetails(w http.ResponseWriter, r *http.Request, title string, statusCode int, detail string) {
	path := r.URL.Path
}

//func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		m := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$").
//			FindStringSubmatch(r.URL.Path)
//		if m == nil {
//			http.NotFound(w, r)
//			return
//		}
//		fn(w, r, m[2])
//	}
//}
