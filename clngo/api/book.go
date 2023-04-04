package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

var books = map[string]Book{
	"666": {Title: "The Hitchhiker's Gide to the Galaxy", Author: "Douglas Adams", ISBN: "666", Description: "Awesome!"},
	"777": {Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "777"},
	"999": {Title: "The Missing Readme", Author: "Dmitriy Ryaboy", ISBN: "999"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	// implement logic for /api/books
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		out_json, err := json.Marshal(books)
		if err != nil {
			panic(err)
		}
		w.Header().Add("Content-Type", "application/json; charset-utf-8")
		w.Write(out_json)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}

}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			json_out := ToJSON(book)
			w.Header().Add("Content-Type", "application/json; charset-utf-8")
			w.Write(json_out)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))

	}

}

func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}

func CreateBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]
	if exists {
		return "", false
	}
	books[book.ISBN] = book

	return book.ISBN, true
}

func GetBook(isbn string) (Book, bool) {
	book, exists := books[isbn]

	return book, exists
}

func UpdateBook(isbn string, book Book) bool {
	_, exists := books[isbn]
	if exists {
		books[isbn] = book
		return exists
	}

	return false
}

func DeleteBook(isbn string) {
	delete(books, isbn)
}

func ToJSON(b Book) []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return ToJSON
}

func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}

	return book
}
