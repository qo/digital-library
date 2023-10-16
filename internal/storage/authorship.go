package storage

import (
	"fmt"
)

type Authorship struct {
	AuthorId int `json:"author_id"`
	BookId   int `json:"book_id"`
}

func (s *Storage) initAuthorships() error {
	const errMsg = "can't init authorships table"

	stmt, err := s.db.Prepare(`
    CREATE TABLE IF NOT EXISTS authorships(
      author_id INTEGER,
      book_id INTEGER,
      FOREIGN KEY (author_id) REFERENCES authors (id),
      FOREIGN KEY (book_id) REFERENCES books (id),
      PRIMARY KEY (author_id, book_id)
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

func (s *Storage) GetAuthorship(authorId, bookId int) (*Authorship, error) {
	const errMsg = "can't get authorship"

	stmt, err := s.db.Prepare(`
    SELECT * FROM authorships
    WHERE author_id = ?
    AND book_id = ?
  `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	row := stmt.QueryRow(authorId, bookId)

	var bookReview Authorship

	err = row.Scan(&bookReview.AuthorId, &bookReview.BookId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return &bookReview, nil
}

func (s *Storage) PutAuthorship(bookReview *Authorship) (int, int, error) {
	const errMsg = "can't put authorship"

	stmt, err := s.db.Prepare(`
    INSERT INTO authorships
    (author_id, book_id)
    VALUES
    (?, ?);
  `)
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	stmt.QueryRow(bookReview.AuthorId, bookReview.BookId)

	return bookReview.AuthorId, bookReview.BookId, nil
}

func (s *Storage) DeleteAuthorship(authorId, bookId int) (int, int, error) {
	const errMsg = "can't delete authorship"

	stmt, err := s.db.Prepare(`
    DELETE FROM authorships
    WHERE author_id = ?
    AND book_id = ?
  `)
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	stmt.QueryRow(authorId, bookId)

	return authorId, bookId, nil
}
