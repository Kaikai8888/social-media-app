//go:build k8s

package config

var Config = config{
	DBConfig: DBConfig{
		DSN: "root:root@tcp(localhost:3306)/webook",
	},
	RedisConfig: RedisConfig{
		Addr: "localhost:6379",
	},
}
