package database

import (
	"database/sql"
	"errors"
	"strings"
	"sync"
)

type Driver struct {
	db     *sql.DB
	name   string  // 驱动名，mysql、postgres、...
	dsn    string  // data source name
	schema *Schema // 配置
}

func (x *Driver) Close() error {
	if x.db != nil {
		return x.db.Close()
	} else {
		return nil
	}
}

func (x *Driver) GetDb() *sql.DB {
	return x.db
}

func (x *Driver) GetName() string {
	return x.name
}

func (x *Driver) GetDsn() string {
	return x.dsn
}

func (x *Driver) GetSchema() *Schema {
	return x.schema
}

type DriverBuilder struct {
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	name   string
	schema *Schema
	joiner func(s *Schema) string // dsn拼接函数
}

func (x *DriverBuilder) Build() (*Driver, error) {
	if x.name == "" {
		return nil, errors.New("name can't be empty")
	}

	if x.schema == nil {
		return nil, errors.New("schema can't be nil")
	}

	joiner := x.joiner
	if joiner == nil {
		joiner = DsnJoiner
	}

	dsn := joiner(x.schema)
	if dsn == "" {
		return nil, errors.New("dsn can't be empty")
	}

	db, err := sql.Open(x.name, dsn)
	if err != nil {
		return nil, err
	}

	maxOpen := x.schema.GetMaxOpen()
	if maxOpen > 0 {
		db.SetMaxOpenConns(maxOpen)
	}

	maxIdle := x.schema.GetMaxIdle()
	if maxIdle > 0 {
		db.SetMaxIdleConns(maxIdle)
	}

	maxLifetime := x.schema.GetMaxLifetime()
	if maxLifetime > 0 {
		db.SetConnMaxLifetime(maxLifetime)
	}

	return &Driver{
		db:     db,
		name:   x.name,
		dsn:    dsn,
		schema: x.schema,
	}, nil
}

func (x *DriverBuilder) SetName(s string) *DriverBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.name = s
	return x
}

func (x *DriverBuilder) SetSchema(s *Schema) *DriverBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.schema = s
	return x
}

func (x *DriverBuilder) SetJoiner(f func(s *Schema) string) *DriverBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.joiner = f
	return x
}
