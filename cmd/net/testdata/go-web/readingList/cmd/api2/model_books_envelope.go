package main

import "github.com/hertz-contrib/migrate/cmd/net/testdata/go-web/readingList/internal/data"

type BooksEnvelope struct {
	Books []*Book `json:"books,omitempty"`
}

func NewBooksEnvelope(books []*data.Book) *BooksEnvelope {
	booksModel := NewBooksModel(books)
	return &BooksEnvelope{Books: booksModel}
}
