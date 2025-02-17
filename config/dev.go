//go:build !k8s

package config

var Config = config{
	DBConfig: DBConfig{
		DSN: "root:root@tcp(localhost:13316)/social_media_app",
	},
	RedisConfig: RedisConfig{
		Addr: "localhost:6379",
	},
}
