package database

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Sharding struct {
	id      string    // 唯一标识
	size    int       // 分库数
	drivers []*Driver // 主库列表，分库时，不支持从库和备库
}

// 取第n个库
func (x *Sharding) GetDriver(num int) (*Driver, error) {
	if x.drivers == nil {
		return nil, errors.New("drivers can't be nil")
	}

	if x.size <= 0 {
		return nil, fmt.Errorf("size %d can't be less or equal than 0", x.size)
	}

	if num < 0 {
		return nil, fmt.Errorf("num %d can't be less than 0", num)
	}

	if num >= x.size {
		return nil, fmt.Errorf("num %d can't be greater or equal than size %d", num, x.size)
	}

	return x.drivers[num], nil
}

func (x *Sharding) GetId() string {
	return x.id
}

func (x *Sharding) GetSize() int {
	return x.size
}

func (x *Sharding) GetDrivers() []*Driver {
	return x.drivers
}

// 关闭全部sql.DB
func (x *Sharding) Close() []error {
	var r []error

	if x.drivers != nil {
		for _, d := range x.drivers {
			if err := d.Close(); err != nil {
				r = append(r, err)
			}
		}
	}

	return r
}

type ShardingBuilder struct {
	mu             sync.Mutex // ensures atomic writes; protects the following fields
	id             string
	name           string                                // 驱动名，mysql、postgres、...
	dsnJoiner      func(s *Schema) string                // dsn拼接函数
	shardingJoiner func(database string, num int) string // sharding拼接函数
	schemas        map[int]*Schema                       // 第n个库 => 库
}

func (x *ShardingBuilder) Build() (*Sharding, error) {
	if x.id == "" {
		return nil, errors.New("id can't be empty")
	}

	if x.name == "" {
		return nil, errors.New("name can't be empty")
	}

	if x.schemas == nil {
		return nil, errors.New("schemas can't be nil")
	}

	size := len(x.schemas)
	if size == 0 {
		return nil, errors.New("schemas can't be empty")
	}

	for num, s := range x.schemas {
		if s == nil {
			return nil, errors.New("schema can't be nil")
		}

		if num < 0 {
			return nil, errors.New("num can't be less than 0")
		}

		if num >= size {
			return nil, fmt.Errorf("num %d can't be greater or equal than schema's size %d", num, size)
		}
	}

	dsnJoiner := x.dsnJoiner
	if dsnJoiner == nil {
		dsnJoiner = DsnJoiner
	}

	shardingJoiner := x.shardingJoiner
	if shardingJoiner == nil {
		shardingJoiner = ShardingJoiner
	}

	builder := DriverBuilder{}
	builder.SetName(x.name).SetJoiner(dsnJoiner)

	var drivers []*Driver
	for num := 0; num < size; num++ {
		s := x.schemas[num]
		s.database = shardingJoiner(s.GetDatabase(), num)

		d, err := builder.SetSchema(s).Build()
		if err != nil {
			return nil, err
		}

		drivers = append(drivers, d)
	}

	return &Sharding{
		id:      x.id,
		size:    len(drivers),
		drivers: drivers,
	}, nil
}

func (x *ShardingBuilder) SetId(s string) error {
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

func (x *ShardingBuilder) SetName(s string) error {
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

func (x *ShardingBuilder) SetDsnJoiner(f func(s *Schema) string) *ShardingBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.dsnJoiner = f
	return x
}

func (x *ShardingBuilder) SetShardingJoiner(f func(database string, num int) string) *ShardingBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.shardingJoiner = f
	return x
}

func (x *ShardingBuilder) SetSchema(num int, s *Schema) error {
	if num < 0 {
		return errors.New("num can't be less than 0")
	}

	if s == nil {
		return errors.New("schema can't be nil")
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	if x.schemas == nil {
		x.schemas = make(map[int]*Schema)
	}

	if x.schemas[num] != nil {
		return fmt.Errorf("num %d has been contained in schemas", num)
	}

	x.schemas[num] = s
	return nil
}

func (x *ShardingBuilder) AddProfile(p *Profile) error {
	if p == nil {
		return errors.New("profile can't be nil")
	}

	first := p.GetShardingFirst()
	last := p.GetShardingLast()
	id := p.GetId()
	name := p.GetDriver()

	if first < 0 {
		return fmt.Errorf("first %d can't be less than 0", first)
	}

	if last <= 0 {
		return fmt.Errorf("last %d can't be less or equal than 0", last)
	}

	if first >= last {
		return fmt.Errorf("first %d can't be greater or equal than last %d", first, last)
	}

	if err := x.SetId(id); err != nil {
		return err
	}

	if err := x.SetName(name); err != nil {
		return err
	}

	for num := first; num <= last; num++ {
		schema, err := NewSchema(p)
		if err != nil {
			return err
		}

		if err := x.SetSchema(num, schema); err != nil {
			return err
		}
	}

	return nil
}
