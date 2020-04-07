package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: time.Second * 30}
var baseUrl = "https://enpopwf3y505.x.pipedream.net"

func CadastrarUsuario(user Usuario) {

	content, err := json.Marshal(user)

	if err != nil {
		fmt.Println("Erro ao deserializar objeto ", err)
	}

	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(content))
	response, err := client.Do(req)
	//	defer response.Body.close()

	if err != nil {
		fmt.Println("Ocorreu um erro ao fazer post ", err)
		return
	}

	fmt.Println(" respnse is  ", response)
	return
}

type Usuario struct {
	Id   int    `json:id`
	Nome string `json:name`
}
