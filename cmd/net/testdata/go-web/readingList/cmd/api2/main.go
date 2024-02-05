package main

import (
	"flag"
	"fmt"
	"github.com/hertz-contrib/migrate/cmd/net/testdata/go-web/readingList/internal/data"
	"log"
	"net/http"
)

type options struct {
	port int
	env  string
	dsn  string
}

func main() {
	var serverOptions options

	flag.IntVar(&serverOptions.port, "port", 9000, "HTTP Listen Port")
	flag.StringVar(&serverOptions.env, "env", "production", "server environment")
	flag.StringVar(&serverOptions.dsn, "dsn", "readingList.db", "database service name")
	flag.Parse()
	serverAddress := fmt.Sprintf(":%d", serverOptions.port)

	repository, err := data.NewSqliteRepository(serverOptions.dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := repository.Close(); err != nil {
			log.Print(err.Error())
			return
		}
	}()

	api := NewApi(repository)

	log.Printf("Server started")
	router := NewRouter(api)

	log.Fatal(http.ListenAndServe(serverAddress, router))
}
