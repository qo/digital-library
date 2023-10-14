package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/qo/digital-library/internal/config"
)

func Open(options config.SQLiteOptions) (*sql.DB, error) {
	const errMsg = "can't open sqlite db"

	foreignKeys := ""
	if options.ForeignKeys {
		foreignKeys = "on"
	} else {
		foreignKeys = "off"
	}

	optionsString := fmt.Sprintf("file:%s?_foreign_keys=%s", options.Path, foreignKeys)

	db, err := sql.Open("sqlite3", optionsString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return db, nil
}
