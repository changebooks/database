package database

import (
	"database/sql"
	"errors"
)

func Scan(rows *sql.Rows) ([]map[string]interface{}, error) {
	if rows == nil {
		return nil, errors.New("rows can't be nil")
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	size := len(columns)
	if size == 0 {
		return nil, errors.New("column size can't be equal than 0")
	}

	attributes := make([]interface{}, size)
	for i := range attributes {
		attributes[i] = new(interface{})
	}

	var r []map[string]interface{}
	for rows.Next() {
		if err := rows.Scan(attributes...); err != nil {
			return nil, err
		}

		pairs := make(map[string]interface{}, size)
		for i := 0; i < size; i++ {
			pairs[columns[i]] = *(attributes[i].(*interface{}))
		}

		r = append(r, pairs)
	}

	if len(r) == 0 {
		return nil, NewEmptyResult()
	}

	return r, nil
}
