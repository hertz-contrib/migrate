package main

import "github.com/hertz-contrib/migrate/cmd/net/testdata/go-web/readingList/internal/data"

func getBooks() []data.Book {
	return []data.Book{
		{
			Title:     "The Andromeda Strain",
			Pages:     350,
			Published: 1969,
			Rating:    4.0,
			Genres:    []string{"Techno-Thriller"},
		},
		{
			Title:     "The Terminal Man",
			Pages:     247,
			Published: 1972,
			Rating:    4.1,
			Genres:    []string{"Science Fiction"},
		},
		{
			Title:     "The Great Train Robbery",
			Pages:     266,
			Published: 1975,
			Rating:    4.2,
			Genres:    []string{"Historical", "Crime"},
		},
		{
			Title:     "Eaters of the Dead",
			Pages:     288,
			Published: 1976,
			Rating:    4.3,
			Genres:    []string{"Historical"},
		},
		{
			Title:     "Congo",
			Pages:     348,
			Published: 1980,
			Rating:    4.4,
			Genres:    []string{"Science Fiction", "Adventure"},
		},
		{
			Title:     "Sphere",
			Pages:     385,
			Published: 1987,
			Rating:    4.5,
			Genres:    []string{"Science Fiction"},
		},
		{
			Title:     "Jurassic Park",
			Pages:     448,
			Published: 1990,
			Rating:    4.6,
			Genres:    []string{"Science Fiction", "Action"},
		},
		{
			Title:     "Rising Sun",
			Pages:     385,
			Published: 1992,
			Rating:    4.7,
			Genres:    []string{"Crime", "Thriller"},
		},
		{
			Title:     "Disclosure",
			Pages:     597,
			Published: 1994,
			Rating:    4.8,
			Genres:    []string{"Crime"},
		},
		{
			Title:     "The Lost World",
			Pages:     393,
			Published: 1995,
			Rating:    4.9,
			Genres:    []string{"Science Fiction", "Action"},
		},
		{
			Title:     "Airframe",
			Pages:     352,
			Published: 1996,
			Rating:    5.0,
			Genres:    []string{"Techno-Thriller"},
		},
		{
			Title:     "Timeline",
			Pages:     464,
			Published: 1999,
			Rating:    4.1,
			Genres:    []string{"Science Fiction", "Historical", "Time Travel"},
		},
		{
			Title:     "Prey",
			Pages:     502,
			Published: 2002,
			Rating:    4.2,
			Genres:    []string{"Science Fiction", "Techno-Thriller", "Horror", "Nanopunk"},
		},
		{
			Title:     "State of Fear",
			Pages:     641,
			Published: 2004,
			Rating:    4.3,
			Genres:    []string{"Science Fiction", "Techno-Thriller", "Dystopian"},
		},
		{
			Title:     "Next",
			Pages:     528,
			Published: 2006,
			Rating:    4.4,
			Genres:    []string{"Science Fiction", "Techno-Thriller", "Satire"},
		},
		{
			Title:     "Pirate Latitudes",
			Pages:     313,
			Published: 2009,
			Rating:    4.5,
			Genres:    []string{"Historical", "Adventure"},
		},
		{
			Title:     "Micro",
			Pages:     424,
			Published: 2011,
			Rating:    4.6,
			Genres:    []string{"Science Fiction", "Techno-Thriller", "Adventure"},
		},
		{
			Title:     "Dragon Teeth",
			Pages:     320,
			Published: 2017,
			Rating:    4.7,
			Genres:    []string{"Historical", "Adventure"},
		},
	}
}
