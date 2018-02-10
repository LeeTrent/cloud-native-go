package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "ML Reimer", ISBN: "0123456789"}
	json := book.ToJSON()
	assert.Equal(t,
		`{"title":"Cloud Native Go","author":"ML Reimer","isbn":"0123456789"}`,
		string(json),
		"Book to JSON marshalling is incorrect.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"ML Reimer","isbn":"0123456789"}`)
	book := FromJSON(json)
	assert.Equal(t,
		Book{Title: "Cloud Native Go", Author: "ML Reimer", ISBN: "0123456789"},
		book,
		"Book unmarshalling to JSON is incorrect.")
}
