package models

import (
	"database/sql"
	"fmt"

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

	results, err := db.Query("SELECT * FROM Task")
	if err != nil {
		return err
	}
	defer results.Close()

	fmt.Printf("%s\t\t%s\t\t%s\t\t%s\t\t%s", "ID", "TITLE", "DESCRIPTION", "DUEDATE", "STATUS")
	for results.Next() {
		var id int
		var title, description, duedate, status string

		err := results.Scan(&id, &title, &description, &duedate, &status)
		if err != nil {
			return err
		}
		fmt.Printf("%d\t\t%s\t\t%s\t\t%s\t\t%s\n", id, title, description, duedate, status)
	}

	if err = results.Err(); err != nil {
		return err
	}

	return nil
}
