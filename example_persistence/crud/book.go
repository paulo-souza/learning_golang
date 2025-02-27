package crud

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/paulo-souza/learning_golang/example_persistence/entity"
)

type Book struct {
	ID        int
	Price     float64
	Name      string
	Format    string
	Author    string
	Genre     string
	Publisher string
	Db        *sql.DB
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func (b *Book) isIdZero() bool {
	const notFound int = 0
	return notFound == b.ID
}

func (b *Book) execSQL(sqlStatement string, vals ...interface{}) int64 {
	prepStmt, err := b.Db.Prepare(sqlStatement)
	checkError(err)

	result, err := prepStmt.Exec(vals...)
	checkError(err)

	affect, err := result.RowsAffected()
	checkError(err)

	return affect
}

func (b *Book) Insert() {
	sqlStatement := "INSERT INTO book (price,name,format,author,genre,publisher) VALUES ($1,$2,$3,$4,$5,$6)"

	affect := b.execSQL(sqlStatement, b.Price, b.Name, b.Format, b.Author, b.Genre, b.Publisher)
	fmt.Println("Números de registros inserido(s):", affect)
}

func (b *Book) InsertList(books []Book) {
	sqlStatement := "INSERT INTO book (price,name,format,author,genre,publisher) VALUES "

	var inserts []string
	var vals []interface{}
	var argNum int
	var row string

	increArg := func() int { argNum += 1; return argNum }

	for _, book := range books {
		row = fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d)", increArg(), increArg(), increArg(), increArg(), increArg(), increArg())
		inserts = append(inserts, row)
		vals = append(vals, book.Price, book.Name, book.Format, book.Author, book.Genre, book.Publisher)
	}

	sqlStatement += strings.Join(inserts, ",")

	affect := b.execSQL(sqlStatement, vals...)
	fmt.Println("Números de registros inserido(s):", affect)
}

func (b *Book) Update() {
	if b.isIdZero() {
		log.Printf("ID nao foi encontrado. %#v", b)
		return
	}

	sqlStatement := `UPDATE book 
					SET 
					price=$1,name=$2,
					format=$3,author=$4,
					genre=$5,publisher=$6
					WHERE id=$7`

	affect := b.execSQL(sqlStatement, b.Price, b.Name, b.Format, b.Author, b.Genre, b.Publisher, b.ID)
	fmt.Println("Números de registros atualizado(s):", affect)
}

func (b *Book) Delete() {
	if b.isIdZero() {
		log.Printf("ID nao foi encontrado. %#v", b)
		return
	}

	sqlStatement := "DELETE FROM book WHERE id=$1"

	affect := b.execSQL(sqlStatement, b.ID)
	fmt.Println("Números de registros deletado(s):", affect)
}

func (b *Book) SelectById() (bookFound entity.Book) {
	sqlStatement := "SELECT id,price,name,format,author,genre,publisher FROM book WHERE id = $1"
	bookValues := []interface{}{
		&bookFound.ID, &bookFound.Price, &bookFound.Name,
		&bookFound.Format, &bookFound.Author,
		&bookFound.Genre, &bookFound.Publisher,
	}

	err := b.Db.QueryRow(sqlStatement, b.ID).Scan(bookValues...)

	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("ID: %d nao foi encontrado. %#v\n", b.ID, b)
		return
	}

	checkError(err)

	return
}

func (b *Book) SelectAll() (books []entity.Book) {
	sqlStatement := "SELECT id,price,name,format,author,genre,publisher FROM book"
	res, err := b.Db.Query(sqlStatement)

	checkError(err)

	for res.Next() {
		var book entity.Book

		err = res.Scan(&book.ID, &book.Price, &book.Name,
			&book.Format, &book.Author,
			&book.Genre, &book.Publisher)

		checkError(err)
		books = append(books, book)
	}
	return
}
