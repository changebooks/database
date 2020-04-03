package database

import (
	"database/sql"
	"errors"
)

func Exec(driver *Driver, query string, args ...interface{}) (sql.Result, error) {
	if driver == nil {
		return nil, errors.New("driver can't be nil")
	}

	db := driver.GetDb()
	if db == nil {
		return nil, errors.New("db can't be nil")
	}

	return db.Exec(query, args...)
}
