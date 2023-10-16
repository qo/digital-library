package storage

import (
	"fmt"
)

type BookReview struct {
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
	Rating int `json:"rating"`
}

func (s *Storage) initBookReviews() error {
	const errMsg = "can't init book reviews table"

	stmt, err := s.db.Prepare(`
    CREATE TABLE IF NOT EXISTS book_reviews(
      user_id INTEGER,
      book_id INTEGER,
      rating INTEGER,
      FOREIGN KEY (user_id) REFERENCES users (id),
      FOREIGN KEY (book_id) REFERENCES books (id),
      PRIMARY KEY (user_id, book_id)
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

func (s *Storage) GetBookReview(userId, bookId int) (*BookReview, error) {
	const errMsg = "can't get book review"

	stmt, err := s.db.Prepare(`
    SELECT * FROM book_reviews
    WHERE user_id = ?
    AND book_id = ?
  `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	row := stmt.QueryRow(userId, bookId)

	var bookReview BookReview

	err = row.Scan(&bookReview.UserId, &bookReview.BookId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return &bookReview, nil
}

func (s *Storage) PutBookReview(bookReview *BookReview) (int, int, error) {
	const errMsg = "can't put book review"

	stmt, err := s.db.Prepare(`
    INSERT INTO book_reviews
    (user_id, book_id)
    VALUES
    (?, ?);
  `)
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	stmt.QueryRow(bookReview.UserId, bookReview.BookId)

	return bookReview.UserId, bookReview.BookId, nil
}

func (s *Storage) DeleteBookReview(userId, bookId int) (int, int, error) {
	const errMsg = "can't delete book review"

	stmt, err := s.db.Prepare(`
    DELETE FROM book_reviews
    WHERE user_id = ?
    AND book_id = ?
  `)
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	stmt.QueryRow(userId, bookId)

	return userId, bookId, nil
}
