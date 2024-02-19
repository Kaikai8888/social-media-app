//go:build k8s

package config

var Config = config{
	DBConfig: DBConfig{
		DSN: "root:root@tcp(webook:3306)/webook",
	},
	RedisConfig: RedisConfig{
		Addr: "webook:6379",
	},
}
