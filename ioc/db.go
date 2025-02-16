package ioc

import (
	"webook/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DBConfig.DSN))
	if err != nil {
		panic(err)
	}

	// err = dao.InitTables(db)
	// if err != nil {
	// 	panic(err)
	// }
	return db
}
