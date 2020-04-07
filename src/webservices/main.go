package main

import (
	"fmt"
	"webservices/reqplacehoder"
)

func main() {

	//	resp, err := reqbin.CadastrarCasa(reqbin.Casa{Id: 666, Endereco: "Av piassanguaba", Bairro: "Saúde"})
	resp3, err := reqplacehoder.GetBook(3)
	fmt.Println("  resp3", resp3, err)
}
