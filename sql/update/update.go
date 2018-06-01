package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/cursogo")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
	stmt.Exec("Rodrigo", 1)
	stmt.Exec("Marcos", 2)

	stmt, _ = db.Prepare("DELETE FROM usuarios WHERE id = ?")
	stmt.Exec(1)
}
