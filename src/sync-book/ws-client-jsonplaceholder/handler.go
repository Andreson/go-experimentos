package ph_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	handler_resty "sync-book/ws-handler-resty"
)

var hostWsUrl = "https://jsonplaceholder.typicode.com/todos/"

func GetAll() ([]Book, error) {

	if resp, err := handler_resty.Get(hostWsUrl); err == nil {
		dataResp := make([]Book, 1)
		if err := json.Unmarshal(resp.Body(), &dataResp); err == nil {
			return dataResp, nil
		} else {
			fmt.Println("error Unmarshal ", resp.Body())
			return nil, err
		}
	}
	return nil, errors.New("Ocorreu um erro nao esperado ao consumir serviço")
}

func FindById(id int) (Book, error) {

	if resp, err := handler_resty.Get(hostWsUrl + strconv.Itoa(id)); err == nil {
		dataResp := Book{}
		if err := json.Unmarshal(resp.Body(), &dataResp); err == nil {
			return dataResp, nil
		} else {
			fmt.Println("error Unmarshal ", resp.Body())
			return Book{}, err
		}
	}
	return Book{}, errors.New("Ocorreu um erro nao esperado ao consumir serviço")

}
