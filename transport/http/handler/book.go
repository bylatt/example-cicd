package handler

import (
	"net/http"

	"github.com/clozed2u/example-cicd/book"
	"github.com/clozed2u/example-cicd/usecase"
	"github.com/kjk/betterguid"
	"github.com/labstack/echo"
)

// BookHandler ...
type BookHandler struct {
	BookUsecase usecase.BookUsecase
}

// FindBook ...
func (bh BookHandler) FindBook(c echo.Context) error {
	id := c.Param("id")
	b, err := bh.BookUsecase.Find(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Something went wrong"})
	}
	return c.JSON(http.StatusOK, Response{Message: "ok", Data: b})
}

// CreateBook ...
func (bh BookHandler) CreateBook(c echo.Context) error {
	b := new(book.Book)
	err := c.Bind(b)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to create new book"})
	}
	b.ID = betterguid.New()
	actualBook := book.Book{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		ImageURL:    b.ImageURL,
		Author:      b.Author,
	}
	err = bh.BookUsecase.Create(actualBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to create new book"})
	}
	return c.JSON(http.StatusOK, Response{Message: "ok", Data: actualBook})
}

// DeleteBook ...
func (bh BookHandler) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	err := bh.BookUsecase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to delete book"})
	}
	return c.JSON(http.StatusOK, Response{Message: "ok"})
}

// UpdateBook ...
func (bh BookHandler) UpdateBook(c echo.Context) error {
	id := c.Param("id")
	b := new(book.Book)
	err := c.Bind(b)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to update book"})
	}
	actualBook := book.Book{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		ImageURL:    b.ImageURL,
		Author:      b.Author,
	}
	err = bh.BookUsecase.Update(id, actualBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to update book"})
	}
	return c.JSON(http.StatusOK, Response{Message: "ok", Data: actualBook})
}

// NewBookHandler ...
func NewBookHandler(buc usecase.BookUsecase) BookHandler {
	return BookHandler{BookUsecase: buc}
}
