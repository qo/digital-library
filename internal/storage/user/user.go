package user

import (
	"database/sql"
	"fmt"

	"github.com/qo/digital-library/internal/storage/author"
	"github.com/qo/digital-library/internal/storage/book"
	"github.com/qo/digital-library/internal/storage/book_review"
)

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Role       int    `json:"role"` // 1 - user, 2 - mod, 3 - admin
}

func InitTable(db *sql.DB) error {
	const errMsg = "can't init users table"

	stmt, err := db.Prepare(`
    CREATE TABLE IF NOT EXISTS users(
      id INTEGER PRIMARY KEY,
      first_name TEXT,
      second_name TEXT,
      role INTEGER
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

func PostUser(db *sql.DB, user *User) error {
	const errMsg = "can't post user"

	stmt, err := db.Prepare(`
    INSERT INTO users
    (id, first_name, second_name, role)
    VALUES 
    (?, ?, ?, ?);
  `)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	_, err = stmt.Exec(user.Id, user.FirstName, user.SecondName, user.Role)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	return nil
}

func GetUser(db *sql.DB, id int) (*User, error) {
	const errMsg = "can't get user"

	stmt, err := db.Prepare(`
    SELECT * FROM users
    WHERE id = ?;
  `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	row := stmt.QueryRow(id)

	var user User

	err = row.Scan(&user.Id, &user.FirstName, &user.SecondName, &user.Role)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return &user, nil
}

func PutUser(db *sql.DB, user *User) error {
	const errMsg = "can't put user"

	stmt, err := db.Prepare(`
    UPDATE users
    SET first_name = ?, second_name = ?, role = ?
    WHERE id = ?;
  `)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	_, err = stmt.Exec(user.FirstName, user.SecondName, user.Role, user.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	return nil
}

func DeleteUser(db *sql.DB, id int) error {
	const errMsg = "can't delete user"

	stmt, err := db.Prepare(`
    DELETE FROM users
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

func GetFavoriteBooks(db *sql.DB, id int) ([]book.Book, error) {
	const errMsg = "can't get favorite books"

	stmt, err := db.Prepare(`
    SELECT b.* FROM favorite_books AS fb
    JOIN books AS b
    ON fb.book_id = b.id
    WHERE fb.user_id = ?;
  `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	books := make([]book.Book, 0)

	for rows.Next() {
		var book book.Book
		err := rows.Scan(&book.Id, &book.Isbn, &book.Title, &book.Year, &book.Publisher)
		if err != nil {
			return nil, fmt.Errorf("%s: can't scan book: %s", errMsg, err)
		}
		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%s: error occured while iterating over books: %s", errMsg, err)
	}

	return books, nil
}

func GetFavoriteAuthors(db *sql.DB, id int) ([]author.Author, error) {
	const errMsg = "can't get favorite authors"

	stmt, err := db.Prepare(`
    SELECT * FROM favorite_authors
    JOIN authors
    ON favorite_authors.author_id = authors.id
    WHERE favorite_books.user_id = ?;
  `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	authors := make([]author.Author, 0)

	for rows.Next() {
		var author author.Author
		err := rows.Scan(&author.Id, &author.FullName)
		if err != nil {
			return nil, fmt.Errorf("%s: can't scan author: %s", errMsg, err)
		}
		authors = append(authors, author)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%s: error occured while iterating over authors: %s", errMsg, err)
	}

	return authors, nil
}

func GetBookReviews(db *sql.DB, id int) ([]book_review.BookReview, error) {
	const errMsg = "can't get book reviews"

	stmt, err := db.Prepare(`
    SELECT * FROM book_reviews
    WHERE user_id = ?;
  `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	reviews := make([]book_review.BookReview, 0)

	for rows.Next() {
		var review book_review.BookReview
		err := rows.Scan(&review.UserId, &review.BookId, &review.Rating)
		if err != nil {
			return nil, fmt.Errorf("%s: can't scan book review: %s", errMsg, err)
		}
		reviews = append(reviews, review)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%s: error occured while iterating over book reviews: %s", errMsg, err)
	}

	return reviews, nil
}
