package main

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
