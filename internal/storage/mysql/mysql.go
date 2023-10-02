package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qo/digital-library/internal/config"
)

func Open(options config.MySQLOptions) (*sql.DB, error) {
	const errMsg = "can't open mysql db"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", options.User, options.Password, options.Name))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	// See https://github.com/go-sql-driver/mysql#important-settings
	db.SetConnMaxLifetime(options.MaxConnLifetime)
	db.SetMaxOpenConns(options.MaxOpenConns)
	db.SetMaxIdleConns(options.MaxIdleConns)

	return db, nil
}
