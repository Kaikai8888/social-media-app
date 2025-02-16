package util

import (
	"fmt"
	"testing"
	"webook/internal/repository/dao"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	tablesToInit := []interface{}{
		dao.User{},
		dao.DraftArticle{},
	}

	for _, table := range tablesToInit {
		if err := db.AutoMigrate(table); err != nil {
			panic(err)
		}
	}

	return nil
}

func TruncateTables(t *testing.T, db *gorm.DB, tables ...string) {
	for _, table := range tables {
		if err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table)).Error; err != nil {
			assert.FailNow(t, fmt.Sprintf("failed to truncate table %s", table), err)
		}
	}
}
