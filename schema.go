package database

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type Schema struct {
	mu          sync.Mutex    // ensures atomic writes; protects the following fields
	proto       string        // 协议，如：tcp，缺省：define.Proto
	host        string        // 域名或Ip
	port        int           // 端口，如：3306，缺省：define.Port
	database    string        // 数据库名
	username    string        // 用户名
	password    string        // 密码
	charset     string        // 缺省：define.Charset
	collation   string        // 缺省：define.Collation
	timeout     string        // 系统默认：90s
	maxOpen     int           // 最大连接数，缺省：0，不设置，无限制
	maxIdle     int           // 最大空闲连接数，缺省：0，不设置，默认2
	maxLifetime time.Duration // 连接最大生命周期，缺省：0，不设置，永不过期
	dsn         string        // data source name
}

func (x *Schema) ToString() string {
	return fmt.Sprintf("proto:       %v\n"+
		"host:        %v\n"+
		"port:        %v\n"+
		"database:    %v\n"+
		"username:    %v\n"+
		"password:    %v\n"+
		"charset:     %v\n"+
		"collation:   %v\n"+
		"timeout:     %v\n"+
		"maxOpen:     %v\n"+
		"maxIdle:     %v\n"+
		"maxLifetime: %v\n"+
		"dsn:         %v\n",
		x.GetProto(), x.GetHost(), x.GetPort(), x.GetDatabase(), x.GetUsername(), x.GetPassword(),
		x.GetCharset(), x.GetCollation(), x.GetTimeout(),
		x.GetMaxOpen(), x.GetMaxIdle(), x.GetMaxLifetime(), x.GetDsn(),
	)
}

func (x *Schema) GetProto() string {
	return x.proto
}

func (x *Schema) GetHost() string {
	return x.host
}

func (x *Schema) GetPort() int {
	return x.port
}

func (x *Schema) GetDatabase() string {
	return x.database
}

func (x *Schema) GetUsername() string {
	return x.username
}

func (x *Schema) GetPassword() string {
	return x.password
}

func (x *Schema) GetCharset() string {
	return x.charset
}

func (x *Schema) GetCollation() string {
	return x.collation
}

func (x *Schema) GetTimeout() string {
	return x.timeout
}

func (x *Schema) GetMaxOpen() int {
	return x.maxOpen
}

func (x *Schema) GetMaxIdle() int {
	return x.maxIdle
}

func (x *Schema) GetMaxLifetime() time.Duration {
	return x.maxLifetime
}

func (x *Schema) GetDsn() string {
	return x.dsn
}

type SchemaBuilder struct {
	mu          sync.Mutex // ensures atomic writes; protects the following fields
	proto       string
	host        string
	port        int
	database    string
	username    string
	password    string
	charset     string
	collation   string
	timeout     string
	maxOpen     int
	maxIdle     int
	maxLifetime time.Duration
	dsn         string
}

func (x *SchemaBuilder) Build() (*Schema, error) {
	if x.host == "" {
		return nil, errors.New("host can't be empty")
	}

	if x.port < 0 {
		return nil, errors.New("port can't be less than 0")
	}

	if x.database == "" {
		return nil, errors.New("database can't be empty")
	}

	if x.username == "" {
		return nil, errors.New("username can't be empty")
	}

	if x.maxOpen < 0 {
		return nil, errors.New("max open can't be less than 0")
	}

	if x.maxIdle < 0 {
		return nil, errors.New("max idle can't be less than 0")
	}

	if x.maxLifetime < 0 {
		return nil, errors.New("max lifetime can't be less than 0")
	}

	proto := x.proto
	if proto == "" {
		proto = Proto
	}

	port := x.port
	if port == 0 {
		port = Port
	}

	charset := x.charset
	if charset == "" {
		charset = Charset
	}

	collation := x.collation
	if collation == "" {
		collation = Collation
	}

	return &Schema{
		proto:       proto,
		host:        x.host,
		port:        port,
		database:    x.database,
		username:    x.username,
		password:    x.password,
		charset:     charset,
		collation:   collation,
		timeout:     x.timeout,
		maxOpen:     x.maxOpen,
		maxIdle:     x.maxIdle,
		maxLifetime: x.maxLifetime,
		dsn:         x.dsn,
	}, nil
}

func (x *SchemaBuilder) SetProto(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.proto = s
	return x
}

func (x *SchemaBuilder) SetHost(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.host = s
	return x
}

func (x *SchemaBuilder) SetPort(p int) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.port = p
	return x
}

func (x *SchemaBuilder) SetDatabase(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.database = s
	return x
}

func (x *SchemaBuilder) SetUsername(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.username = s
	return x
}

func (x *SchemaBuilder) SetPassword(s string) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.password = s
	return x
}

func (x *SchemaBuilder) SetCharset(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.charset = s
	return x
}

func (x *SchemaBuilder) SetCollation(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.collation = s
	return x
}

func (x *SchemaBuilder) SetTimeout(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.timeout = s
	return x
}

func (x *SchemaBuilder) SetMaxOpen(n int) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.maxOpen = n
	return x
}

func (x *SchemaBuilder) SetMaxIdle(n int) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.maxIdle = n
	return x
}

func (x *SchemaBuilder) SetMaxLifetime(d time.Duration) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.maxLifetime = d
	return x
}

func (x *SchemaBuilder) SetDsn(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.dsn = s
	return x
}
