package blog

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var sqlEngine *sql.DB

func init() {
	var err error
	sqlEngine, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go?charset=utf8")
	if err != nil {
		panic(err)
	}

}
