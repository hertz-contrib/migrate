package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Published int      `json:"published"`
	Pages     int      `json:"pages"`
	Genres    []string `json:"genres"`
	Rating    float64  `json:"rating"`
}

type AddOrUpdateBook struct {
	Title     string   `json:"title"`
	Published int      `json:"published"`
	Pages     int      `json:"pages"`
	Genres    []string `json:"genres"`
	Rating    float64  `json:"rating"`
}

type BookResponse struct {
	Book *Book `json:"book"`
}

type BooksResponse struct {
	Books *[]Book `json:"books"`
}

type ReadingListClient struct {
	Endpoint string
}

func (c *ReadingListClient) Create(book *AddOrUpdateBook) (*Book, error) {
	body, err := json.Marshal(*book)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.Endpoint, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected response %d", resp.StatusCode)
	}

	var bookResponse BookResponse
	if err := ReadJSONObject(resp.Body, &bookResponse); err != nil {
		return nil, err
	}
	return bookResponse.Book, nil
}

func (c *ReadingListClient) GetAll() (*[]Book, error) {
	resp, err := http.Get(c.Endpoint)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}
	var booksResponse BooksResponse
	if err := ReadJSONObject(resp.Body, &booksResponse); err != nil {
		return nil, err
	}
	return booksResponse.Books, nil
}

func (c *ReadingListClient) Get(id int64) (*Book, error) {
	url := fmt.Sprintf("%s/%d", c.Endpoint, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var bookResponse BookResponse
	if err := ReadJSONObject(resp.Body, &bookResponse); err != nil {
		return nil, err
	}
	return bookResponse.Book, nil
}
