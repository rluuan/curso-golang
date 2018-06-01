package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/cursogo")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("INSERT INTO usuarios(id, nome) values(?, ?)")

	stmt.Exec(9, "Bia")
	stmt.Exec(10, "carlos")
	_, err = stmt.Exec(11, "maria")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
