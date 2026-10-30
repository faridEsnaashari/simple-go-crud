package common

import (
	"fmt"
	"os"
)

type MainConfigs struct {
	AppPort string
	HasLog  bool
}

func GetMainConfigs() MainConfigs {

	cnf := MainConfigs{AppPort: ":8000", HasLog: true}

	if appPort := os.Getenv("APP_PORT"); appPort != "" {
		appPort = fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
		cnf.AppPort = appPort
	}

	if os.Getenv("HAS_LOG") == "false" {
		cnf.HasLog = false
	}

	return cnf
}
