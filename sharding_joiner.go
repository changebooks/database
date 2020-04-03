package database

import "fmt"

func ShardingJoiner(database string, num int) string {
	return fmt.Sprintf("%s%s%d", database, ShardingSeparator, num)
}
