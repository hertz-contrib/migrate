package main

import (
	"flag"
	"github.com/hertz-contrib/migrate/cmd/net/testdata/go-web/readingList/internal/data"
	"log"
)

func main() {

	var dsn string
	flag.StringVar(&dsn, "dsn", "readingList.db", "database service name")
	flag.Parse()

	r, err := data.NewSqliteRepository(dsn)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := r.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := r.Migrate(); err != nil {
		log.Fatal(err)
	}

	for _, b := range getBooks() {
		if _, err := r.InsertOne(b.Title, b.Published, b.Pages, b.Rating, b.Genres); err != nil {
			log.Fatal(err)
		}
	}

}
