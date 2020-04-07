package crud_book

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Conn() *sql.DB {

	db, err := sql.Open("mysql", "root:123@tcp(172.17.0.3 :3306)/books")

	if err != nil {
		fmt.Println("Ocorreu um erro nao esperado ao iniciar conexão", err)
		if err != nil {
			panic(err.Error())
		}

	}

	return db
}
