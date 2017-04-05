package db

import (
	"gowork/extern/logging"
	e "gowork/error"
	_ "database/sql"
	"regexp"
	"strings"
)

var (
	MySQL *sql.DB
)

func InitMySQL(url string) *e.WError {
	if ok, err := regexp.MatchString("^mysql://.*:.*@.*/.*$", url); ok == false || err != nil {
		logging.Error("[InitMySQL] Mysql config syntax error")
		return e.NewWError(e.ERR_CODE_PARA, "Invalid URL to MySQL server[%s], syntax error", url)
	}

	addr := strings.Replace(url, "mysql://", "", 1)
	mysql, err := sql.Open("mysql", addr)
	if err != nil {
		logging.Error("[InitMySQL] Failed to connect to MySQL server: %s, error = %s", addr, err.Error())
		return e.NewWError(e.ERR_CODE_DB, "Failed to connect to MySQL server: %s", url)
	}

	MySQL = mysql
	logging.Info("[InitMySQL] Connect to MySQL server: %s ok", addr)
	return nil
}