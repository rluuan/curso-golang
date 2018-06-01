package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// User estrutura
type User struct {
	ID   int    `json:"id"`
	Name string `json:"nome"`
}

// UserHandler | analisa o request e delega para funcao adequada
func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		UserID(w, r, id)
	case r.Method == "GET":
		UserAll(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(w, "Sorry... :(")
	}
}

// UserID | busca por id
func UserID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var u User
	db.QueryRow("SELECT id, nome FROM usuarios WHERE id = ?", id).Scan(&u.ID, &u.Name)

	json, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

// UserAll | busca todos
func UserAll(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios")
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}
	json, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
