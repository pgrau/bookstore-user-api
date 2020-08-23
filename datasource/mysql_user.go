package datasource

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	mysql_user_username = "mysql_user_username"
	mysql_user_password = "mysql_user_password"
	mysql_user_host     = "mysql_user_host"
	mysql_user_schema   = "mysql_user_schema"
)

var (
	MysqlClient *sql.DB

	username = os.Getenv(mysql_user_username)
	password = os.Getenv(mysql_user_password)
	host     = os.Getenv(mysql_user_host)
	schema   = os.Getenv(mysql_user_schema)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", username, password, host, schema,)

	var err error
	MysqlClient, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = MysqlClient.Ping(); err != nil {
		panic(err)
	}
}
