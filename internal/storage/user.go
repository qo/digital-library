package storage

import (
	"fmt"
)

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Role       int    `json:"role"` // 1 - user, 2 - mod, 3 - admin
}

func (s *Storage) initUsers() error {
	const errMsg = "can't init users table"

	stmt, err := s.db.Prepare(`
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

func (s *Storage) PostUser(user *User) error {
	const errMsg = "can't post user"

	stmt, err := s.db.Prepare(`
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

func (s *Storage) GetUser(id int) (*User, error) {
	const errMsg = "can't get user"

	stmt, err := s.db.Prepare(`
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

func (s *Storage) PutUser(user *User) error {
	const errMsg = "can't put user"

	stmt, err := s.db.Prepare(`
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

func (s *Storage) DeleteUser(id int) error {
	const errMsg = "can't delete user"

	stmt, err := s.db.Prepare(`
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

func (s *Storage) GetFavoriteBooks(id int) ([]Book, error) {
	const errMsg = "can't get favorite books"

	stmt, err := s.db.Prepare(`
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

	books := make([]Book, 0)

	for rows.Next() {
		var book Book
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

func (s *Storage) GetFavoriteAuthors(id int) ([]Author, error) {
	const errMsg = "can't get favorite authors"

	stmt, err := s.db.Prepare(`
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

	authors := make([]Author, 0)

	for rows.Next() {
		var author Author
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

func (s *Storage) GetBookReviews(id int) ([]BookReview, error) {
	const errMsg = "can't get book reviews"

	stmt, err := s.db.Prepare(`
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

	reviews := make([]BookReview, 0)

	for rows.Next() {
		var review BookReview
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
