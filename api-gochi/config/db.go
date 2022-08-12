package dbConfig

import (
	"database/sql"
	"fmt"
)

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "admin"
const DbName = "course_gochi"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)


func Connect() (*sql.DB, error) {
	db, err := sql.Open(PostgresDriver, DataSourceName)
	if err != nil {
		return nil, err
	}
	
	return db, nil
}