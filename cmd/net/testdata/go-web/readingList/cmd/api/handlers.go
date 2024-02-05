package main

import (
	"errors"
	"github.com/hertz-contrib/migrate/cmd/net/testdata/go-web/readingList/internal/data"
	_ "github.com/hertz-contrib/migrate/cmd/net/testdata/go-web/readingList/internal/data"
	"net/http"
	"strconv"
)

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
