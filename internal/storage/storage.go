package storage

import (
	"database/sql"
	"fmt"

	"github.com/qo/digital-library/internal/config"
	"github.com/qo/digital-library/internal/storage/mysql"
	"github.com/qo/digital-library/internal/storage/sqlite"
)

type Storage struct {
	db *sql.DB
}

const (
	mysqlDb  = "mysql"
	sqliteDb = "sqlite"
)

func Init(options config.StorageOptions) (*Storage, error) {
	const errMsg = "can't init storage"

	var (
		db  *sql.DB
		err error
	)

	switch options.Db {
	case mysqlDb:
		db, err = mysql.Open(options.MySQLOptions)
	case sqliteDb:
		db, err = sqlite.Open(options.SQLiteOptions)
	default:
		err = fmt.Errorf("db option %s is unknown", options.Db)
	}

	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	st := Storage{db}

	err = st.initUsers()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	return &st, err
}
