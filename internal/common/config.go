package common

import (
	"fmt"
	"os"
)

type MainConfigs struct {
	AppPort  string
	HasLog   bool
	TimeZone string
}

func GetMainConfigs() MainConfigs {
	cnf := MainConfigs{AppPort: ":8000", HasLog: true, TimeZone: "Asia/Tehran"}

	if appPort := os.Getenv("APP_PORT"); appPort != "" {
		appPort = fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
		cnf.AppPort = appPort
	}

	if timeZone := os.Getenv("TIMEZONE"); timeZone != "" {
		cnf.TimeZone = timeZone
	}

	if os.Getenv("HAS_LOG") == "false" {
		cnf.HasLog = false
	}

	return cnf
}
