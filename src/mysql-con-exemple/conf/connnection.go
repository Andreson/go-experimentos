package conf

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	strConnection := "root:123@tcp(172.17.0.3:3306)/funcionarios"

	db, err := sql.Open(dbDriver, strConnection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
