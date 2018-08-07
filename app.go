package main

//https://tehnojam.pro/category/development/sozdanie-odnostranichnogo-veb-prilozhenija-na-go-echo-i-vue.html

import (
	"database/sql"
	"fmt"
	"go-echo-vue/handlers"

	"github.com/labstack/echo"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//Inittialize DB
	db := initDB("storage.db")
	migrate(db)

	// Create a new instance of Echo
	e := echo.New()
	//Define rotes
	e.File("/", "public/index.html")
	e.File("/main.js", "public/scripts/main.js")
	e.File("/style.css", "public/styles/style.css")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))
	// Start as a web server
	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	//Open db, or create db
	db, err := sql.Open("sqlite3", filepath)
	//Check connection errors
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}
	fmt.Println("DB connected")
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL
	);
	`

	_, err := db.Exec(sql)
	//Check errors
	if err != nil {
		panic(err)
	}
}
