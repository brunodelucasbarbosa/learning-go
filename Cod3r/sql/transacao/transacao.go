package main

import (
	"database/sql"
	"learning-go/Cod3r/sql/dbConfig"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	transaction, _ := db.Begin()
	statement, _ := transaction.Prepare("INSERT INTO todos(todo_name, done) VALUES($1, $2)")

	statement.Exec("Arrumar sala", false)
	_, err = statement.Exec("Arrumar cozinha", false)

	if err != nil {
		transaction.Rollback()
	}
	transaction.Commit()

}
