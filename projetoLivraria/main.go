package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	//"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	api_livros = "/api/v1/livros"
)

type Livro struct {
	Id, Nome, Isbn string
}

type livraria struct {
	dbHost, dbPass, dbName string
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost:3306"
	}

	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		dbPass = "tech#1000"
	}

	apiPath := os.Getenv("API_PATH")
	if apiPath == "" {
		apiPath = api_livros
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "livraria"
	}

	l := livraria{
		dbHost: dbHost,
		dbPass: dbPass,
		dbName: dbName,
	}

	r := mux.NewRouter()
	r.HandleFunc(api_livros, l.buscarLivros).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func (l livraria) buscarLivros(http.ResponseWriter, *http.Request) {
	// abrir conexão
	db := l.openConnection()

	// carregar todos os livros
	rows, err := db.Query("select * from livros")
	if err != nil {
		log.Fatalf("Query error failed livros: %s\n", err.Error())
	}

	livros := []Livro{}

	for rows.Next() {
		var id, nome, isbn string
		err := rows.Scan(&id, &nome, &isbn)
		if err != nil {
			log.Fatalf("Query scanning the row error: %s\n", err.Error())
		}

		aLivro := Livro{
			Id:   id,
			Nome: nome,
			Isbn: isbn,
		}
		livros = append(livros, aLivro)
	}

	// fechar conexão
}

func (l livraria) openConnection() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s", "root", l.dbPass, l.dbHost, l.dbName))

	if err != nil {
		log.Fatal("Open connection failed %s\n", err.Error())
	}

	return db
}
