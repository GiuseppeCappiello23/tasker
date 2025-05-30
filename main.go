package main

/*
FUNZIONI PRINCIPALI DELL'APPLICAZIONE DA TERMINALE:
1. Aggiungere nuovo task (add)
	1. Dare titolo al task (--title -t)
	2. Dare descrizione al task (--description -d)
	3. Definire data completamento task (--due-date)
2. Modificare informazioni sul task (edit)

3. Listare tutti i task

Appena creato il task avrà lo stato "pending"

Si potranno aggiornare informazioni relative al task con il comando "tasker {id} edit"

Il comando edit potrà modificare tutte le informazioni relative al task (Titolo, Descrizione, Due Date, Status)

I flag accettati dal comando edit sono i medesimi del comando add con l'aggiunta di :
-- status per modificare lo status
Gli stati accettati sono :
p (pending - task non ancora iniziato)
i (in-progress - task iniziato)
d (done - task completato)
*/

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./miodb.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlCreate := `
	CREATE TABLE IF NOT EXISTS Tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		Title TEXT NOT NULL UNIQUE,
		Description TEXT,
		DueDate TEXT,
		Status TEXT
	);`
	_, err = db.Exec(sqlCreate)
	if err != nil {
		log.Fatalf("Errore nella creazione tabella: %s", err)
	}

	args := os.Args

	for i, arg := range args {
		fmt.Printf("%d: %s\n", i, arg)
	}

	// fmt.Println("Hello world")
}
