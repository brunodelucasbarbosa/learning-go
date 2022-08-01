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

	statement, _ := db.Prepare("DELETE FROM todos WHERE id = $1")

	statement.Exec(1)
}
