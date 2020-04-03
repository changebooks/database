package database

import (
	"fmt"
	"strconv"
)

// "AS 'aggregate'" must be contained in Query
func AggregateInt(driver *Driver, query string, args ...interface{}) (result int64, err error, closeErr error) {
	r, err, closeErr := Aggregate(driver, query, args...)
	if err != nil {
		return
	}

	if r == nil {
		return
	}

	switch r.(type) {
	case int64:
		result = r.(int64)
		break
	case []uint8:
		result, err = strconv.ParseInt(string(r.([]uint8)), 10, 64)
		break
	default:
		err = fmt.Errorf("unsupported type %v", r)
		break
	}

	return
}

// "AS 'aggregate'" must be contained in Query
func AggregateFloat(driver *Driver, query string, args ...interface{}) (result float64, err error, closeErr error) {
	r, err, closeErr := Aggregate(driver, query, args...)
	if err != nil {
		return
	}

	if r == nil {
		return
	}

	switch r.(type) {
	case float64:
		result = r.(float64)
		break
	case []uint8:
		result, err = strconv.ParseFloat(string(r.([]uint8)), 64)
		break
	default:
		err = fmt.Errorf("unsupported type %v", r)
		break
	}

	return
}

// "AS 'aggregate'" must be contained in Query
func Aggregate(driver *Driver, query string, args ...interface{}) (result interface{}, err error, closeErr error) {
	r, err, closeErr := First(driver, query, args...)
	if err != nil {
		return
	}

	if r2, ok := r[AggregateAlias]; ok {
		result = r2
	} else {
		err = fmt.Errorf(`"%s" must be contained in map`, AggregateAlias)
	}

	return
}
