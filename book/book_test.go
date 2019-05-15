package book

import (
	"errors"
	"testing"
)

func TestBookHasExpectAttribute(t *testing.T) {
	b := Book{
		ID:          "id",
		Title:       "title",
		Description: "description",
		ImageURL:    "image_url",
		Author:      "author",
	}

	if result, _ := getType(b.ID); result != "string" {
		t.Errorf("ID is not string")
	}

	if result, _ := getType(b.Title); result != "string" {
		t.Errorf("Title is not string")
	}

	if result, _ := getType(b.Description); result != "string" {
		t.Errorf("Description is not string")
	}

	if result, _ := getType(b.ImageURL); result != "string" {
		t.Errorf("ImageURL is not string")
	}

	if result, _ := getType(b.Author); result != "string" {
		t.Errorf("Author is not string")
	}
}

func getType(i interface{}) (string, error) {
	result := ""
	switch i.(type) {
	case string:
		result = "string"
	default:
		return result, errors.New("Type not found")
	}
	return result, nil
}
