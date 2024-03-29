package favorite_author

import (
	"database/sql"
	"fmt"
)

type FavoriteAuthor struct {
	UserId   int `json:"user_id"`
	AuthorId int `json:"author_id"`
}

func InitTable(db *sql.DB) error {
	const errMsg = "can't init favorite authors table"

	stmt, err := db.Prepare(`
    CREATE TABLE IF NOT EXISTS favorite_authors(
      user_id INTEGER,
      author_id INTEGER,
      FOREIGN KEY (user_id) REFERENCES users (id),
      FOREIGN KEY (author_id) REFERENCES authors (id),
      PRIMARY KEY (user_id, author_id)
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

func GetFavoriteAuthor(db *sql.DB, userId, authorId int) (*FavoriteAuthor, error) {
	const errMsg = "can't get favorite author"

	stmt, err := db.Prepare(`
    SELECT * FROM favorite_authors
    WHERE user_id = ?
    AND author_id = ?
  `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	row := stmt.QueryRow(userId, authorId)

	var favoriteAuthor FavoriteAuthor

	err = row.Scan(&favoriteAuthor.UserId, &favoriteAuthor.AuthorId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return &favoriteAuthor, nil
}

func PutFavoriteAuthor(db *sql.DB, favoriteAuthor *FavoriteAuthor) (int, int, error) {
	const errMsg = "can't put favorite author"

	stmt, err := db.Prepare(`
    INSERT INTO favorite_authors
    (user_id, author_id)
    VALUES
    (?, ?);
  `)
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	stmt.QueryRow(favoriteAuthor.UserId, favoriteAuthor.AuthorId)

	return favoriteAuthor.UserId, favoriteAuthor.AuthorId, nil
}

func DeleteFavoriteAuthor(db *sql.DB, userId, authorId int) (int, int, error) {
	const errMsg = "can't delete favorite author"

	stmt, err := db.Prepare(`
    DELETE FROM favorite_authors
    WHERE user_id = ?
    AND author_id = ?
  `)
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	stmt.QueryRow(userId, authorId)

	return userId, authorId, nil
}
