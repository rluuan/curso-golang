package main

import (
	"database/sql"

	_ "github.com/Go-SQL-Driver/MySQL"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, error := db.Exec(sql)
	if error != nil {
		panic(error)
	}
	return result
}
func main() {
	db, error := sql.Open("mysql", "root:@/")
	if error != nil {
		panic(error)
	}
	defer db.Close()

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY(id)
	)`)
}
