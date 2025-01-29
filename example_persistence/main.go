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

	defer db.Close()

	// bookInsert := crud.LacosDeAmor(db)
	// bookInsert.Insert()

	// bookUpdate := crud.JavaUseACabeca(db)
	// bookUpdate.ID = 42
	// bookUpdate.Update()

	// bookSelectId := crud.Book{ID: 42, Db: db}
	// fmt.Println(bookSelectId.SelectById())

	booksSelectAll := crud.Book{Db: db}
	allBooks := booksSelectAll.SelectAll()
	fmt.Printf("%#v\n", allBooks)
	fmt.Println("total books =", len(allBooks))

	// bookDelete := crud.Book{Db: db, ID: 42}
	// bookDelete.Delete()

	// books := crud.TodosLivros()
	// book := crud.Book{Db: db}
	// book.InsertList(books)
}
