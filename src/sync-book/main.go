package main

import (
	"fmt"
	crud_book "sync-book/bancodedados"
)

/**
Aplicação que consume serviço rest e atualiza a base da dados de cadastro de livros

*/
func main() {

	//resp, _ := ph_client.GetAll()

	//	crud_book.InsertBooks(crud_book.ParseListBookToEntity(resp))

	bookEntitys := crud_book.GetAllBooks()

	fmt.Println(" ------------ livros cadastrados -------------- ")
	for _, v := range bookEntitys {
		fmt.Printf("\n Id:  %d  - Titulo : %s", v.Id, v.Title)
	}
	fmt.Println("\n ------------ livros cadastrados -------------- ")

}
