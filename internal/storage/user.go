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
    INSERT INTO users(id, first_name, second_name, role)
    VALUES (?, ?, ?, ?);
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
