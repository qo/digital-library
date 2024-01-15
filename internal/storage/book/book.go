package book

import (
	"database/sql"
	"fmt"
)

type Book struct {
	Id        int    `json:"id"`
	Isbn      string `json:"isbn"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Publisher string `json:"publisher"`
}

func InitTable(db *sql.DB) error {
	const errMsg = "can't init books table"

	stmt, err := db.Prepare(`
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

func GetBook(db *sql.DB, id int) (*Book, error) {
	const errMsg = "can't get book"

	stmt, err := db.Prepare(`
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

func PostBook(db *sql.DB, book *Book) error {
	const errMsg = "can't post book"

	stmt, err := db.Prepare(`
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

func PutBook(db *sql.DB, book *Book) error {
	const errMsg = "can't put book"

	stmt, err := db.Prepare(`
    UPDATE books
    SET isbn = ?, title = ?, year = ?, publisher = ?
    WHERE id = ?;
  `)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	_, err = stmt.Exec(book.Isbn, book.Title, book.Year, book.Publisher, book.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	return nil
}

func DeleteBook(db *sql.DB, id int) error {
	const errMsg = "can't delete book"

	stmt, err := db.Prepare(`
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
