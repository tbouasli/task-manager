package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db_path = os.Getenv("HOME") + "/.local/share/tasks/db"
var db, _ = sql.Open("sqlite3", "db.sql")


func Init() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, name TEXT, completed BOOLEAN)")
	if err != nil {
		panic(err)
	}
}

type Task struct {
	ID        int
	Name      string
	Completed bool
}

func (t *Task) Save() {
	_, err := db.Exec("INSERT INTO tasks (id, name, completed) VALUES (?, ?, ?)", t.ID, t.Name, t.Completed)
	if err != nil {
		panic(err)
	}
}

func (t *Task) Delete() {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", t.ID)
	if err != nil {
		panic(err)
	}
}

func (t *Task) Complete() {
	t.Completed = true
	_, err := db.Exec("UPDATE tasks SET completed = 1 WHERE id = ?", t.ID)
	if err != nil {
		panic(err)
	}
}

func (t *Task) Uncomplete() {
	t.Completed = false
	_, err := db.Exec("UPDATE tasks SET completed = 0 WHERE id = ?", t.ID)
	if err != nil {
		panic(err)
	}
}

func List() []Task {
	rows, err := db.Query("SELECT id, name, completed FROM tasks")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		t := Task{}
		rows.Scan(&t.ID, &t.Name, &t.Completed)
		tasks = append(tasks, t)
	}

	return tasks
}
