package main

import (
	"flag"
	"fmt"
	"github.com/hertz-contrib/migrate/cmd/net/testdata/go-web/readingList/internal/data"
	"log"
	"net/http"
	"os"
	"time"
)

type options struct {
	port int
	env  string
	dsn  string
}

type service struct {
	settings   options
	logger     *log.Logger
	repository data.Repository
}

func main() {
	var serverOptions options

	flag.IntVar(&serverOptions.port, "port", 9000, "HTTP Listen Port")
	flag.StringVar(&serverOptions.env, "env", "production", "server environment")
	flag.StringVar(&serverOptions.dsn, "dsn", "readingList.db", "database service name")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	repository, err := data.NewSqliteRepository(serverOptions.dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := repository.Ping(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := repository.Close; err != nil {
			log.Printf("%v", err)
		}
	}()

	svc := &service{
		settings:   serverOptions,
		logger:     logger,
		repository: repository,
	}
	serverAddress := fmt.Sprintf(":%d", svc.settings.port)

	svr := http.Server{
		Addr:         serverAddress,
		Handler:      svc.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s", svc.settings.env, serverAddress)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
