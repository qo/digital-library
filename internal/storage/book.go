package storage

import (
	"fmt"
)

type Book struct {
	Id        int    `json:"id"`
	Isbn      string `json:"isbn"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Publisher string `json:"publisher"`
}

func (s *Storage) initBooks() error {
	const errMsg = "can't init books table"

	stmt, err := s.db.Prepare(`
    CREATE TABLE IF NOT EXISTS books(
      id INTEGER PRIMARY KEY,
      isbn TEXT,
      title TEXT,
      year INTEGER,
      publisher TEXT
    );
  `)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	return nil
}

func (s *Storage) PostBook(book *Book) error {
	const errMsg = "can't post book"

	stmt, err := s.db.Prepare(`
    INSERT INTO books(id, isbn, title, year, publisher)
    VALUES (?, ?, ?, ?, ?);
  `)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	_, err = stmt.Exec(book.Id, book.Isbn, book.Title, book.Year, book.Publisher)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	return nil
}

func (s *Storage) GetBook(id int) (*Book, error) {
	const errMsg = "can't get book"

	stmt, err := s.db.Prepare(`
    SELECT * FROM books
    WHERE id = ?;
  `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	row := stmt.QueryRow(id)

	var book Book

	err = row.Scan(&book.Id, &book.Isbn, &book.Title, &book.Year, &book.Publisher)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return &book, nil
}

func (s *Storage) PutBook(book *Book) error {
	const errMsg = "can't put book"

	stmt, err := s.db.Prepare(`
    INSERT INTO books
    (id, isbn, title, year, publisher)
    VALUES
    (?, ?, ?, ?, ?);
  `)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	_, err = stmt.Exec(book.Id, book.Isbn, book.Title, book.Year, book.Publisher)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	return nil
}

func (s *Storage) DeleteBook(id int) error {
	const errMsg = "can't delete book"

	stmt, err := s.db.Prepare(`
    DELETE FROM books
    WHERE id = ?;
  `)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	return nil
}
