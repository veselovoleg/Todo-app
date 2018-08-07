package models

import (
	"database/sql"

	//-sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// Task - id + name of task
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection - array of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

//GetTasks get
func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)
	// Exit if SQL not working
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name)
		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

//PutTask - add new task
func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	// change '?' symbol to 'name'
	result, err2 := stmt.Exec(name)
	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

//DeleteTask - delete task
func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	// change '?' symbol to 'name'
	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}
