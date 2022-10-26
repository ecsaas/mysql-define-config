package mysql

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	MysqlUser     string
	MysqlPassword string
	MysqlIp       string
	MysqlDbName   string
}

func SqlInit(c Connection) (*sql.DB, error) {
	var applyURI = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		url.QueryEscape(c.MysqlUser),
		url.QueryEscape(c.MysqlPassword),
		c.MysqlIp,
		c.MysqlDbName,
	)
	return sql.Open("mysql", applyURI)
}
