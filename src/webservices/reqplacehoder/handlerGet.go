package reqplacehoder

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"webservices/reqbin"
)

func GetBook(id int) (Book, error) {
	req := createRequest(id)

	if byteResponse, err := reqbin.Call(req); err == nil {
		respoModel := Book{}
		if error := json.Unmarshal(byteResponse, &respoModel); error == nil {
			return respoModel, nil
		} else {
			return Book{}, error
		}
	} else {
		fmt.Println("Ocorreu um erro nao esperado ao fazer request")
		return Book{}, errors.New("Erro ao fazer request ")
	}
}

func createRequest(id int) *http.Request {
	url := "https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(id)

	if req, err := http.NewRequest("GET", url, nil); err == nil {
		return req
	} else {
		fmt.Println("Erro ao criar Request ", err)
		return nil
	}
}
