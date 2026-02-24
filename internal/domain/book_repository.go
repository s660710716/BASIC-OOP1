package domain

type BookRepository interface {
	ExistsByTitle(title string) (bool, error)
	Create(book *Book) error
}