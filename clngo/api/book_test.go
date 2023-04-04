package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	json := ToJSON(Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "12345"})

	assert.Equal(t, `{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"12345"}`,
		string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"12345"}`)

	assert.Equal(t, Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "12345"},
		FromJSON(json), "Book JSON unmarshalling wrong.")
}
