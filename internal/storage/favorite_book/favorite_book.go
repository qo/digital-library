package favorite_book

import (
	"database/sql"
	"fmt"
)

type FavoriteBook struct {
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
}

func InitTable(db *sql.DB) error {
	const errMsg = "can't init favorite books table"

	stmt, err := db.Prepare(`
    CREATE TABLE IF NOT EXISTS favorite_books(
      user_id INTEGER,
      book_id INTEGER,
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

func GetFavoriteBook(db *sql.DB, userId, bookId int) (*FavoriteBook, error) {
	const errMsg = "can't get favorite book"

	stmt, err := db.Prepare(`
    SELECT * FROM favorite_books
    WHERE user_id = ?
    AND book_id = ?
  `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	row := stmt.QueryRow(userId, bookId)

	var favoriteBook FavoriteBook

	err = row.Scan(&favoriteBook.UserId, &favoriteBook.BookId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return &favoriteBook, nil
}

func PutFavoriteBook(db *sql.DB, favoriteBook *FavoriteBook) (int, int, error) {
	const errMsg = "can't put favorite book"

	stmt, err := db.Prepare(`
    INSERT INTO favorite_books
    (user_id, book_id)
    VALUES
    (?, ?);
  `)
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	stmt.QueryRow(favoriteBook.UserId, favoriteBook.BookId)

	return favoriteBook.UserId, favoriteBook.BookId, nil
}

func DeleteFavoriteBook(db *sql.DB, userId, bookId int) (int, int, error) {
	const errMsg = "can't delete favorite book"

	stmt, err := db.Prepare(`
    DELETE FROM favorite_books
    WHERE user_id = ?
    AND book_id = ?
  `)
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	stmt.QueryRow(userId, bookId)

	return userId, bookId, nil
}
