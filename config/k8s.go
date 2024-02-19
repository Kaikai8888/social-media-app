//go:build k8s

package config

var Config = config{
	DBConfig: DBConfig{
		DSN: "root:root@tcp(localhost:30006)/webook",
	},
	RedisConfig: RedisConfig{
		Addr: "localhost:30007",
	},
}
