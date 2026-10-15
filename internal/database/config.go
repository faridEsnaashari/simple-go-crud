package database

import (
	"cmp"
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

func GetDatabaseConfigs() DatabaseConfig {

	mysqlPort, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		mysqlPort = 6000
	}

	return DatabaseConfig{
		MysqlHost:     cmp.Or(os.Getenv("MYSQL_HOST"), "localhost"),
		MysqlPort:     mysqlPort,
		MysqlPassword: cmp.Or(os.Getenv("MYSQL_PASSWORD"), "123456"),
		MysqlUsername: cmp.Or(os.Getenv("MYSQL_USERNAME"), "root"),
		MysqlDatabase: cmp.Or(os.Getenv("MYSQL_DATABASE"), "listDb"),
	}
}
