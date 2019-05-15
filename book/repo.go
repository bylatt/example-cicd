package book

// RepoAdapter ...
type RepoAdapter interface {
	Find(id string) (Book, error)
	Create(b Book) error
	Delete(id string) error
	Update(id string, b Book) error
}

// Repo ...
type Repo struct {
	Adapter RepoAdapter
}

// Find ...
func (r Repo) Find(id string) (Book, error) {
	return r.Adapter.Find(id)
}

// Create ...
func (r Repo) Create(b Book) error {
	return r.Adapter.Create(b)
}

// Delete ...
func (r Repo) Delete(id string) error {
	return r.Adapter.Delete(id)
}

// Update ...
func (r Repo) Update(id string, b Book) error {
	return r.Adapter.Update(id, b)
}

// NewRepo ...
func NewRepo(r RepoAdapter) Repo {
	return Repo{Adapter: r}
}
