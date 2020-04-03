package database

import (
	"errors"
	"testing"
)

func TestIsEmptyResult(t *testing.T) {
	got := IsEmptyResult(NewEmptyResult())
	if got != true {
		t.Errorf("got %v; want true", got)
	}

	got2 := IsEmptyResult(nil)
	if got2 != false {
		t.Errorf("got %v; want true", got2)
	}

	got3 := IsEmptyResult(errors.New(""))
	if got3 != false {
		t.Errorf("got %v; want true", got3)
	}
}
