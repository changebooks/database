package database

import (
	"bytes"
	"fmt"
)

func DsnJoiner(s *Schema) string {
	if s == nil {
		return ""
	}

	if dsn := s.GetDsn(); dsn != "" {
		return dsn
	}

	p := DsnPath(s)
	q := DsnQuery(s)
	if q != "" {
		return p + "?" + q
	} else {
		return p
	}
}

func DsnPath(s *Schema) string {
	if s == nil {
		return ""
	} else {
		return fmt.Sprintf("%s:%s@%s(%s:%d)/%s", s.GetUsername(), s.GetPassword(), s.GetProto(), s.GetHost(), s.GetPort(), s.GetDatabase())
	}
}

func DsnQuery(s *Schema) string {
	if s == nil {
		return ""
	}

	var bucket bytes.Buffer

	charset := s.GetCharset()
	if charset != "" {
		bucket.WriteString("charset=")
		bucket.WriteString(charset)
		bucket.WriteRune('&')
	}

	collation := s.GetCollation()
	if collation != "" {
		bucket.WriteString("collation=")
		bucket.WriteString(collation)
		bucket.WriteRune('&')
	}

	timeout := s.GetTimeout()
	if timeout != "" {
		bucket.WriteString("timeout=")
		bucket.WriteString(timeout)
		bucket.WriteRune('&')
	}

	r := bucket.String()
	if r != "" {
		r = r[:len(r)-1]
	}

	return r
}
