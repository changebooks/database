package database

import (
	"errors"
	"strconv"
	"time"
)

type Profile struct {
	id                string        // 唯一标识
	shardingFirst     int           // 分库开始，缺省：不分库
	shardingLast      int           // 分库结束，缺省：不分库
	shardingSeparator string        // 拼接库名和分库数
	write             bool          // 是否是主库，写库，缺省：非主库
	read              bool          // 是否是从库，只读库，缺省：非从库
	backup            bool          // 是否是备库，复杂查询，缺省：非备库
	host              string        // 域名或Ip
	username          string        // 用户名
	password          string        // 密码
	driver            string        // 驱动名，mysql、postgres、...
	proto             string        // 协议，如：tcp，缺省：define.Proto
	port              int           // 端口，如：3306，缺省：define.Port
	database          string        // 数据库名
	charset           string        // 编码，缺省：define.Charset
	collation         string        // 编码，缺省：define.Collation
	timeout           string        // 系统默认：90s
	maxOpen           int           // 最大连接数，缺省：0-不设置，无限制
	maxIdle           int           // 最大空闲连接数，缺省：0-不设置，默认2
	maxLifetime       time.Duration // 连接最大生命周期，缺省：0-不设置，永不过期
	dsn               string        // data source name，建议置空，缺省：通过host、username、password、...拼接
}

func NewProfile(data map[string]string) (*Profile, error) {
	if data == nil {
		return nil, errors.New("data can't be nil")
	}

	shardingFirst := -1
	if data[ProfileShardingFirst] != "" {
		if n, err := strconv.ParseInt(data[ProfileShardingFirst], 10, 32); err == nil {
			shardingFirst = int(n)
		} else {
			return nil, err
		}
	}

	shardingLast := -1
	if data[ProfileShardingLast] != "" {
		if n, err := strconv.ParseInt(data[ProfileShardingLast], 10, 32); err == nil {
			shardingLast = int(n)
		} else {
			return nil, err
		}
	}

	write := false
	if data[ProfileWrite] != "" {
		if b, err := strconv.ParseBool(data[ProfileWrite]); err == nil {
			write = b
		} else {
			return nil, err
		}
	}

	read := false
	if data[ProfileRead] != "" {
		if b, err := strconv.ParseBool(data[ProfileRead]); err == nil {
			read = b
		} else {
			return nil, err
		}
	}

	backup := false
	if data[ProfileBackup] != "" {
		if b, err := strconv.ParseBool(data[ProfileBackup]); err == nil {
			backup = b
		} else {
			return nil, err
		}
	}

	port := 0
	if data[ProfilePort] != "" {
		if p, err := strconv.ParseInt(data[ProfilePort], 10, 32); err == nil {
			port = int(p)
		} else {
			return nil, err
		}
	}

	maxOpen := 0
	if data[ProfileMaxOpen] != "" {
		if n, err := strconv.ParseInt(data[ProfileMaxOpen], 10, 32); err == nil {
			maxOpen = int(n)
		} else {
			return nil, err
		}
	}

	maxIdle := 0
	if data[ProfileMaxIdle] != "" {
		if n, err := strconv.ParseInt(data[ProfileMaxIdle], 10, 32); err == nil {
			maxIdle = int(n)
		} else {
			return nil, err
		}
	}

	var maxLifetime time.Duration = 0
	if data[ProfileMaxLifetime] != "" {
		if d, err := strconv.ParseInt(data[ProfileMaxLifetime], 10, 64); err == nil {
			maxLifetime = time.Duration(d)
		} else {
			return nil, err
		}
	}

	id := data[ProfileId]
	shardingSeparator := data[ProfileShardingSeparator]
	host := data[ProfileHost]
	username := data[ProfileUsername]
	password := data[ProfilePassword]
	driver := data[ProfileDriver]
	proto := data[ProfileProto]
	database := data[ProfileDatabase]
	charset := data[ProfileCharset]
	collation := data[ProfileCollation]
	timeout := data[ProfileTimeout]
	dsn := data[ProfileDsn]

	return &Profile{
		id:                id,
		shardingFirst:     shardingFirst,
		shardingLast:      shardingLast,
		shardingSeparator: shardingSeparator,
		write:             write,
		read:              read,
		backup:            backup,
		host:              host,
		username:          username,
		password:          password,
		driver:            driver,
		proto:             proto,
		port:              port,
		database:          database,
		charset:           charset,
		collation:         collation,
		timeout:           timeout,
		maxOpen:           maxOpen,
		maxIdle:           maxIdle,
		maxLifetime:       maxLifetime,
		dsn:               dsn,
	}, nil
}

func (x *Profile) GetId() string {
	return x.id
}

func (x *Profile) GetShardingFirst() int {
	return x.shardingFirst
}

func (x *Profile) GetShardingLast() int {
	return x.shardingLast
}

func (x *Profile) GetShardingSeparator() string {
	return x.shardingSeparator
}

func (x *Profile) GetWrite() bool {
	return x.write
}

func (x *Profile) GetRead() bool {
	return x.read
}

func (x *Profile) GetBackup() bool {
	return x.backup
}

func (x *Profile) GetHost() string {
	return x.host
}

func (x *Profile) GetUsername() string {
	return x.username
}

func (x *Profile) GetPassword() string {
	return x.password
}

func (x *Profile) GetDriver() string {
	return x.driver
}

func (x *Profile) GetProto() string {
	return x.proto
}

func (x *Profile) GetPort() int {
	return x.port
}

func (x *Profile) GetDatabase() string {
	return x.database
}

func (x *Profile) GetCharset() string {
	return x.charset
}

func (x *Profile) GetCollation() string {
	return x.collation
}

func (x *Profile) GetTimeout() string {
	return x.timeout
}

func (x *Profile) GetMaxOpen() int {
	return x.maxOpen
}

func (x *Profile) GetMaxIdle() int {
	return x.maxIdle
}

func (x *Profile) GetMaxLifetime() time.Duration {
	return x.maxLifetime
}

func (x *Profile) GetDsn() string {
	return x.dsn
}
