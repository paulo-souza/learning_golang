package crud

import (
	"database/sql"
	"fmt"
)

type Book struct {
	ID                                     int
	Price                                  float64
	Name, Format, Author, Genre, Publisher string
	Db                                     *sql.DB
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func (b *Book) Insert() {

	sqlStatement := "INSERT INTO book (price,name, format, author, genre, publisher) VALUES ($1,$2,$3,$4,$5,$6)"

	insert, err := b.Db.Prepare(sqlStatement)
	checkError(err)

	result, err := insert.Exec(b.Price, b.Name, b.Format, b.Author, b.Genre, b.Publisher)
	checkError(err)

	affect, err := result.RowsAffected()
	checkError(err)

	fmt.Println("Números de registros inserido(s):", affect)
}
