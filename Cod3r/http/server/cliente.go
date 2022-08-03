package main

import (
	"cod3r/sql/dbConfig"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Usuario struct {
	Id   int64  `json:"id"`
	Name string `json:"nome"`
}

func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/") // remove /usuarios/ e tudo que tiver atrás do inicio da string

	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorId(w, r, id)
	case r.Method == "POST":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Não encontrado")
	}
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.Id, &usuario.Name)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func usuarioPorId(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var usuario Usuario
	db.QueryRow("SELECT * FROM usuarios where id = $1", id).Scan(&usuario.Id, &usuario.Name)

	json, _ := json.Marshal(usuario)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
