package storage

import (
	"fmt"
)

type Author struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
}

func (s *Storage) initAuthors() error {
	const errMsg = "can't init authors table"

	stmt, err := s.db.Prepare(`
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

func (s *Storage) PostAuthor(author *Author) error {
	const errMsg = "can't post author"

	stmt, err := s.db.Prepare(`
    INSERT INTO books(id, full_name)
    VALUES (?, ?);
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

func (s *Storage) GetAuthor(id int) (*Author, error) {
	const errMsg = "can't get author"

	stmt, err := s.db.Prepare(`
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

func (s *Storage) PutAuthor(author *Author) error {
	const errMsg = "can't put author"

	stmt, err := s.db.Prepare(`
    INSERT INTO authors
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

func (s *Storage) DeleteAuthor(id int) error {
	const errMsg = "can't delete author"

	stmt, err := s.db.Prepare(`
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
