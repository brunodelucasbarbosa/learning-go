package main

import (
	"database/sql"
	"fmt"
	"learning-go/Cod3r/sql/dbConfig"

	_ "github.com/lib/pq"
)

type todo struct {
	id   int
	name string
	done bool
}

func main() {
	db, err := sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM todos where id > $1", 2)
	defer rows.Close()

	for rows.Next() {
		t := todo{}
		rows.Scan(&t.id, &t.name, &t.done)
		fmt.Println(t)
	}

}
