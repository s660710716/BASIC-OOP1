package memory

import "book-api/internal/domain"

type BookRepo struct {
	data []domain.Book
	nextID int
}

func NewBookRepo() *BookRepo {
	return &BookRepo{nextID: 1}
}

func (r *BookRepo) ExistsByTitle(title string) (bool, error) {
	for _, b := range r.data {
	if b.Title == title {
		return true, nil
		}
	}
	return false, nil
}

func (r *BookRepo) Create(book *domain.Book) error {
	book.ID = r.nextID
	r.nextID++
	r.data = append(r.data, *book)
	return nil
}