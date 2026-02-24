package usecase

import (
	"errors"
	"strings"

	"book-api/internal/domain"
)

type CreateBookUsecase struct {
	repo domain.BookRepository
}

func NewCreateBookUsecase(repo domain.BookRepository) *CreateBookUsecase{
	return &CreateBookUsecase{repo:repo}
	
}

func (u *CreateBookUsecase) Execute(title string) (*domain.Book, error) {
	title = strings.TrimSpace(title)
	if title == "" {
	return nil, errors.New("title is required")
}

exists, err := u.repo.ExistsByTitle(title)
	if err != nil {
	return nil, err
}

if exists {
	return nil, errors.New("title already exists")
}

book := &domain.Book{Title: title}
	if err := u.repo.Create(book); err != nil {
	return nil, err
	}
	return book, nil
}