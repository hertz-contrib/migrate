package main

import (
	"github.com/hertz-contrib/migrate/cmd/net/testdata/go-web/readingList/internal/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	books, err := app.client.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", books)
	if err != nil {
		app.logger.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) bookView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	book, err := app.client.Get(int64(id))
	if err != nil {
		app.logger.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/view.html",
	}
	functions := template.FuncMap{"join": strings.Join}

	ts, err := template.New("showBook").Funcs(functions).ParseFiles(files...)
	if err != nil {
		app.logger.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", book)
	if err != nil {
		app.logger.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) bookCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.bookCreateForm(w, r)
	case http.MethodPost:
		app.bookCreateProcess(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) bookCreateForm(w http.ResponseWriter, r *http.Request) {
	_ = r

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/create.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) bookCreateProcess(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.logger.Print(err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	pages, err := strconv.Atoi(r.PostForm.Get("pages"))
	if err != nil || pages < 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	published, err := strconv.Atoi(r.PostForm.Get("pages"))
	if err != nil || published < 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	genres := strings.Split(r.PostForm.Get("genres"), " ")
	rating, err := strconv.ParseFloat(r.PostForm.Get("rating"), 64)
	if err != nil || rating < 0.0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	addOrUpdateBook := &models.AddOrUpdateBook{
		Title:     title,
		Pages:     pages,
		Published: published,
		Genres:    genres,
		Rating:    rating,
	}

	if _, err = app.client.Create(addOrUpdateBook); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) bookEdit(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.bookEditForm(w, r)
	case http.MethodPost:
		app.bookEditProcess(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) bookEditForm(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r
}

func (app *application) bookEditProcess(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r
}
