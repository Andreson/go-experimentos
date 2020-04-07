package crud_book

import (
	"fmt"
	"log"
)

func InsertBooks(bs []BookEntity) {

	for _, v := range bs {
		InsertBook(v)
	}
}

func InsertBook(b BookEntity) {
	db := Conn()

	if insert, err := db.Prepare("INSERT INTO book (userId,title,ok) VALUES (?,?,? )"); err == nil {
		log.Println("INSERT NEW BOOK : ", b)
		if _, err := insert.Exec(b.UserId, b.Title, b.Ok); err != nil {
			log.Println("Erro ao executar insert : ", err)
		}
	} else {
		log.Println("Ocorreu um erro nao esperado  : ", err)
	}
	defer db.Close()

}

func GetAllBooks() []BookEntity {
	db := Conn()
	selDB, err := db.Query("SELECT * FROM book")
	if err != nil {
		fmt.Println("Erro ao executar query select ", err)
	}
	bookResp := make([]BookEntity, 0, 0)
	for selDB.Next() {
		var id, userId int
		var title string
		var ok bool
		selDB.Scan(&id, &userId, &title, &ok)
		bookResp = append(bookResp, BookEntity{Id: id, Title: title, Ok: ok, UserId: userId})
	}
	return bookResp
}
