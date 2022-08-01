package main

import (
	"database/sql"
	"fmt"
	"learning-go/Cod3r/sql/dbConfig"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	statement, _ := db.Prepare("INSERT INTO todos(todo_name, done) VALUES($1, $2)")
	statement.Exec("Arrumar quarto1", false)
	statement.Exec("Arrumar quarto2", false)

	res, _ := statement.Exec("Comprar pão", false)

	id, _ := res.LastInsertId()

	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}
