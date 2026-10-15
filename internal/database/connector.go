package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

func Open(c DatabaseConfig) (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = c.MysqlUsername
	cfg.Passwd = c.MysqlPassword
	cfg.Net = "tcp"
	cfg.Addr = c.MysqlHost
	cfg.DBName = c.MysqlDatabase
	cfg.ParseTime = true // DATETIME columns come back as time.Time

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	fmt.Printf("connected to mysql: %v | %v | %v", c.MysqlHost, c.MysqlPort, c.MysqlDatabase)
	fmt.Println()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(3 * time.Minute)

	return db, nil
}
