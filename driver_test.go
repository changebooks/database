package database

import "testing"

func TestDriverBuilder(t *testing.T) {
	b := &DriverBuilder{}
	_, got := b.Build()
	want := "name can't be empty"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	b.SetName("mysql")
	_, got2 := b.Build()
	want2 := "schema can't be nil"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}
