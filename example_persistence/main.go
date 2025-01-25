package main

import (
	"fmt"

	"github.com/paulo-souza/learning_golang/example_persistence/crud"
	"github.com/paulo-souza/learning_golang/example_persistence/dbconnection"
)

func main() {
	db, err := dbconnection.OpenConnection()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("You are connected to the postgres database...")
	fmt.Println("This is amazing!")

	defer db.Close()

	bookInsert := crud.LacosDeAmor(db)
	bookInsert.Insert()
}
