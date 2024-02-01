package data

import "time"

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Published int       `json:"published,omitempty"`
	Pages     int       `json:"pages,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Rating    float64   `json:"rating,omitempty"`
	Version   int32     `json:"-"`
}
