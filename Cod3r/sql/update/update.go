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

	statement, _ := db.Prepare("UPDATE todos SET done = $1 WHERE id = $2")

	statement.Exec(true, 1)
}
