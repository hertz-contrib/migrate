package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/hertz-contrib/migrate/cmd/net/testdata/internal/data"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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
	flag.StringVar(&serverOptions.dsn, "dsn", "readingList.db", "database _service name")
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

func (svc *service) route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health", svc.healthcheck)
	mux.HandleFunc("/api/v1/books", svc.getOrCreateBooks)
	mux.HandleFunc("/api/v1/books/", svc.getUpdateOrDeleteBooks)
	return mux
}

//------------------utils

type envelope map[string]any

func (svc *service) writeJSON(w http.ResponseWriter, status int, data any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func (svc *service) readJSONObject(w http.ResponseWriter, r *http.Request, dto any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var maxBytes int64 = 2_097_152 // 2MB
	http.MaxBytesReader(w, r.Body, maxBytes)

	if err := decoder.Decode(dto); err != nil {
		return err
	}

	err := decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body should  only contain a single object")
	}
	return nil
}

// ------------------ svc logic
func (svc *service) writeBadRequest(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, "", http.StatusBadRequest)
}

func (svc *service) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	status := map[string]string{
		"status":      "available",
		"environment": svc.settings.env,
	}

	err := svc.writeJSON(w, http.StatusOK, status)
	if err != nil {
		svc.logger.Print(err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (svc *service) getOrCreateBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		books, err := svc.repository.FindAll(true)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		if err := svc.writeJSON(w, http.StatusOK, envelope{"books": books}); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}
	if r.Method == http.MethodPost {
		var bookDto struct {
			Title     string   `json:"title"`
			Published int      `json:"published"`
			Pages     int      `json:"pages"`
			Genres    []string `json:"genres"`
			Rating    float64  `json:"rating"`
		}

		if err := svc.readJSONObject(w, r, &bookDto); err != nil {
			// custom error handling details: Alex Edwards, Let's Go Further Chapter 4
			svc.writeBadRequest(w, err)
			return
		}

		book, err := svc.repository.
			InsertOne(bookDto.Title, bookDto.Published, bookDto.Pages, bookDto.Rating, bookDto.Genres)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
		}
		if err := svc.writeJSON(w, http.StatusCreated, envelope{"book": book}); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (svc *service) getUpdateOrDeleteBooks(w http.ResponseWriter, r *http.Request) {

	getId := func(r *http.Request) (int64, error) {
		id, err := strconv.ParseInt(r.URL.Path[len("api/v1/books/"):], 10, 64)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
		}
		return id, err
	}

	switch r.Method {
	case http.MethodGet:
		if id, err := getId(r); err == nil {
			svc.getBook(id, w, r)
		}
	case http.MethodPut:
		if id, err := getId(r); err == nil {
			svc.updateBook(id, w, r)
		}
	case http.MethodDelete:
		if id, err := getId(r); err == nil {
			svc.deleteBook(id, w, r)
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (svc *service) getBook(id int64, w http.ResponseWriter, r *http.Request) {
	_ = r

	book, err := svc.repository.FindById(id, true)
	if err != nil {
		switch {
		case errors.Is(err, data.NotFoundError):
			http.Error(w, "", http.StatusNotFound)
			return
		default:
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}

	if err := svc.writeJSON(w, http.StatusOK, envelope{"book": book}); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}
func (svc *service) updateBook(id int64, w http.ResponseWriter, r *http.Request) {
	var bookDto struct {
		Title     *string  `json:"title"`
		Published *int     `json:"published"`
		Pages     *int     `json:"pages"`
		Genres    []string `json:"genres"`
		Rating    *float64 `json:"rating"`
	}

	book, err := svc.repository.FindById(id, true)
	if err != nil {
		switch {
		case errors.Is(err, data.NotFoundError):
			http.Error(w, "", http.StatusNotFound)
		default:
			http.Error(w, "", http.StatusInternalServerError)
		}
	}

	if err := svc.readJSONObject(w, r, &bookDto); err != nil {
		svc.writeBadRequest(w, err)
		return
	}

	if bookDto.Title != nil {
		book.Title = *bookDto.Title
	}
	if bookDto.Published != nil {
		book.Published = *bookDto.Published
	}
	if bookDto.Pages != nil {
		book.Pages = *bookDto.Pages
	}
	if len(bookDto.Genres) > 0 {
		book.Genres = bookDto.Genres
	}
	if bookDto.Rating != nil {
		book.Rating = *bookDto.Rating
	}

	if err := svc.repository.Update(book); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
	if err := svc.writeJSON(w, http.StatusOK, envelope{"book": book}); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}
func (svc *service) deleteBook(id int64, w http.ResponseWriter, r *http.Request) {
	_ = r

	err := svc.repository.DeleteById(id)
	if err != nil {
		switch {
		case errors.Is(err, data.NotFoundError):
			http.Error(w, "", http.StatusNotFound)
		default:
			http.Error(w, "", http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
