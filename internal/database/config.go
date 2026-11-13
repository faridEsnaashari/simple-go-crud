package database

import (
	"fmt"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	MysqlHost     string `json:"mysqlHost"`
	MysqlPort     int    `json:"mysqlPort"`
	MysqlPassword string `json:"mysqlPassword"`
	MysqlUsername string `json:"mysqlUsername"`
	MysqlDatabase string `json:"mysqlDatabase"`
}

func GetDatabaseConfigs() (DatabaseConfig, error) {
	mysqlPort, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPasswor := os.Getenv("MYSQL_PASSWORD")
	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	if mysqlPort == 0 ||
		err != nil ||
		mysqlHost == "" ||
		mysqlPasswor == "" ||
		mysqlUsername == "" ||
		mysqlDatabase == "" {
		return DatabaseConfig{},
			fmt.Errorf("wrong database confic")
	}

	return DatabaseConfig{
		MysqlHost:     mysqlHost,
		MysqlPort:     mysqlPort,
		MysqlPassword: mysqlPasswor,
		MysqlUsername: mysqlUsername,
		MysqlDatabase: mysqlDatabase,
	}, nil
}
