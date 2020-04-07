package crud_book

import (
	ph_client "sync-book/ws-client-jsonplaceholder"
)

type BookEntity struct {
	Id     int
	UserId int
	Title  string
	Ok     bool
}

func ParseListBookToEntity(param []ph_client.Book) []BookEntity {

	resp := make([]BookEntity, len(param), len(param))
	for _, b := range param {
		resp = append(resp, BookEntity{Id: b.Id, Ok: b.Ok, Title: b.Title, UserId: b.UserId})
	}
	return resp
}

func ParseBookToEntity(b ph_client.Book) BookEntity {
	return BookEntity{Id: b.Id, Ok: b.Ok, Title: b.Title, UserId: b.UserId}
}

/**
Id     int `json:"userid"`
	UserId int `json:"id"`
	Title  string
	Ok     bool `json:"completed"`
*/
