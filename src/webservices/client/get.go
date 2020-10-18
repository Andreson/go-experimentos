package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func BuscarUsuario(id int) {
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos/2", nil)

	if err != nil {
		fmt.Println("Erro ao criar GET request ", err)
		return
	}
	resp, respError := client.Do(req)

	body, errReadall := ioutil.ReadAll(resp.Body)

	if errReadall != nil {
		fmt.Println("Erro ao criar errReadall  ", errReadall)
		return
	}

	if respError != nil {
		fmt.Println("Erro ao criar GET request ", respError)
		return
	}
	respoModel := []Book{}
	error := json.Unmarshal(body, &respoModel)

	if error != nil {
		fmt.Println("Erro Unmarshal ", error)
		return
	}

	fmt.Println("respoModel  ", respoModel)

}

type Book struct {
	Id   int    `json:"userId"`
	Nome string `json:"title"`
}
