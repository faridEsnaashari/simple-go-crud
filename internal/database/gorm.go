package database

import (
	"database/sql"
	"fmt"

	"a.com/test/internal/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var repo *gorm.DB

func NewGorm(sqlDB *sql.DB) (*gorm.DB, error) {
	log := logger.Default.LogMode(logger.Info)

	if common.GetMainConfigs().HasLog != true {
		log = logger.Default.LogMode(logger.Error)
	}

	initDb, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		Logger:         log,
		TranslateError: true,
	})
	if err != nil {
		return nil, fmt.Errorf("open gorm: %w", err)
	}

	repo = initDb

	return repo, nil
}

func GetRepo() *gorm.DB {
	return repo
}
