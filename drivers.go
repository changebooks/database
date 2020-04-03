package database

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

type Drivers struct {
	id      string    // 唯一标识
	writers []*Driver // 主库列表
	readers []*Driver // 从库列表
	backups []*Driver // 备库列表
}

// 主库列表，随机取一个Driver
func (x *Drivers) GetWriter() (*Driver, error) {
	if x.writers == nil {
		return nil, errors.New("writers can't be nil")
	}

	size := len(x.writers)
	if size == 0 {
		return nil, errors.New("writers can't be empty")
	}

	num := 0
	if size > 1 {
		num = rand.Intn(size)
	}

	return x.writers[num], nil
}

// 从库列表，随机取一个Driver
func (x *Drivers) GetReader() (*Driver, error) {
	if x.readers == nil {
		return nil, errors.New("readers can't be nil")
	}

	size := len(x.readers)
	if size == 0 {
		return nil, errors.New("readers can't be empty")
	}

	num := 0
	if size > 1 {
		num = rand.Intn(size)
	}

	return x.readers[num], nil
}

// 备库列表，随机取一个Driver
func (x *Drivers) GetBackup() (*Driver, error) {
	if x.backups == nil {
		return nil, errors.New("backups can't be nil")
	}

	size := len(x.backups)
	if size == 0 {
		return nil, errors.New("backups can't be empty")
	}

	num := 0
	if size > 1 {
		num = rand.Intn(size)
	}

	return x.backups[num], nil
}

func (x *Drivers) GetId() string {
	return x.id
}

func (x *Drivers) GetWriters() []*Driver {
	return x.writers
}

func (x *Drivers) GetReaders() []*Driver {
	return x.readers
}

func (x *Drivers) GetBackups() []*Driver {
	return x.backups
}

// 关闭全部sql.DB
func (x *Drivers) Close() []error {
	var r []error

	if x.writers != nil {
		for _, driver := range x.writers {
			if err := driver.Close(); err != nil {
				r = append(r, err)
			}
		}
	}

	if x.readers != nil {
		for _, driver := range x.readers {
			if err := driver.Close(); err != nil {
				r = append(r, err)
			}
		}
	}

	if x.backups != nil {
		for _, driver := range x.backups {
			if err := driver.Close(); err != nil {
				r = append(r, err)
			}
		}
	}

	return r
}

type DriversBuilder struct {
	mu        sync.Mutex // ensures atomic writes; protects the following fields
	id        string
	name      string                 // 驱动名，mysql、postgres、...
	dsnJoiner func(s *Schema) string // dsn拼接函数
	writers   []*Schema              // 主库列表
	readers   []*Schema              // 从库列表
	backups   []*Schema              // 备库列表
}

func (x *DriversBuilder) Build() (*Drivers, error) {
	if x.id == "" {
		return nil, errors.New("id can't be empty")
	}

	if x.name == "" {
		return nil, errors.New("name can't be empty")
	}

	dsnJoiner := x.dsnJoiner
	if dsnJoiner == nil {
		dsnJoiner = DsnJoiner
	}

	builder := DriverBuilder{}
	builder.SetName(x.name).SetJoiner(dsnJoiner)

	var w []*Driver
	if x.writers != nil {
		for _, schema := range x.writers {
			if schema == nil {
				return nil, errors.New("writer's schema can't be nil")
			}

			driver, err := builder.SetSchema(schema).Build()
			if err != nil {
				return nil, err
			}

			w = append(w, driver)
		}
	}

	var r []*Driver
	if x.readers != nil {
		for _, schema := range x.readers {
			if schema == nil {
				return nil, errors.New("reader's schema can't be nil")
			}

			driver, err := builder.SetSchema(schema).Build()
			if err != nil {
				return nil, err
			}

			r = append(r, driver)
		}
	}

	var b []*Driver
	if x.backups != nil {
		for _, schema := range x.backups {
			if schema == nil {
				return nil, errors.New("backup's schema can't be nil")
			}

			driver, err := builder.SetSchema(schema).Build()
			if err != nil {
				return nil, err
			}

			b = append(b, driver)
		}
	}

	if w == nil && r == nil && b == nil {
		return nil, errors.New("no driver, writer & reader & backup is nil")
	}

	return &Drivers{
		id:      x.id,
		writers: w,
		readers: r,
		backups: b,
	}, nil
}

func (x *DriversBuilder) SetId(s string) error {
	if s = strings.TrimSpace(s); s == "" {
		return errors.New("id can't be empty")
	}

	if x.id == s {
		return nil
	}

	if x.id != "" {
		return fmt.Errorf(`id "%s" must be equal than x.id "%s"`, s, x.id)
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	x.id = s
	return nil
}

func (x *DriversBuilder) SetName(s string) error {
	if s = strings.TrimSpace(s); s == "" {
		return errors.New("name can't be empty")
	}

	if x.name == s {
		return nil
	}

	if x.name != "" {
		return fmt.Errorf(`name "%s" must be equal than x.name "%s"`, s, x.name)
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	x.name = s
	return nil
}

func (x *DriversBuilder) SetDsnJoiner(f func(s *Schema) string) *DriversBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.dsnJoiner = f
	return x
}

func (x *DriversBuilder) AddWriter(s *Schema) error {
	if s == nil {
		return errors.New("schema can't be nil")
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	x.writers = append(x.writers, s)
	return nil
}

func (x *DriversBuilder) AddReader(s *Schema) error {
	if s == nil {
		return errors.New("schema can't be nil")
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	x.readers = append(x.readers, s)
	return nil
}

func (x *DriversBuilder) AddBackup(s *Schema) error {
	if s == nil {
		return errors.New("schema can't be nil")
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	x.backups = append(x.backups, s)
	return nil
}

func (x *DriversBuilder) AddSchema(s *Schema, write bool, read bool, backup bool) error {
	if write {
		if err := x.AddWriter(s); err != nil {
			return err
		}
	}

	if read {
		if err := x.AddReader(s); err != nil {
			return err
		}
	}

	if backup {
		if err := x.AddBackup(s); err != nil {
			return err
		}
	}

	return nil
}

func (x *DriversBuilder) AddProfile(p *Profile) error {
	if p == nil {
		return errors.New("profile can't be nil")
	}

	id := p.GetId()
	name := p.GetDriver()
	write := p.GetWrite()
	read := p.GetRead()
	backup := p.GetBackup()

	if err := x.SetId(id); err != nil {
		return err
	}

	if err := x.SetName(name); err != nil {
		return err
	}

	schema, err := NewSchema(p)
	if err != nil {
		return err
	}

	return x.AddSchema(schema, write, read, backup)
}
