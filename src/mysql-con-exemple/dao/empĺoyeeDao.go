package dao

import (
	"log"
	"mysql-con-exemple/conf"
	"mysql-con-exemple/model"
)

func CadastrarEmployee(e model.Employee) {
	db := conf.DbConn()
	insForm, err := db.Prepare("INSERT INTO employee(name, city) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(e.Name, e.City)
	log.Println("INSERT: Name: " + e.Name + " | City: " + e.City)
	defer db.Close()
}

func BuscarEmployer() []model.Employee {
	db := conf.DbConn()
	selDB, err := db.Query("SELECT * FROM employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := model.Employee{}
	res := []model.Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}

	defer db.Close()

	return res
}
