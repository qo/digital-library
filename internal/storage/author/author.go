package author

import (
	"database/sql"
	"fmt"
)

type Author struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
}

func InitTable(db *sql.DB) error {
	const errMsg = "can't init authors table"

	stmt, err := db.Prepare(`
    CREATE TABLE IF NOT EXISTS authors(
      id INTEGER PRIMARY KEY,
      full_name TEXT
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

func GetAuthor(db *sql.DB, id int) (*Author, error) {
	const errMsg = "can't get author"

	stmt, err := db.Prepare(`
    SELECT * FROM authors
    WHERE id = ?;
  `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	row := stmt.QueryRow(id)

	var author Author

	err = row.Scan(&author.Id, &author.FullName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return &author, nil
}

func PostAuthor(db *sql.DB, author *Author) error {
	const errMsg = "can't post author"

	stmt, err := db.Prepare(`
    INSERT INTO books
    (id, full_name)
    VALUES 
    (?, ?);
  `)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	_, err = stmt.Exec(author.Id, author.FullName)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	return nil
}

func PutAuthor(db *sql.DB, author *Author) error {
	const errMsg = "can't put author"

	stmt, err := db.Prepare(`
    UPDATE authors
    SET full_name = ?
    WHERE id = ?;
  `)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	_, err = stmt.Exec(author.FullName, author.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	return nil
}

func DeleteAuthor(db *sql.DB, id int) error {
	const errMsg = "can't delete author"

	stmt, err := db.Prepare(`
    DELETE FROM authors
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
