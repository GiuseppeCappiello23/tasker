package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Name        string
	DueDate     string
	Status      string
	Description string
}

func (t *Task) NewTask(s string) error {
	db, err := sql.Open("sqlite3", "./miodb.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO utenti(nome, email) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
}
