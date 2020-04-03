package database

import "errors"

func First(driver *Driver, query string, args ...interface{}) (result map[string]interface{}, err error, closeErr error) {
	if driver == nil {
		err = errors.New("driver can't be nil")
		return
	}

	db := driver.GetDb()
	if db == nil {
		err = errors.New("db can't be nil")
		return
	}

	rows, err := db.Query(query, args...)
	if err == nil {
		result, err = ScanFirst(rows)
	}

	if rows != nil {
		closeErr = rows.Close()
	}

	return
}
