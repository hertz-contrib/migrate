package testdata

import (
	"fmt"
	"net/http"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello afei!") //这个写入到w的是输出到客户端的
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayhelloName) //设置访问的路由
	svr := http.Server{
		Addr:         ":9090",
		Handler:      mux,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	svr.ListenAndServe()
}
