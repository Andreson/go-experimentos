package reqbin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: time.Second * 30}

func CadastrarCasa(casa Casa) (c string, err error) {

	if req, err := createRequest(casa); err == nil {
		if c, err := Call(req); err == nil {
			return string(c), nil
		}
	}
	return "body", err
}

func Call(req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer request", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("HTTP Status code invalido " + string(resp.StatusCode))
	}

	body, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		fmt.Println("Erro ao fazer parse body response ", erro)
		return nil, erro
	}
	return body, nil
}

func createRequest(c Casa) (*http.Request, error) {
	casaJson, err := json.Marshal(c)

	if err != nil {
		fmt.Println("Erro ao gerar json de request ", err)
		return nil, err
	}
	return http.NewRequest("POST", "https://enpopwf3y505.x.pipedream.net/", bytes.NewBuffer(casaJson))

}
