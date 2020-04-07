package main

import (
	"fmt"
	"mysql-con-exemple/dao"
	"mysql-con-exemple/model"
)

func main() {

	dao.CadastrarEmployee(model.Employee{Id: 2, Name: "Jurema", City: "Belem"})

	emps := dao.BuscarEmployer()

	for _, v := range emps {
		fmt.Println(" employee ", v)
	}

}
