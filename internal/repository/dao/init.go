package dao

import "gorm.io/gorm"

// 並不是很好的做法, 需要確認生成的DDL
func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
