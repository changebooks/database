package database

import "testing"

func TestDsnJoiner(t *testing.T) {
	got := DsnJoiner(nil)
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := DsnJoiner(&Schema{
		proto:       Proto,
		host:        "127.0.0.1",
		port:        Port,
		database:    "test",
		username:    "root",
		password:    "123456",
		charset:     Charset,
		collation:   Collation,
		timeout:     "90s",
		maxOpen:     0,
		maxIdle:     0,
		maxLifetime: 0,
		dsn:         "",
	})
	want2 := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&collation=utf8mb4_general_ci&timeout=90s"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := DsnJoiner(&Schema{
		proto:       Proto,
		host:        "127.0.0.1",
		port:        Port,
		database:    "test",
		username:    "root",
		password:    "123456",
		charset:     Charset,
		collation:   Collation,
		timeout:     "90s",
		maxOpen:     0,
		maxIdle:     0,
		maxLifetime: 0,
		dsn:         "admin:123123@tcp(192.168.0.1:3306)/user?charset=utf8&collation=utf8_general_ci&timeout=10s",
	})
	want3 := "admin:123123@tcp(192.168.0.1:3306)/user?charset=utf8&collation=utf8_general_ci&timeout=10s"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}
}

func TestDsnPath(t *testing.T) {
	got := DsnPath(nil)
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := DsnPath(&Schema{
		proto:       Proto,
		host:        "127.0.0.1",
		port:        Port,
		database:    "test",
		username:    "root",
		password:    "123456",
		charset:     Charset,
		collation:   Collation,
		timeout:     "90s",
		maxOpen:     0,
		maxIdle:     0,
		maxLifetime: 0,
		dsn:         "",
	})
	want2 := "root:123456@tcp(127.0.0.1:3306)/test"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}

func TestDsnQuery(t *testing.T) {
	got := DsnQuery(nil)
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := DsnQuery(&Schema{
		proto:       Proto,
		host:        "127.0.0.1",
		port:        Port,
		database:    "test",
		username:    "root",
		password:    "123456",
		charset:     Charset,
		collation:   Collation,
		timeout:     "90s",
		maxOpen:     0,
		maxIdle:     0,
		maxLifetime: 0,
		dsn:         "",
	})
	want2 := "charset=utf8mb4&collation=utf8mb4_general_ci&timeout=90s"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := DsnQuery(&Schema{
		proto:       Proto,
		host:        "127.0.0.1",
		port:        Port,
		database:    "test",
		username:    "root",
		password:    "123456",
		charset:     Charset,
		collation:   "",
		timeout:     "90s",
		maxOpen:     0,
		maxIdle:     0,
		maxLifetime: 0,
		dsn:         "",
	})
	want3 := "charset=utf8mb4&timeout=90s"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4 := DsnQuery(&Schema{
		proto:       Proto,
		host:        "127.0.0.1",
		port:        Port,
		database:    "test",
		username:    "root",
		password:    "123456",
		charset:     "",
		collation:   "",
		timeout:     "90s",
		maxOpen:     0,
		maxIdle:     0,
		maxLifetime: 0,
		dsn:         "",
	})
	want4 := "timeout=90s"
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	got5 := DsnQuery(&Schema{
		proto:       Proto,
		host:        "127.0.0.1",
		port:        Port,
		database:    "test",
		username:    "root",
		password:    "123456",
		charset:     "",
		collation:   "",
		timeout:     "",
		maxOpen:     0,
		maxIdle:     0,
		maxLifetime: 0,
		dsn:         "",
	})
	want5 := ""
	if got5 != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}

	got6 := DsnQuery(&Schema{
		proto:       Proto,
		host:        "127.0.0.1",
		port:        Port,
		database:    "test",
		username:    "root",
		password:    "123456",
		charset:     Charset,
		collation:   "",
		timeout:     "",
		maxOpen:     0,
		maxIdle:     0,
		maxLifetime: 0,
		dsn:         "",
	})
	want6 := "charset=utf8mb4"
	if got6 != want6 {
		t.Errorf("got %q; want %q", got6, want6)
	}

	got7 := DsnQuery(&Schema{
		proto:       Proto,
		host:        "127.0.0.1",
		port:        Port,
		database:    "test",
		username:    "root",
		password:    "123456",
		charset:     "",
		collation:   Collation,
		timeout:     "",
		maxOpen:     0,
		maxIdle:     0,
		maxLifetime: 0,
		dsn:         "",
	})
	want7 := "collation=utf8mb4_general_ci"
	if got7 != want7 {
		t.Errorf("got %q; want %q", got7, want7)
	}
}
