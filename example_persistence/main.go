package main

import (
	"github.com/paulo-souza/learning_golang/example_persistence/dbconnection"
)

func main() {
	db, err := dbconnection.OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// bookInsert := crud.LacosDeAmor(db)
	// bookInsert.Insert()

	// bookUpdate := crud.ALendaDeAng(db)
	// bookUpdate.ID = 3
	// bookUpdate.Update()

	// bookSelectId := crud.Book{ID: 1, Db: db}
	// fmt.Println(bookSelectId.SelectById())
}
