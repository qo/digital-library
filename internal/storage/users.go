package storage

import (
	"fmt"
)

type User struct {
	Id         int
	FirstName  string
	SecondName string
}

func (s *Storage) initUsers() error {
	const errMsg = "can't init users table"

	stmt, err := s.db.Prepare(`
    CREATE TABLE IF NOT EXISTS users(
      id INTEGER PRIMARY KEY,
      first_name TEXT,
      second_name TEXT
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

	err = row.Scan(&user.Id, &user.FirstName, &user.SecondName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return &user, nil
}

func (s *Storage) PutUser(user *User) (int, error) {
	const errMsg = "can't put user"

	stmt, err := s.db.Prepare(`
    INSERT INTO users
    (id, first_name, second_name)
    VALUES
    (?, ?, ?);
  `)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	stmt.QueryRow(user.Id, user.FirstName, user.SecondName)

	return user.Id, nil
}

func (s *Storage) DeleteUser(id int) (int, error) {
	const errMsg = "can't delete user"

	stmt, err := s.db.Prepare(`
    DELETE FROM users
    WHERE id = ?;
  `)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", errMsg, err)
	}

	stmt.QueryRow(id)

	return id, nil
}
