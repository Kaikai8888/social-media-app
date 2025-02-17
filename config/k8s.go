//go:build k8s

package config

var Config = config{
	DBConfig: DBConfig{
		DSN: "root:root@tcp(social-media-app:3306)/social_media_app",
	},
	RedisConfig: RedisConfig{
		Addr: "social-media-app:6379",
	},
}
