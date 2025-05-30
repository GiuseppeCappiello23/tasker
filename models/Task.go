package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Title       string
	DueDate     string
	Status      string
	Description string
}

func (t *Task) NewTask(task Task) error {
	db, err := sql.Open("sqlite3", "./miodb.sqlite")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Task(Title, Description, DueDate, Status) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Title, task.Description, task.DueDate, task.Status)
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) ListTasks() error {
	db, err := sql.Open("sqlite3", "./miodb.sqlite")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM Task")
	if err != nil {
		return err
	}
	defer stmt.Close()

	results, err := stmt.Exec()
	if err != nil {
		return err
	}
	
	for r := range results {

	}
}
