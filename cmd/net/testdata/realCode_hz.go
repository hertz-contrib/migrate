package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/migrate/cmd/net/testdata/internal/data"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type _options struct {
	port int
	env  string
	dsn  string
}

type _service struct {
	settings   _options
	logger     *log.Logger
	repository data.Repository
}

func main() {
	var serverOptions _options

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

	svc := &_service{
		settings:   serverOptions,
		logger:     logger,
		repository: repository,
	}
	serverAddress := fmt.Sprintf(":%d", svc.settings.port)

	svr := http.Server{
		Addr: serverAddress,
		//Handler:      svc.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s", svc.settings.env, serverAddress)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (svc *_service) route() *server.Hertz {
	mux := server.Default()
	mux.Any("/api/v1/health", svc.healthcheck)
	mux.Any("/api/v1/books", svc.getOrCreateBooks)
	mux.Any("/api/v1/books/", svc.getUpdateOrDeleteBooks)
	return mux
}

//------------------utils

type _envelope map[string]any

func (svc *_service) writeJSON(c *app.RequestContext, status int, data any) error {
	//w.WriteHeader(status)
	c.SetStatusCode(status)
	//w.Header().Set("Content-Type", "application/json")
	c.Response.Header.Set("Content-Type", "application/json")
	c.JSON(status, data)
	return nil
}

func (svc *_service) readJSONObject(c *app.RequestContext, dto any) error {
	//decoder := json.NewDecoder(r.Body)
	//decoder.DisallowUnknownFields()
	//var maxBytes int64 = 2_097_152 // 2MB
	//http.MaxBytesReader(w, r.Body, maxBytes)
	//
	//if err := decoder.Decode(dto); err != nil {
	//	return err
	//}
	err := json.Unmarshal(c.Request.Body(), dto)
	//err := decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body should  only contain a single object")
	}
	return nil
}

// ------------------ svc logic
func (svc *_service) writeBadRequest(c *app.RequestContext, err error) {
	log.Println(err)
	c.AbortWithStatus(http.StatusBadRequest)
}

func (svc *_service) healthcheck(ctx context.Context, c *app.RequestContext) {
	if string(c.Method()) != http.MethodGet {
		c.AbortWithMsg(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	status := map[string]string{
		"status":      "available",
		"environment": svc.settings.env,
	}

	//err := svc.writeJSON(w, http.StatusOK, status)
	c.JSON(http.StatusOK, status)
	//if err != nil {
	//	svc.logger.Print(err.Error())
	//	http.Error(w, "", http.StatusInternalServerError)
	//	return
	//}
}

func (svc *_service) getOrCreateBooks(ctx context.Context, c *app.RequestContext) {
	if string(c.Method()) == http.MethodGet {
		books, err := svc.repository.FindAll(true)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		//if err := svc.writeJSON(w, http.StatusOK, envelope{"books": books}); err != nil {
		//	http.Error(w, "", http.StatusInternalServerError)
		//}
		c.JSON(http.StatusOK, envelope{"books": books})
		return
	}
	if string(c.Method()) == http.MethodPost {
		var bookDto struct {
			Title     string   `json:"title"`
			Published int      `json:"published"`
			Pages     int      `json:"pages"`
			Genres    []string `json:"genres"`
			Rating    float64  `json:"rating"`
		}

		if err := svc.readJSONObject(c, &bookDto); err != nil {
			// custom error handling details: Alex Edwards, Let's Go Further Chapter 4
			svc.writeBadRequest(c, err)
			return
		}

		book, err := svc.repository.
			InsertOne(bookDto.Title, bookDto.Published, bookDto.Pages, bookDto.Rating, bookDto.Genres)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		if err := svc.writeJSON(c, http.StatusCreated, envelope{"book": book}); err != nil {
			//http.Error(w, "", http.StatusInternalServerError)
			c.AbortWithStatus(http.StatusBadRequest)
		}
		return
	}
	//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	c.AbortWithMsg(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (svc *_service) getUpdateOrDeleteBooks(ctx context.Context, c *app.RequestContext) {

	getId := func(c *app.RequestContext) (int64, error) {
		id, err := strconv.ParseInt(c.Request.URI().String()[len("api/v1/books/"):], 10, 64)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		return id, err
	}

	switch string(c.Request.Method()) {
	case http.MethodGet:
		if id, err := getId(c); err == nil {
			svc.getBook(id, c)
		}
	case http.MethodPut:
		if id, err := getId(c); err == nil {
			svc.updateBook(id, c)
		}
	case http.MethodDelete:
		if id, err := getId(c); err == nil {
			svc.deleteBook(id, c)
		}
	default:
		c.AbortWithMsg(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (svc *_service) getBook(id int64, c *app.RequestContext) {
	book, err := svc.repository.FindById(id, true)
	if err != nil {
		switch {
		case errors.Is(err, data.NotFoundError):
			//http.Error(w, "", http.StatusNotFound)
			c.AbortWithStatus(http.StatusNotFound)
			return
		default:
			//http.Error(w, "", http.StatusInternalServerError)
			c.AbortWithStatus(http.StatusInternalServerError)
			return

		}
	}

	if err := svc.writeJSON(c, http.StatusOK, envelope{"book": book}); err != nil {
		//http.Error(w, "", http.StatusInternalServerError)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
func (svc *_service) updateBook(id int64, c *app.RequestContext) {
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
			//http.Error(w, "", http.StatusNotFound)
			c.AbortWithStatus(http.StatusNotFound)
		default:
			//http.Error(w, "", http.StatusInternalServerError)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	if err := svc.readJSONObject(c, &bookDto); err != nil {
		svc.writeBadRequest(c, err)
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
	}
	if err := svc.writeJSON(c, http.StatusOK, envelope{"book": book}); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
func (svc *_service) deleteBook(id int64, c *app.RequestContext) {
	err := svc.repository.DeleteById(id)
	if err != nil {
		switch {
		case errors.Is(err, data.NotFoundError):
			c.AbortWithStatus(http.StatusNotFound)
		default:
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	c.SetStatusCode(http.StatusNoContent)
}
