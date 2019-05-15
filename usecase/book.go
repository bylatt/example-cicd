package usecase

import "github.com/clozed2u/example-cicd/book"

// BookUsecase ...
type BookUsecase struct {
	BookRepo book.Repo
}

// Find ...
func (bu BookUsecase) Find(id string) (book.Book, error) {
	return bu.BookRepo.Find(id)
}

// Create ...
func (bu BookUsecase) Create(b book.Book) error {
	return bu.BookRepo.Create(b)
}

// Delete ...
func (bu BookUsecase) Delete(id string) error {
	return bu.BookRepo.Delete(id)
}

// Update ...
func (bu BookUsecase) Update(id string, b book.Book) error {
	return bu.BookRepo.Update(id, b)
}

// NewBookUsecase ...
func NewBookUsecase(b book.Repo) BookUsecase {
	return BookUsecase{BookRepo: b}
}
