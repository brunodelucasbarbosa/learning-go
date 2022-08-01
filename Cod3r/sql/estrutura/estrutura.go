package main

import (
	"cod3r/sql/dbConfig"
	"database/sql"

	_ "github.com/lib/pq"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	db, err := sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "DROP DATABASE IF EXISTS curso_go_cod3r")
	exec(db, "CREATE DATABASE curso_go_cod3r")

}
