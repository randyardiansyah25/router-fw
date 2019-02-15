package connection

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func MysqlConnect(host, port, uname, pass, dbname string) error {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		uname,
		pass,
		host,
		port,
		dbname,
	)

	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}
