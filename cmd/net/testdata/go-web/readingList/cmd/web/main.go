package main

import (
	"flag"
	"fmt"
	"github.com/hertz-contrib/migrate/cmd/net/testdata/go-web/readingList/internal/models"
	_ "github.com/hertz-contrib/migrate/cmd/net/testdata/go-web/readingList/internal/models"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	client *models.ReadingListClient
	logger *log.Logger
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	var port int
	flag.IntVar(&port, "port", 9001, "HTTP Listen Port")
	apiEndpoint := flag.String("api-endpoint", "http://localhost:9000/api/v1/books", "API endpoint")
	flag.Parse()
	address := fmt.Sprintf("localhost:%d", port)

	app := &application{
		client: &models.ReadingListClient{
			Endpoint: *apiEndpoint,
		},
		logger: logger,
	}

	svr := http.Server{
		Addr:         address,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting server on %s", address)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
