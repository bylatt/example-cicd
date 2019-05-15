package book

import (
	"reflect"
	"testing"
)

type MockAdapter struct{}

func (m MockAdapter) Find(id string) (Book, error) {
	b := Book{
		ID:          "id",
		Title:       "title",
		Description: "description",
		ImageURL:    "image_url",
		Author:      "author",
	}
	return b, nil
}

func (m MockAdapter) Create(b Book) error {
	return nil
}

func (m MockAdapter) Delete(id string) error {
	return nil
}

func (m MockAdapter) Update(id string, b Book) error {
	return nil
}

func TestRepoFind(t *testing.T) {
	r := NewRepo(MockAdapter{})
	want := Book{
		ID:          "id",
		Title:       "title",
		Description: "description",
		ImageURL:    "image_url",
		Author:      "author",
	}
	got, _ := r.Find("id")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want %+v but got %+v\n", want, got)
	}
}
