package database

import "errors"

func NewEmptyResult() error {
	return errors.New(EmptyResult)
}

func IsEmptyResult(err error) bool {
	if err == nil {
		return false
	} else {
		return EmptyResult == err.Error()
	}
}
