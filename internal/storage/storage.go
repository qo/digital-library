package storage

import (
	"database/sql"
	"fmt"

	"github.com/qo/digital-library/internal/config"
	"github.com/qo/digital-library/internal/storage/author"
	"github.com/qo/digital-library/internal/storage/authorship"
	"github.com/qo/digital-library/internal/storage/book"
	"github.com/qo/digital-library/internal/storage/book_review"
	"github.com/qo/digital-library/internal/storage/favorite_author"
	"github.com/qo/digital-library/internal/storage/favorite_book"
	"github.com/qo/digital-library/internal/storage/mysql"
	"github.com/qo/digital-library/internal/storage/sqlite"
	"github.com/qo/digital-library/internal/storage/user"
)

type Storage struct {
	db *sql.DB
}

const (
	mysqlDb  = "mysql"
	sqliteDb = "sqlite"
)

func Init(options config.StorageOptions) (*Storage, error) {
	const errMsg = "can't init storage"

	var (
		db  *sql.DB
		err error
	)

	switch options.Db {
	case mysqlDb:
		db, err = mysql.Open(options.MySQLOptions)
	case sqliteDb:
		db, err = sqlite.Open(options.SQLiteOptions)
	default:
		err = fmt.Errorf("db option %s is unknown", options.Db)
	}

	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	st := Storage{db}

	err = st.initTables()
	if err != nil {
		return nil, err
	}

	return &st, err
}

func (st *Storage) initTables() error {
	db := st.db

	var err error

	err = author.InitTable(db)
	if err != nil {
		return fmt.Errorf("can't init author: %w", err)
	}

	err = authorship.InitTable(db)
	if err != nil {
		return fmt.Errorf("can't init authorship: %w", err)
	}

	err = book.InitTable(db)
	if err != nil {
		return fmt.Errorf("can't init book: %w", err)
	}

	err = book_review.InitTable(db)
	if err != nil {
		return fmt.Errorf("can't init book_review: %w", err)
	}

	err = favorite_author.InitTable(db)
	if err != nil {
		return fmt.Errorf("can't init favorite_author: %w", err)
	}

	err = favorite_book.InitTable(db)
	if err != nil {
		return fmt.Errorf("can't init favorite_book: %w", err)
	}

	err = user.InitTable(db)
	if err != nil {
		return fmt.Errorf("can't init user: %w", err)
	}

	return nil
}

func (s *Storage) GetAuthor(id int) (*author.Author, error) {
	return author.GetAuthor(s.db, id)
}

func (s *Storage) PostAuthor(a *author.Author) error {
	return author.PostAuthor(s.db, a)
}

func (s *Storage) PutAuthor(a *author.Author) error {
	return author.PutAuthor(s.db, a)
}

func (s *Storage) DeleteAuthor(id int) error {
	return author.DeleteAuthor(s.db, id)
}

func (s *Storage) GetBook(id int) (*book.Book, error) {
	return book.GetBook(s.db, id)
}

func (s *Storage) PostBook(a *book.Book) error {
	return book.PostBook(s.db, a)
}

func (s *Storage) PutBook(a *book.Book) error {
	return book.PutBook(s.db, a)
}

func (s *Storage) DeleteBook(id int) error {
	return book.DeleteBook(s.db, id)
}

func (s *Storage) GetUser(id int) (*user.User, error) {
	return user.GetUser(s.db, id)
}

func (s *Storage) PostUser(u *user.User) error {
	return user.PostUser(s.db, u)
}

func (s *Storage) PutUser(u *user.User) error {
	return user.PutUser(s.db, u)
}

func (s *Storage) DeleteUser(id int) error {
	return user.DeleteUser(s.db, id)
}

func (s *Storage) GetBookReviews(id int) ([]book_review.BookReview, error) {
	return user.GetBookReviews(s.db, id)
}

func (s *Storage) GetFavoriteAuthors(id int) ([]author.Author, error) {
	return user.GetFavoriteAuthors(s.db, id)
}

func (s *Storage) GetFavoriteBooks(id int) ([]book.Book, error) {
	return user.GetFavoriteBooks(s.db, id)
}
