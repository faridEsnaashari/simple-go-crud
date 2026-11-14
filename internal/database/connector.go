package database

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"

	"a.com/test/internal/common"
	logx "a.com/test/internal/httpx/logs"
)

func Open(ctx context.Context) (*sql.DB, error) {
	c, cnfError := GetDatabaseConfigs()
	if cnfError != nil {
		return nil, cnfError
	}

	timeZone, timeZoneErr := time.LoadLocation(common.GetMainConfigs().TimeZone)
	if timeZoneErr != nil {
		return nil, timeZoneErr
	}

	cfg := mysql.NewConfig()
	cfg.User = c.MysqlUsername
	cfg.Passwd = c.MysqlPassword
	cfg.Net = "tcp"
	cfg.DBName = c.MysqlDatabase
	cfg.ParseTime = true // DATETIME columns come back as time.Time
	cfg.Addr = net.JoinHostPort(c.MysqlHost, strconv.Itoa(c.MysqlPort))
	cfg.Timeout = 5 * time.Second      // dial timeout
	cfg.ReadTimeout = 30 * time.Second // per-query read timeout
	cfg.WriteTimeout = 30 * time.Second
	cfg.ClientFoundRows = true
	cfg.Loc = timeZone

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	if err := db.PingContext(pingCtx); err != nil {
		_ = db.Close()

		return nil, fmt.Errorf("ping database: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(3 * time.Minute)

	logx.InfoString(
		"DATABASE_CONNECTED",
		fmt.Sprintf("connected to mysql: %v | %v | %v | %v", c.MysqlHost, c.MysqlPort, c.MysqlDatabase, timeZone),
	)

	return db, nil
}
