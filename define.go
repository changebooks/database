package database

const (
	Charset           = "utf8mb4"
	Collation         = "utf8mb4_general_ci"
	Proto             = "tcp"
	Port              = 3306
	Write             = 1              // 主库，写库
	Read              = 2              // 从库，只读库
	Backup            = 4              // 备库，复杂查询
	EmptyResult       = "empty result" // 查询结果空
	AggregateAlias    = "aggregate"    // 统计字段别名
	ShardingSeparator = "_"            // 拼接库名和分库数
)
