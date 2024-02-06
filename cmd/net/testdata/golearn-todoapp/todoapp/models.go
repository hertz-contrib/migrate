package todoapp

import (
	"net/url"
	"time"
)

// Defines a to do entry
type TodoEntry struct {
	Title string
	Text  string
	Date  time.Time
}

// Defines a template context for an index of “to do“ entries
type indexContextData struct {
	Entries []*TodoEntry
}

// Defines a template context for forms
type formContextData struct {
	Form  url.Values
	Error error
}
