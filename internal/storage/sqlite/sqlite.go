package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/qo/digital-library/internal/config"
)

func Open(options config.SQLiteOptions) (*sql.DB, error) {
	const errMsg = "can't open sqlite db"

	db, err := sql.Open("sqlite3", options.Path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return db, nil
}
