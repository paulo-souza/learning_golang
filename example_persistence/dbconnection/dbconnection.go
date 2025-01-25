package dbconnection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const postgresDriver = "postgres"
const user = "dev_user"
const host = "localhost"
const port = "2222"
const password = "dev_user"
const dbName = "minha_empresa"

var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open(postgresDriver, dataSourceName)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}
