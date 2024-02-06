package todoapp

import (
	"fmt"
	"net/http"
	"time"
)

var SavedEntries []*TodoEntry

func _raiseError(response http.ResponseWriter, err error) {
	http.Error(response, err.Error(), http.StatusInternalServerError)
}

func IndexView(response http.ResponseWriter, request *http.Request) {
	context := indexContextData{Entries: SavedEntries}
	err := renderTemplate(response, "index.html", context)

	if err != nil {
		_raiseError(response, err)
	}
}

func _handleCreateEntry(request *http.Request) error {
	var err error

	err = request.ParseForm()

	if err != nil {
		return err
	}

	title := request.PostForm.Get("title")
	text := request.PostForm.Get("text")

	if title == "" || text == "" {
		err = fmt.Errorf("Please fill-up everything.") // noqa
	} else {
		SavedEntries = append(SavedEntries, &TodoEntry{
			Title: title,
			Text:  text,
			Date:  time.Now()})
	}

	return err
}

func CreateEntry(response http.ResponseWriter, request *http.Request) {
	var err error
	isPostRequest := request.Method == "POST"

	if isPostRequest {
		err = _handleCreateEntry(request)
	}

	// If the user did not submit or data was invalid, render the form
	if !isPostRequest || err != nil {
		context := formContextData{request.PostForm, err}
		err = renderTemplate(response, "form.html", context)

		if err != nil {
			_raiseError(response, err)
		}
	} else {
		// The form was successfully submitted,
		// redirect the user to the index
		http.Redirect(response, request, "/", http.StatusFound)
	}
}
