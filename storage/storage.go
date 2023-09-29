package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func open(path string) (*Storage, error) {
	const errMsg = "can't open db"

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return &Storage{db}, nil
}

func Init(path string) (*Storage, error) {
	const errMsg = "can't init storage"

	st, err := open(path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	err = st.initUsers()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return st, err
}
