package mysql

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	"github.com/ecsaas/mysql-define-config/DEFINE_VARIABLES/MYSQL_ENV"
	_ "github.com/go-sql-driver/mysql"
)

var SqlDb *sql.DB
var SqlErr error

func SqlInit() (db *sql.DB) {
	var applyURI = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		url.QueryEscape(os.Getenv(MYSQL_ENV.MYSQL_USER)),
		url.QueryEscape(os.Getenv(MYSQL_ENV.MYSQL_PASSWORD)),
		os.Getenv(MYSQL_ENV.MYSQL_IP),
		os.Getenv(MYSQL_ENV.MYSQL_DB_NAME),
	)
	SqlDb, SqlErr = sql.Open("mysql", applyURI)
	return
}
