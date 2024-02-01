package data

type Repository interface {
	Migrate() error
	Ping() error
	InsertOne(title string, published int, pages int, rating float64, genres []string) (*Book, error)
	FindById(id int64, includeGenres bool) (*Book, error)
	FindAll(includeGenres bool) ([]*Book, error)
	LoadGenres(book *Book) ([]string, error)
	Update(book *Book) error
	DeleteById(id int64) error
	Close() error
}
